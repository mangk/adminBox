import { login } from '@/api/auth'
import router from '@/router'
import { loadBackendPrefix } from '@/utils/routerFormat'
import { defineStore } from 'pinia'
import { ref } from 'vue'
import { useRouterStore } from './useRouterStore'
import { ElMessage } from 'element-plus'
import { handler401 } from '@/utils/401'

export const useUserStore = defineStore('user', () => {
  const initialized = ref(0)
  const user = ref({})
  const _tokenStorageKey = 'x-token'
  const _userTypeStroageKey = 'x-user-type'
  const _userIdStroageKey = 'x-user-id'

  const userAuth = () => {
    let v1 = localStorage.getItem(_tokenStorageKey)
    let v2 = localStorage.getItem(_userTypeStroageKey)
    let v3 = localStorage.getItem(_userIdStroageKey)
    if (v1 && v2 && v3) {
      return {
        token: v1,
        userType: v2,
        userId: v3
      }
    } else {
      handler401(true)
      return false
    }
  }

  const userInfo = (refresh = false) => {
    if (refresh || initialized.value === 0) {
      initialized.value = 1
    }
    return user.value
  }

  const setUserData = (u) => {
    user.value = u
  }

  const logIn = async (form) => {
    const backendPrefix = loadBackendPrefix()
    var res = await login(form)
    if (res.code != 0) {
      ElMessage.error(res.msg)
      return ''
    }

    user.value = res.data
    initialized.value = 1
    await localStorage.setItem(_tokenStorageKey, res.data.jwt_token)
    await localStorage.setItem(_userTypeStroageKey, 'default')
    await localStorage.setItem(_userIdStroageKey, res.data.id)

    const routerStore = useRouterStore()
    await routerStore.loadServerRouter(true)

    if (router.currentRoute.value.query.redirect) {
      return decodeURIComponent(router.currentRoute.value.query.redirect)
    }

    return res.data.user_config.home_page
      ? res.data.user_config.home_page
      : `/${backendPrefix}/welcome`
  }

  const isLogIn = () => {
    return (
      localStorage.getItem(_tokenStorageKey) &&
      localStorage.getItem(_userTypeStroageKey) &&
      localStorage.getItem(_userIdStroageKey)
    )
  }

  const logOut = () => {
    localStorage.clear()
    window.location.href = window.location.origin
    // router.replace({ name: 'login' })
  }

  return {
    initialized,
    userInfo,
    setUserData,
    logIn,
    logOut,
    isLogIn,
    userAuth
  }
})
