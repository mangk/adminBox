import * as Vue from 'vue'
import { defineAsyncComponent, markRaw } from 'vue'
import { compileScript, compileStyle, compileTemplate, parse } from '@vue/compiler-sfc'
import http from './requester'

// 用一个对象来跟踪已加载的脚本及其回调队列
const loadedScripts = {};
export function loadJS(urls, callback = null) {
  if (typeof urls === 'string') {
    urls = [urls];
  }

  const promises = urls.map(url => {
    if (loadedScripts[url]) {
      return loadedScripts[url]; // 返回同一个Promise
    }

    loadedScripts[url] = new Promise((resolve, reject) => {
      const script = document.createElement('script');
      script.src = url;
      script.onload = resolve;
      script.onerror = reject;
      document.head.appendChild(script);
    });

    return loadedScripts[url];
  });

  Promise.all(promises).then(() => {
    callback && callback();
  });
}

function hashString(input) {
  let hash = 2166136261
  for (let i = 0; i < input.length; i += 1) {
    hash ^= input.charCodeAt(i)
    hash = Math.imul(hash, 16777619)
  }
  return (hash >>> 0).toString(16)
}

function stripVueMacroImports(source) {
  const macroSet = new Set([
    'defineProps',
    'defineEmits',
    'defineExpose',
    'defineOptions',
    'defineSlots',
    'withDefaults'
  ])

  const importRegex = /^import\s+\{([^}]+)\}\s+from\s+['"]vue['"]\s*;?\s*$/gm
  return source.replace(importRegex, (_, namedPart) => {
    const keepList = namedPart
      .split(',')
      .map((part) => part.trim())
      .filter(Boolean)
      .filter((part) => {
        const m = part.match(/^([A-Za-z_$][\w$]*)(?:\s+as\s+([A-Za-z_$][\w$]*))?$/)
        if (!m) return true
        const importedName = m[1]
        return !macroSet.has(importedName)
      })

    if (keepList.length === 0) return ''
    return `import { ${keepList.join(', ')} } from 'vue'`
  })
}

function formatCompileErrors(title, errors) {
  const arr = Array.isArray(errors) ? errors : [errors]
  return arr
    .map((err) => {
      if (!err) return ''
      if (typeof err === 'string') return err
      return err.message || String(err)
    })
    .filter(Boolean)
    .map((msg) => `[${title}] ${msg}`)
    .join('\n')
}

