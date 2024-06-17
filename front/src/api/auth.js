import http from '@/utils/requester'

const adminx = window.adminX ? window.adminX : {}
const prefix = adminx.BackendPrefix ? '/' + adminx.BackendPrefix : ''

export const userPermission = () => {
  return http({
    url: prefix + '/sys/auth/userPermission',
    method: 'GET'
  })
}

export const permissionAll = () => {
  return http({
    url: prefix + '/sys/auth/permissionAll',
    method: 'GET'
  })
}

export const permissionGetByIdAndModule = (id, module) => {
  return http({
    url: prefix + '/sys/auth/permissionGetByIdAndModule',
    method: 'POST',
    data: { id: id, module: module }
  })
}

export const permissionSave = (data) => {
  return http({
    url: prefix + '/sys/auth/permissionSave',
    method: 'PUT',
    data
  })
}

export const verificationCode = () => {
  return http({
    url: prefix + '/sys/verificationCode',
    method: 'GET'
  })
}

export const login = (data) => {
  return http({
    url: prefix + '/sys/login',
    method: 'POST',
    data
  })
}
