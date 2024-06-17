import axios from 'axios' // 引入axios
import { emitter } from '@/utils/bus.js'
import router from '@/router/index'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useUserStore } from '@/pinia/useUserStore'

const adminx = window.adminX ? window.adminX : {}
let serverHost = adminx.RunAt ? adminx.RunAt : ''

const http = axios.create({
  baseURL: serverHost,
  timeout: 30000
})
let acitveAxios = 0
let timer
const showLoading = () => {
  acitveAxios++
  if (timer) {
    clearTimeout(timer)
  }
  timer = setTimeout(() => {
    if (acitveAxios > 0) {
      emitter.emit('showLoading')
    }
  }, 400)
}

const closeLoading = () => {
  acitveAxios--
  if (acitveAxios <= 0) {
    clearTimeout(timer)
    emitter.emit('closeLoading')
  }
}
// http request 拦截器
http.interceptors.request.use(
  (config) => {
    if (!config.donNotShowLoading) {
      showLoading()
    }

    const userStore = useUserStore()
    config.headers = {
      'Content-Type': 'application/json; charset=utf-8',
      Authorization: 'Bearer ' + userStore.token(),
      ...config.headers
    }
    return config
  },
  (error) => {
    if (!error.config.donNotShowLoading) {
      closeLoading()
    }
    ElMessage({
      showClose: true,
      message: error,
      type: 'error'
    })
    return error
  }
)

// http response 拦截器
http.interceptors.response.use(
  (response) => {
    // const userStore = useUserStore()
    // if (!response.config.donNotShowLoading) {
    //     closeLoading()
    // }
    // if (response.headers['new-token']) {
    //     userStore.setToken(response.headers['new-token'])
    // }
    if (response.data.code === 0 || response.headers.success === 'true') {
      if (response.headers.msg) {
        response.data.msg = decodeURI(response.headers.msg)
      }
      return response.data
    } else {
      ElMessage({
        showClose: true,
        message: response.data.msg || decodeURI(response.headers.msg),
        type: 'error'
      })
      if (response.data.data && response.data.data.reload) {
        // userStore.token = ''
        localStorage.clear()
        router.push({ name: 'login', replace: true })
      }
      return response.data
    }
  },
  (error) => {
    if (!error.config.donNotShowLoading) {
      closeLoading()
    }

    if (!error.response) {
      ElMessageBox.confirm(
        `
        <p>检测到请求错误</p>
        <p>${error}</p>
        `,
        '请求报错',
        {
          dangerouslyUseHTMLString: true,
          distinguishCancelAndClose: true,
          confirmButtonText: '稍后重试',
          cancelButtonText: '取消'
        }
      )
      return
    }

    switch (error.response.status) {
      case 401:
        ElMessage({
          showClose: true,
          message: error.response.data.msg,
          type: 'error'
        })
        router.push({ name: 'login', replace: true })
        break
      case 500:
        ElMessageBox.confirm(
          `
        <p>检测到接口错误${error}</p>
        <p>错误码<span style="color:red"> 500 </span>：此类错误内容常见于后台panic，请先查看后台日志，如果影响您正常使用可强制登出清理缓存</p>
        `,
          '接口报错',
          {
            dangerouslyUseHTMLString: true,
            distinguishCancelAndClose: true,
            confirmButtonText: '清理缓存',
            cancelButtonText: '取消'
          }
        ).then(() => {
          // const userStore = useUserStore()
          // userStore.token = ''
          localStorage.clear()
          router.push({ name: 'Login', replace: true })
        })
        break
      case 404:
        ElMessageBox.confirm(
          `
          <p>检测到接口错误${error}</p>
          <p>错误码<span style="color:red"> 404 </span>：此类错误多为接口未注册（或未重启）或者请求路径（方法）与api路径（方法）不符--如果为自动化代码请检查是否存在空格</p>
          `,
          '接口报错',
          {
            dangerouslyUseHTMLString: true,
            distinguishCancelAndClose: true,
            confirmButtonText: '我知道了',
            cancelButtonText: '取消'
          }
        )
        break
    }

    return error
  }
)
export default http
