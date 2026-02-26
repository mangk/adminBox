import router from '@/router'
import { ElMessage } from 'element-plus'

export const handler401 = (back = false) => {
  var whiteList = ['/', '/login']
  var redirect = ''
  if (whiteList.indexOf(router.currentRoute.value.path) == -1) {
    // 不在白名单时
    ElMessage({
      showClose: true,
      message: '身份认证失败',
      type: 'error'
    })

    redirect = encodeURIComponent(router.currentRoute.value.fullPath)
  }

  if (router.currentRoute.value.query.redirect) {
    redirect = router.currentRoute.value.query.redirect
  }

  let query = { path: '/login' }
  // if (redirect) {
  //   query.query = { redirect: redirect }
  // }
  router.replace(query)
}
