import axios from 'axios' // 引入axios
import { emitter } from '@/utils/bus.js'
import router from '@/router/index'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useUserStore } from '@/pinia/useUserStore'
import { handler401 } from './401'

const http = axios.create({
  baseURL: serverHost(),
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

    const u = useUserStore().userAuth()

    config.headers = {
      'Content-Type': 'application/json; charset=utf-8',
      Authorization: 'Bearer ' + u.token,
      'X-User-Type': u.userType,
      'X-User-Id': u.userId,

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
    if (
      response.headers['content-type'] == 'application/json; charset=utf-8' &&
      response.data.code != 0
    ) {
      if (response.data.code == 401) {
        localStorage.clear()
        if (window.adminBox?.RunAt) {
          window.location.href = window.adminBox.RunAt
          return
        }
        handler401(true)
        return
      }
      ElMessage({
        showClose: true,
        message: response.data.msg,
        type: 'error'
      })
    }
    return response.data
  },
  (error) => {
    closeLoading()

    switch (error) {
      case 401:
        handler401(true)
        break
      default:
        // ElMessageBox.confirm(
        //   `
        //   <p>检测到接口错误${error}</p>
        //   <p>错误码<span style="color:red"> ${error.response.status} </span>：此类错误多为接口未注册（或未重启）或者请求路径（方法）与api路径（方法）不符--如果为自动化代码请检查是否存在空格</p>
        //   `,
        //   '接口报错',
        //   {
        //     dangerouslyUseHTMLString: true,
        //     distinguishCancelAndClose: true,
        //     confirmButtonText: '我知道了',
        //     cancelButtonText: '取消'
        //   }
        // )
        ElMessage({
          showClose: true,
          message: error,
          type: 'error'
        })
        break
    }

    return error
  }
)
export default http

export function serverHost() {
  let str = []
  if (window.adminBox.RunAt) {
    str.push(window.adminBox.RunAt)
  }
  if (window.adminBox.BackendRouterPrefix) {
    str.push(window.adminBox.BackendRouterPrefix)
  }

  return str.join('/')
}
