import { login } from '@/api/auth'
import router from '@/router'
import { loadBackendPrefix } from '@/utils/routerFormat'
import { defineStore } from 'pinia'
import { ref } from 'vue'
import { useRouterStore } from './useRouterStore'

export const useUserStore = defineStore('user', () => {
  const initialized = ref(0)
  const user = ref({})
  const _tokenStorageKey = 'x-token'
  const _userTypeStroageKey = 'x-user-type'
  const _userIdStroageKey = 'x-user-id'

  const token = () => {
    var token = localStorage.getItem(_tokenStorageKey)
    if (!token) {
      router.replace({ name: 'login' })
      return
    }
    return token
  }

  const userType = () => {
    return localStorage.getItem(_userTypeStroageKey)
  }

  const userId = () => {
    return localStorage.getItem(_userIdStroageKey)
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
      return ''
    }

    user.value = res.data
    initialized.value = 1
    await localStorage.setItem(_tokenStorageKey, res.data.jwt_token)
    await localStorage.setItem(_userTypeStroageKey, 'default')
    await localStorage.setItem(_userIdStroageKey, res.data.id)

    const routerStore = useRouterStore()
    await routerStore.loadServerRouter(true)

    return res.data.user_config.home_page
      ? res.data.user_config.home_page
      : `/${backendPrefix}/welcome`
  }

  const logOut = () => {
    localStorage.clear()
    router.replace({ name: 'login' })
  }

  return {
    initialized,
    userInfo,
    setUserData,
    token,
    userType,
    userId,
    logIn,
    logOut
  }
})
