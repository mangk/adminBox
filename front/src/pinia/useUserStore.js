import { login } from '@/api/auth'
import router from '@/router'
import { loadBackendPrefix } from '@/utils/routerFormat'
import { defineStore } from 'pinia'
import { ref } from 'vue'
import { useRouterStore } from './useRouterStore'

export const useUserStore = defineStore('user', () => {
  const initialized = ref(0)
  const user = ref({})
  const _tokenStorageKey = 'token'

  const token = () => {
    var token = localStorage.getItem(_tokenStorageKey)
    if (!token) {
      router.replace({ name: 'login' })
      return
    }
    return token
  }

  const userInfo = (refresh = false) => {
    if (refresh || initialized.value === 0) {
      initialized.value = 1
    }
    return user.value
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

    const routerStore = useRouterStore()
    await routerStore.loadServerRouter(true)

    return res.data.user_config.home_page
      ? res.data.user_config.home_page
      : `/${backendPrefix}/welcome`
  }

  const logOut = () => {
    localStorage.removeItem(_tokenStorageKey)
    router.replace({ name: 'login' })
  }

  return {
    initialized,
    userInfo,
    token,
    logIn,
    logOut
  }
})
