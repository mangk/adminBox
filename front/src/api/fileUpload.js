import http from '@/utils/requester.js'

export const fileUploadCfg = () => {
  return http({
    url: '/sys/fileUpload/cfg',
    method: 'GET'
  })
}

export const fileUploadPage = (page = 1, page_size = 20, query = {}) => {
  return http({
    url: '/sys/fileUpload/page',
    method: 'POST',
    data: { page: page, page_size: page_size, query: query }
  })
}

export const fileDelete = (id) => {
  return http({
    url: '/sys/fileUpload',
    method: 'DELETE',
    data: { id: id }
  })
}

export const fileGroupTree = () => {
  return http({
    url: '/sys/fileGroup/tree',
    method: 'GET'
  })
}