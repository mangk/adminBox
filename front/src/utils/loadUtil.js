import { loadModule } from 'vue3-sfc-loader'
import * as Vue from 'vue'
import { defineAsyncComponent, markRaw } from 'vue'
import http from './requester'

// 用一个对象来跟踪已加载的脚本及其回调队列
const loadedScripts = {}
export function loadJS(url, callback = null) {
  // 如果该脚本已经加载过，直接将回调加入队列并执行
  if (loadedScripts[url]) {
    if (typeof callback === 'function') {
      loadedScripts[url].push(callback)
      callback() // 立即执行新的回调
    }
    return
  }

  // 初始化回调队列
  loadedScripts[url] = []
  if (typeof callback === 'function') {
    loadedScripts[url].push(callback)
  }

  const script = document.createElement('script')
  script.type = 'text/javascript'
  script.src = url

  script.onload = () => {
    // 脚本加载完成，依次执行所有队列中的回调函数
    loadedScripts[url].forEach((cb) => cb())
  }

  document.head.appendChild(script)
}

export function loadTMPL(url, name = 'myConvert') {
  const options = {
    moduleCache: { vue: Vue },
    getFile() {
      return http(url, {
        method: 'GET'
      }).then((response) => {
        return response
      })
    },
    addStyle(styleString) {
      const style = document.createElement('style')
      style.setAttribute('id', name)
      style.textContent = styleString
      const ref =
        document.head.getElementsByTagName('style')[
          document.head.getElementsByTagName('style').length - 1
        ] || null
      document.head.insertBefore(style, ref)
    }
  }

  const component = defineAsyncComponent(() => loadModule(`${name}.vue`, options))
  return markRaw(component)
}
