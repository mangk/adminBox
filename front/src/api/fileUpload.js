import http from '@/utils/requester.js'

const adminx = window.adminX ? window.adminX : {}
const prefix = adminx.BackendPrefix ? '/' + adminx.BackendPrefix : ''

export const fileUploadCfg = () => {
  return http({
    url: prefix + '/sys/fileUpload/cfg',
    method: 'GET'
  })
}

export const fileUploadPage = (page = 1, page_size = 20, query = {}) => {
  return http({
    url: prefix + '/sys/fileUpload/page',
    method: 'POST',
    data: { page: page, page_size: page_size, query: query }
  })
}

export const fileDelete = (id) => {
  return http({
    url: prefix + '/sys/fileUpload',
    method: 'DELETE',
    data: { id: id }
  })
}
