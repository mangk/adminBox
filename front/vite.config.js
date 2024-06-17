import {fileURLToPath, URL} from 'node:url'

import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
    base: './', // index.html文件所在位置
    root: './', // js导入的资源路径，src
    plugins: [
        vue(),
    ],
    resolve: {
        alias: {
            '@': fileURLToPath(new URL('./src', import.meta.url))
        }
    },
    build: {
        // minify: 'terser', // 是否进行压缩,boolean | 'terser' | 'esbuild',默认使用terser
        // terserOptions: {
        //     compress: {
        //         drop_console: true,
        //         drop_debugger: true,
        //     },
        // },
        manifest: false, // 是否产出manifest.json
        sourcemap: false, // 是否产出sourcemap.json
        outDir: 'dist', // 产出目录
        cssCodeSplit: false,
        rollupOptions: {
            output: {
                // 用于从入口点创建的块的打包输出格式[name]表示文件名,[hash]表示该文件内容hash值
                entryFileNames: 'assets/X[name][hash].js', // TODO 这里的文件命名
                // 用于命名代码拆分时创建的共享块的输出命名
                chunkFileNames: 'assets/X[name][hash].js', // TODO 这里的文件命名
            }
        },
        // rollupOptions,
    },
})
