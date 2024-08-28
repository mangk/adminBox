import http from '@/utils/requester.js'

export const apiPage = (page = 1, page_size = 20, query = {}) => {
  return http({
    url: '/sys/setting/api/page',
    method: 'POST',
    data: { page: page, page_size: page_size, query: query }
  })
}

export const apiDetail = (id) => {
  return http({
    url: '/sys/setting/api/getById',
    method: 'POST',
    data: { id: id }
  })
}

export const apiCreate = (data) => {
  return http({
    url: '/sys/setting/api',
    method: 'POST',
    data
  })
}

export const apiUpdate = (data) => {
  return http({
    url: '/sys/setting/api',
    method: 'PUT',
    data
  })
}

export const apiDelete = (id) => {
  return http({
    url: '/sys/setting/api',
    method: 'DELETE',
    data: { id: id }
  })
}
