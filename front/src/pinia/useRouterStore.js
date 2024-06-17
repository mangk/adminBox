import { defineStore } from 'pinia'
import { ref } from 'vue'
import { userPermission } from '@/api/auth'
import { formatRouter, importView, loadBackendPrefix } from '@/utils/routerFormat'
import router from '@/router/index.js'

const prefix = loadBackendPrefix()

export const useRouterStore = defineStore('router', () => {
  // 服务端路由是否经过了初始化
  const initialized = ref(0)
  // 服务端路由列表
  const serverRouter = ref([])
  // 加载服务端路由
  const loadServerRouter = async (refresh = false) => {
    if (refresh || initialized.value === 0) {
      initialized.value = 1
      const menuList = await userPermission()
      serverRouter.value = menuList.data.menu[0].children
      formatRouter(serverRouter.value)
    }
    router.addRoute({
      path: '/' + prefix,
      name: prefix,
      meta: { icon: 'add' },
      component: importView('views/main/styleDefault.vue'), // => import('@/views/main/styleDefault.vue'), // TODO 这里的引入改为引入所有，从中选择
      children: serverRouter.value
    })
    return serverRouter.value
  }

  return {
    initialized,
    loadServerRouter
  }
})
