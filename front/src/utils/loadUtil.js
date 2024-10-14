import { loadModule } from 'vue3-sfc-loader'
import * as Vue from 'vue'
import { defineAsyncComponent, markRaw } from 'vue'
import http from './requester'


export function loadJS(url, callback = null) {
    // 检查是否已经加载过该脚本
    if (document.querySelector(`script[src="${url}"]`)) {
        if (typeof callback === 'function') {
            callback()
        }
        return
    }

    const script = document.createElement('script')
    script.type = 'text/javascript'
    script.src = url

    // 当脚本加载完成后执行回调
    script.onload = () => {
        if (typeof callback === 'function') {
            callback()
        }
    }

    // 将脚本插入到页面中
    document.head.appendChild(script)
}

export function loadTMPL(url, name = 'myConvert') {
    const options = {
        moduleCache: { vue: Vue },
        getFile() {
            return http(url, {
                method: 'GET'
            }).then((response) => {
                return response
            })
        },
        addStyle(styleString) {
            const style = document.createElement('style')
            style.setAttribute('id', name)
            style.textContent = styleString
            const ref = document.head.getElementsByTagName('style')[0] || null
            document.head.insertBefore(style, ref)
        }
    }

    const component = defineAsyncComponent(() => loadModule(`${name}.vue`, options))
    return markRaw(component)
}