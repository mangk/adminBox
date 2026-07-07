import http from '@/utils/requester'

export const menuPage = (loadSystem = false) => {
  return http({
    url: '/sys/setting/menu/page',
    method: 'POST',
    data: { loadSystem: loadSystem }
  })
}

export const menuDetail = (id) => {
  return http({
    url: '/sys/setting/menu/getById',
    method: 'POST',
    data: { id: id }
  })
}

export const menuCreate = (data) => {
  return http({
    url: '/sys/setting/menu',
    method: 'POST',
    data
  })
}

export const menuUpdate = (data) => {
  return http({
    url: '/sys/setting/menu',
    method: 'PUT',
    data
  })
}

export const menuDelete = (id) => {
  return http({
    url: '/sys/setting/menu',
    method: 'DELETE',
    data: { id: id }
  })
}
