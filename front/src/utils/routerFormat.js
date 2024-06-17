const viewModules = import.meta.glob('../views/**/*.vue')
const pluginModules = import.meta.glob('../plugins/**/*.vue')

export function importView(componentPath) {
  const keys = Object.keys(viewModules)
  const matchKeys = keys.filter((key) => {
    const k = key.replace('../', '')
    return k === componentPath
  })
  const matchKey = matchKeys[0]

  return viewModules[matchKey]
}

export function importPlugin(componentPath) {
  const keys = Object.keys(pluginModules)
  const matchKeys = keys.filter((key) => {
    const k = key.replace('../', '')
    return k === componentPath
  })
  const matchKey = matchKeys[0]

  return pluginModules[matchKey]
}

export function formatRouter(serverRouterList) {
  if (serverRouterList.length) {
    serverRouterList.forEach((item) => {
      if (item.component) {
        if (item.component.split('/')[0] === 'views') {
          item.component = importView(item.component)
        } else if (item.component.split('/')[0] === 'plugins') {
          item.component = importPlugin(item.component)
        }
      }

      if (item.children && item.children.length) {
        delete item['component']
        formatRouter(item.children)
      }
    })
  }
}

export function loadBackendPrefix() {
  var prefix = 'admin'
  if (window._adminX && window._adminX.backgroundPrefix) {
    prefix = window._adminX.backgroundPrefix
  }

  return prefix
}
