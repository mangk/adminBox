import http from '@/utils/requester.js'

export const rolePage = (page = 1, page_size = 20, query = {}) => {
  return http({
    url: '/sys/setting/role/page',
    method: 'POST',
    data: { page: page, page_size: page_size, query: query }
  })
}

export const roleDetail = (id) => {
  return http({
    url: '/sys/setting/role/getById',
    method: 'POST',
    data: { id: id }
  })
}

export const roleCreate = (data) => {
  return http({
    url: '/sys/setting/role',
    method: 'POST',
    data
  })
}

export const roleUpdate = (data) => {
  return http({
    url: '/sys/setting/role',
    method: 'PUT',
    data
  })
}

export const roleDelete = (id) => {
  return http({
    url: '/sys/setting/role',
    method: 'DELETE',
    data: { id: id }
  })
}

export const roleAll = () => {
  return http({
    url: '/sys/setting/role/all',
    method: 'GET'
  })
}
