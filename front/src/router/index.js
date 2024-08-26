import { createRouter, createWebHashHistory } from 'vue-router'
import Nprogress from 'nprogress'
import { useRouterStore } from '@/pinia/useRouterStore.js'
import { loadBackendPrefix } from '@/utils/routerFormat.js'
import { useUserStore } from '@/pinia/useUserStore'

const prefix = loadBackendPrefix()

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: '/' + `${prefix}/welcome`
    },
    // 后台管理页面
    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/login.vue')
    },
    {
      path: '/' + prefix,
      name: prefix,
      component: () => import('@/views/main/styleDefault.vue')
    },
    {
      path: '/:catchAll(.*)',
      meta: {
        closeTab: true
      },
      component: () => import('@/views/404.vue')
    }
  ]
})

router.beforeEach(async (to, from) => {
  Nprogress.start()

  if (to.meta.title) {
    window.document.title = to.meta.title
  }

  if (to.path === '/login') {
    return true
  }

  const routerStore = useRouterStore()
  const userStore = useUserStore()
  if (!routerStore.initialized) {
    await routerStore.loadServerRouter(true)
    return { path: to.path }
  }

  if (!userStore.isLogIn()) {
    userStore.logOut()
  }

  return true
})

router.afterEach(() => {
  Nprogress.done()
})

router.onError(() => {
  // 路由发生错误后销毁进度条
  Nprogress.done()
})

export default router
