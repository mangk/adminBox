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
    },
    devMode: true
  }

  const component = defineAsyncComponent(() => loadModule(`${name}.vue`, options))
  return markRaw(component)
}
