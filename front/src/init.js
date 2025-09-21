import * as ElIconModules from '@element-plus/icons-vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import http from './utils/requester'
import { importView } from './utils/routerFormat'
import { loadJS, loadTMPL } from './utils/loadUtil'
import { useUserStore } from '@/pinia/useUserStore'
import { useRouterStore } from '@/pinia/useRouterStore.js'

const register = async (app) => {
  for (const iconName in ElIconModules) {
    app.component(iconName, ElIconModules[iconName])
  }

  // 注册js加载方法
  app.config.globalProperties.$loadJS = loadJS

  // 注册服务端模版加载方法
  app.config.globalProperties.$loadTMPL = loadTMPL

  // 注册全局Element消息
  app.config.globalProperties.$message = ElMessage
  app.config.globalProperties.$messageBox = ElMessageBox

  // 注册http请求方法
  app.config.globalProperties.$http = http

  app.config.globalProperties.$useUserStore = useUserStore

  app.config.globalProperties.$useRouterStore = useRouterStore

  // 注册全局组件
  const MenuTree = await importView('views/main/menuTree.vue')()
  app.component('MenuTree', MenuTree.default || MenuTree)
  const FileUpload = await importView('views/util/fileUpload.vue')()
  app.component('FileUpload', FileUpload.default || FileUpload)
  const FileManagerBox = await importView('views/file/fileManagerBox.vue')()
  app.component('FileManagerBox', FileManagerBox.default || FileManagerBox)
}

export default {
  install: (app) => {
    register(app)
  }
}
