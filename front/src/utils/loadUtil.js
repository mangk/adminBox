import { loadModule } from 'vue3-sfc-loader'
import * as Vue from 'vue'
import { defineAsyncComponent, markRaw } from 'vue'
import http from './requester'

// 用一个对象来跟踪已加载的脚本及其回调队列
const loadedScripts = {}
export function loadJS(urls, callback = null) {
  // 确保urls是一个数组
  if (typeof urls === 'string') {
    urls = [urls];
  }

  // 创建一个Promise数组，用于并行加载所有脚本
  const promises = urls.map((url) => new Promise((resolve, reject) => {
    // 如果该脚本已加载过，直接跳过加载
    if (loadedScripts[url]) {
      resolve();
      return;
    }

    const script = document.createElement('script');
    script.type = 'text/javascript';
    script.src = url;

    script.onload = resolve; // 脚本加载成功
    script.onerror = reject; // 脚本加载失败

    document.head.appendChild(script);
    loadedScripts[url] = true; // 标记脚本已加载
  }));

  // 所有脚本加载完毕后，执行回调
  Promise.all(promises).then(() => {
    if (typeof callback === 'function') {
      callback();
    }
  }).catch((error) => {
    console.error('脚本加载失败:', error);
  });
}
/**
 * Robust loader for .vue single-file components fetched as text.
 *
 * @param {string} url - 请求 SFC 文本的 URL（返回值可为 string，或 axios-like response，或 fetch Response）
 * @param {string} name - 用作样式 id / 日志前缀
 * @param {object} userOptions - 可选项:
 *    { number timeout = 15000, boolean devMode = true,
 *      boolean retryOnFail = false, number maxRetries = 2,
 *      Component loadingComponent, Component errorComponent, number delay = 200, number errorTimeout = 0 }
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

  // accumulate compiler/runtime logs from vue3-sfc-loader
  let compileLogs = []

  const options = {
    moduleCache: { vue: Vue },
    // getFile must return SFC text. We make it robust to various http wrappers.
    getFile() {
      // prefer dev-mode cache busting to avoid stale files
      const fetchUrl = devMode ? `${url}${url.includes('?') ? '&' : '?'}t=${Date.now()}` : url

      // fetch wrapper that supports: plain string, fetch Response, axios-like {data,status}, custom http*
      const fetchPromise = (async () => {
        try {
          const response = await http(fetchUrl, { method: 'GET' })
          // if using fetch API
          if (response && typeof Response !== 'undefined' && response instanceof Response) {
            if (!response.ok) throw new Error(`HTTP ${response.status} ${response.statusText}`)
            return await response.text()
          }
          // axios-like (response.data)
          if (response && typeof response === 'object' && 'data' in response) {
            // assume response.data is text
            return typeof response.data === 'string' ? response.data : JSON.stringify(response.data)
          }
          // plain string (some wrappers directly return text)
          if (typeof response === 'string') return response
          // some wrappers put body
          if (response && typeof response === 'object' && 'body' in response) {
            return typeof response.body === 'string' ? response.body : JSON.stringify(response.body)
          }

          // unknown shape
          throw new Error('Unsupported response shape from http()')
        } catch (err) {
          // rethrow to be caught by outer Promise.race
          throw err
        }
      })()

      // timeout guard
      const timeoutPromise = new Promise((_, reject) =>
        setTimeout(() => reject(new Error(`Request timed out after ${timeout} ms`)), timeout)
      )

      return Promise.race([fetchPromise, timeoutPromise]).catch((err) => {
        // normalize message so loader sees a clear error
        const e = new Error(`[loadTMPL:getFile] ${err && err.message ? err.message : err}`)
        return Promise.reject(e)
      })
    },

    // insert or replace style tag by id to avoid duplicated styles
    addStyle(styleString) {
      try {
        let style = document.getElementById(name)
        if (!style) {
          style = document.createElement('style')
          style.setAttribute('id', name)
          // append to head end
          document.head.appendChild(style)
        }
        style.textContent = styleString
      } catch (err) {
        // don't break loader for style injection errors, but log
        console.warn(`[${name}] addStyle failed:`, err)
      }
    },

    // hook for vue3-sfc-loader to report compiler/runtime messages
    log(type, ...args) {
      try {
        // record for post-checks (we'll throw if there are compiler errors)
        compileLogs.push({ type, args })
        // forward to dev console with a prefix
        const fn = console[type] || console.log
        fn.call(console, `[vue3-sfc-loader:${name}]`, ...args)
      } catch (e) {
        // swallow any logging errors
      }
    },

    devMode
  }

  // loader wrapper which will call loadModule and perform sanity checks
  const loader = async () => {
    // reset compile logs before each load
    compileLogs = []

    let mod
    try {
      mod = await loadModule(`${name}.vue`, options)
    } catch (err) {
      // this is likely network / parse errors. normalize and rethrow
      console.error(`[${name}] loadModule failed:`, err)
      throw err
    }

    // if loader emitted any error-level logs during compile, surface them as thrown error
    const errs = compileLogs
      .filter((l) => l.type === 'error' || l.type === 'warn')
      .map((l) => (Array.isArray(l.args) ? l.args.map(String).join(' ') : String(l.args)))
    if (errs.length > 0) {
      const aggregated = errs.join('\n')
      console.error(`[${name}] compile/loader reports:\n${aggregated}`)
      // Throw a meaningful error so defineAsyncComponent.onError runs and console shows stack
      throw new Error(`${name} compile/loader errors:\n${aggregated}`)
    }

    // module shape checks: prefer ES module default export, else module itself
    const component = mod && (mod.default || mod)
    if (!component || (typeof component !== 'object' && typeof component !== 'function')) {
      console.error(`[${name}] module did not export a Vue component:`, mod)
      throw new Error(`${name} did not export a valid component`)
    }

    return _markRaw(component)
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
