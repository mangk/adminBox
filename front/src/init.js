import * as ElIconModules from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'
import { loadModule } from 'vue3-sfc-loader'
import * as Vue from 'vue'
import { defineAsyncComponent, markRaw } from 'vue'
import http from './utils/requester'
import { importView } from './utils/routerFormat'

// 公共方法
const loadJS = (src) => {
  return new Promise((resolve, reject) => {
    let script = document.createElement('script')
    script.type = 'text/javascript'
    script.src = src
    document.body.appendChild(script)
    script.onload = () => {
      resolve()
    }
    script.onerror = () => {
      reject()
    }
  })
}
const loadTMPL = (url, name = 'myConvert') => {
  const options = {
    moduleCache: { vue: Vue },
    getFile() {
      return fetch(url, {
        method: 'GET'
      }).then((response) => (response.ok ? response.text() : Promise.reject(response)))
    },
    addStyle(styleString) {
      const style = document.createElement('style')
      style.setAttribute('id', name)
      style.textContent = styleString
      const ref = document.head.getElementsByTagName('style')[0] || null
      document.head.insertBefore(style, ref)
    }
  }

  const component = defineAsyncComponent(() => loadModule(`${name}.vue`, options))
  return markRaw(component)
}

const register = async (app) => {
  for (const iconName in ElIconModules) {
    app.component(iconName, ElIconModules[iconName])
  }

  // 注册全局组件
  const FileUpload = await importView('views/util/fileUpload.vue')()
  app.component('FileUpload', FileUpload.default || FileUpload)

  // 注册全局变量
  app.config.globalProperties.$adminX = window.adminX

  // 注册js加载方法
  app.config.globalProperties.$loadJS = async (src) => {
    await loadJS(src)
  }

  // 注册服务端模版加载方法
  app.config.globalProperties.$loadTMPL = loadTMPL

  // 注册全局Element消息
  app.config.globalProperties.$message = ElMessage

  // 注册http请求方法
  app.config.globalProperties.$http = http
}

export default {
  install: (app) => {
    register(app)
  }
}