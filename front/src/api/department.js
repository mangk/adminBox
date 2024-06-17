import http from '@/utils/requester'

const adminx = window.adminX ? window.adminX : {}
const prefix = adminx.BackendPrefix ? '/' + adminx.BackendPrefix : ''

export const departmentPage = () => {
  return http({
    url: prefix + '/sys/setting/department/page',
    method: 'POST'
  })
}

export const departmentDetail = (id) => {
  return http({
    url: prefix + '/sys/setting/department/getById',
    method: 'POST',
    data: { id: id }
  })
}

export const departmentCreate = (data) => {
  return http({
    url: prefix + '/sys/setting/department',
    method: 'POST',
    data
  })
}

export const departmentUpdate = (data) => {
  return http({
    url: prefix + '/sys/setting/department',
    method: 'PUT',
    data
  })
}

export const departmentDelete = (id) => {
  return http({
    url: prefix + '/sys/setting/department',
    method: 'DELETE',
    data: { id: id }
  })
}
