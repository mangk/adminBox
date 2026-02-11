import http from '@/utils/requester'

export const userPermission = () => {
  return http({
    url: '/sys/auth/userPermission',
    method: 'GET'
  })
}

export const permissionAll = () => {
  return http({
    url: '/sys/auth/permissionAll',
    method: 'GET'
  })
}

export const permissionGetByIdAndModule = (id, module) => {
  return http({
    url: '/sys/auth/permissionGetByIdAndModule',
    method: 'POST',
    data: { id: id, module: module }
  })
}

export const permissionSave = (data) => {
  return http({
    url: '/sys/auth/permissionSave',
    method: 'PUT',
    data
  })
}

export const verificationCode = () => {
  return http({
    url: '/sys/verificationCode',
    method: 'GET'
  })
}

export const login = (data) => {
  return http({
    url: '/sys/login',
    method: 'POST',
    data
  })
}

export const IsRewriteIndex = () => {
  return http({
    url: '/sys/isRewriteIndex',
    method: 'GET'
  })
}
