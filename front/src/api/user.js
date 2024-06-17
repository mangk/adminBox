import http from '@/utils/requester.js'

const adminx = window.adminX ? window.adminX : {}
const prefix = adminx.BackendPrefix ? '/' + adminx.BackendPrefix : ''

export const userPage = (page = 1, page_size = 20, query = {}) => {
  return http({
    url: prefix + '/sys/setting/user/page',
    method: 'POST',
    data: { page: page, page_size: page_size, query: query }
  })
}

export const userDetail = (id) => {
  return http({
    url: prefix + '/sys/setting/user/getById',
    method: 'POST',
    data: { id: id }
  })
}

export const userCreate = (data) => {
  return http({
    url: prefix + '/sys/setting/user',
    method: 'POST',
    data
  })
}

export const userUpdate = (data) => {
  return http({
    url: prefix + '/sys/setting/user',
    method: 'PUT',
    data
  })
}

export const userDelete = (id) => {
  return http({
    url: prefix + '/sys/setting/user',
    method: 'DELETE',
    data: { id: id }
  })
}