function transformVueImports(code) {
  const normalizeNamedImports = (rawNamed) => {
    return rawNamed
      .split(',')
      .map((part) => part.trim())
      .filter(Boolean)
      .map((part) => {
        const match = part.match(/^([A-Za-z_$][\w$]*)(?:\s+as\s+([A-Za-z_$][\w$]*))?$/)
        if (!match) return part
        return match[2] ? `${match[1]}: ${match[2]}` : match[1]
      })
      .join(', ')
  }

  let output = code

  output = output.replace(
    /^import\s+([A-Za-z_$][\w$]*)\s*,\s*\{([^}]+)\}\s+from\s+['"]vue['"]\s*;?\s*$/gm,
    (_, defaultName, namedPart) =>
      `const ${defaultName} = __deps.vue.default || __deps.vue\nconst { ${normalizeNamedImports(namedPart)} } = __deps.vue`
  )

  output = output.replace(
    /^import\s+\{([^}]+)\}\s+from\s+['"]vue['"]\s*;?\s*$/gm,
    (_, namedPart) => `const { ${normalizeNamedImports(namedPart)} } = __deps.vue`
  )

  output = output.replace(
    /^import\s+\*\s+as\s+([A-Za-z_$][\w$]*)\s+from\s+['"]vue['"]\s*;?\s*$/gm,
    (_, localName) => `const ${localName} = __deps.vue`
  )

  output = output.replace(
    /^import\s+([A-Za-z_$][\w$]*)\s+from\s+['"]vue['"]\s*;?\s*$/gm,
    (_, defaultName) => `const ${defaultName} = __deps.vue.default || __deps.vue`
  )

  return output
}

function compileToFactory(code, filename, returnSymbol) {
  let normalized = transformVueImports(code)
  normalized = normalized.replace(
    /export\s*\{\s*([A-Za-z_$][\w$]*)\s+as\s+default\s*\}\s*;?/gm,
    'const __sfc__ = $1'
  )
  normalized = normalized.replace(/export\s+default/g, 'const __sfc__ =')
  normalized = normalized.replace(/export\s+function\s+render/g, 'function render')
  normalized = normalized.replace(/export\s+(const|let|var|function|class)\s+/g, '$1 ')

  const wrapped = `
"use strict"
${normalized}
return ${returnSymbol};
`
  try {
    return new Function('__deps', wrapped)
  } catch (err) {
    throw new Error(`[${filename}] compile factory failed: ${err && err.message ? err.message : err}`, {
      cause: err
    })
  }
}

function runComponentScript(code, filename) {
  const factory = compileToFactory(code, filename, '__sfc__')
  const component = factory({ vue: Vue })
  if (!component || (typeof component !== 'object' && typeof component !== 'function')) {
    throw new Error(`[${filename}] script does not export a valid component`)
  }
  return component
}

function runRenderFactory(code, filename) {
  const factory = compileToFactory(code, filename, 'render')
  const render = factory({ vue: Vue })
  if (typeof render !== 'function') {
    throw new Error(`[${filename}] template does not export render function`)
  }
  return render
}

/**
 * 动态加载并编译远程 .vue SFC（不依赖 vue3-sfc-loader）。
 */
export function loadTMPL(url, name = 'myConvert', userOptions = {}) {
  const {
    timeout = 5000,
    devMode = false,
    retryOnFail = false,
    maxRetries = 2,
    loadingComponent = null,
    errorComponent = null,
    delay = 200,
    errorTimeout = 0 // ms for defineAsyncComponent timeout
  } = userOptions

  // small helpers to be resilient if user environment exposes different Vue APIs
  const _defineAsyncComponent =
    typeof defineAsyncComponent !== 'undefined'
      ? defineAsyncComponent
      : (Vue && Vue.defineAsyncComponent)
        ? Vue.defineAsyncComponent
        : null
  const _markRaw =
    typeof markRaw !== 'undefined'
      ? markRaw
      : (Vue && Vue.markRaw)
        ? Vue.markRaw
        : (v) => v // fallback: identity
  if (!_defineAsyncComponent) {
    throw new Error('defineAsyncComponent is not available in this environment')
  }

  const fetchSFC = () => {
    const fetchUrl = devMode ? `${url}${url.includes('?') ? '&' : '?'}t=${Date.now()}` : url
    const fetchPromise = (async () => {
      const response = await http(fetchUrl, { method: 'GET' })
      if (response && typeof Response !== 'undefined' && response instanceof Response) {
        if (!response.ok) throw new Error(`HTTP ${response.status} ${response.statusText}`)
        return await response.text()
      }
      if (response && typeof response === 'object' && 'data' in response) {
        return typeof response.data === 'string' ? response.data : JSON.stringify(response.data)
      }
      if (typeof response === 'string') return response
      if (response && typeof response === 'object' && 'body' in response) {
        return typeof response.body === 'string' ? response.body : JSON.stringify(response.body)
      }
      throw new Error('Unsupported response shape from http()')
    })()

    const timeoutPromise = new Promise((_, reject) =>
      setTimeout(() => reject(new Error(`Request timed out after ${timeout} ms`)), timeout)
    )

    return Promise.race([fetchPromise, timeoutPromise]).catch((err) => {
      const e = new Error(`[loadTMPL:getFile] ${err && err.message ? err.message : err}`)
      return Promise.reject(e)
    })
  }

  const applyStyles = (styleId, cssText) => {
    if (!cssText) return
    try {
      let style = document.getElementById(styleId)
      if (!style) {
        style = document.createElement('style')
        style.setAttribute('id', styleId)
        document.head.appendChild(style)
      }
      style.textContent = cssText
    } catch (err) {
      console.warn(`[${name}] addStyle failed:`, err)
    }
  }

  const loader = async () => {
    const filename = `${name}.vue`
    const scopeId = `data-v-${hashString(`${name}::${url}`)}`
    const styleId = `${name}__runtime_style`
    const source = stripVueMacroImports(await fetchSFC())

    const parsed = parse(source, { filename })
    const parseErrors = parsed.errors || []
    if (parseErrors.length > 0) {
      throw new Error(formatCompileErrors(`${name}:parse`, parseErrors))
    }

    const descriptor = parsed.descriptor
    const hasScoped = descriptor.styles.some((styleBlock) => styleBlock.scoped)

    const styles = []
    for (const styleBlock of descriptor.styles) {
      const styleRes = compileStyle({
        filename,
        id: scopeId,
        source: styleBlock.content,
        scoped: styleBlock.scoped
      })
      if (styleRes.errors && styleRes.errors.length > 0) {
        throw new Error(formatCompileErrors(`${name}:style`, styleRes.errors))
      }
      styles.push(styleRes.code)
    }
    if (styles.length > 0) {
      applyStyles(styleId, styles.join('\n'))
    }

    let component = {}
    let scriptResult = null
    if (descriptor.script || descriptor.scriptSetup) {
      scriptResult = compileScript(descriptor, {
        id: scopeId,
        inlineTemplate: false
      })
      component = runComponentScript(scriptResult.content, filename)
    }

    if (descriptor.template) {
      const templateRes = compileTemplate({
        id: scopeId,
        filename,
        source: descriptor.template.content,
        scoped: hasScoped,
        compilerOptions: {
          bindingMetadata: scriptResult ? scriptResult.bindings : undefined
        }
      })
      if (templateRes.errors && templateRes.errors.length > 0) {
        throw new Error(formatCompileErrors(`${name}:template`, templateRes.errors))
      }
      if (templateRes.tips && templateRes.tips.length > 0) {
        for (const tip of templateRes.tips) {
          console.warn(`[${name}] template tip:`, tip)
        }
      }
      component.render = runRenderFactory(templateRes.code, filename)
    }

    if (hasScoped) {
      component.__scopeId = scopeId
    }
    component.__file = filename

    try {
      return _markRaw(component)
    } catch (err) {
      console.error(`[${name}] compile/load failed:`, err)
      throw err
    }
  }

  // default friendly error component (shows minimal message in DOM, instructs to check console)
  const DefaultErrorComp =
    errorComponent ||
    (Vue && Vue.defineComponent
      ? Vue.defineComponent({
        props: { message: { type: String, default: '' } },
        setup(props) {
          return () =>
            Vue.h(
              'div',
              {
                style:
                  'padding:12px;border:1px solid rgba(255,0,0,0.25);background:rgba(255,0,0,0.03);color:#900;font-family:monospace'
              },
              [
                Vue.h('div', { style: 'font-weight:600;margin-bottom:6px' }, `Failed to load component "${name}".`),
                Vue.h('div', {}, 'See console for details.'),
                props.message ? Vue.h('pre', { style: 'white-space:pre-wrap;margin-top:8px' }, props.message) : null
              ]
            )
        }
      })
      : null)

  // configure retries via onError
  let attempts = 0
  const asyncComponent = _defineAsyncComponent({
    loader: async () => {
      attempts = 0
      return loader()
    },
    loadingComponent: loadingComponent || null,
    errorComponent: DefaultErrorComp,
    delay,
    timeout: errorTimeout || undefined,
    // don't use suspensible mode (makes SSR / Suspense interactions predictable)
    suspensible: false,
    onError(err, retry, fail) {
      attempts += 1
      console.error(`[${name}] async component error (attempt ${attempts}):`, err)
      if (retryOnFail && attempts <= maxRetries) {
        // small backoff
        setTimeout(() => {
          console.info(`[${name}] retrying load (attempt ${attempts + 1})`)
          retry()
        }, 200 * attempts)
      } else {
        // fail and render errorComponent; also rethrow so it's visible in console
        fail()
      }
    }
  })

  // return a markRaw async component to prevent Vue from observing it
  return _markRaw(asyncComponent)
}
