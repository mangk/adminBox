import { createApp } from 'vue'
import { createPinia } from 'pinia'
import * as ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import init from '@/init'
import App from './App.vue'
import router from './router/index'

import Nprogress from 'nprogress'
import 'nprogress/nprogress.css'

import '@/assets/main.scss'

Nprogress.configure({ showSpinner: false, ease: 'ease', speed: 200 })

const app = createApp(App)

app.use(createPinia())
app.use(ElementPlus)
app.use(init)
app.use(router)

app.mount('#app')
