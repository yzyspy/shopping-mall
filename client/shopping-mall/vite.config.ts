import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import {createSvgIconsPlugin} from 'vite-plugin-svg-icons'
import path from 'path'

function resolve(dir) {
    return path.join(__dirname, '.', dir)
}

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [vue(),
        createSvgIconsPlugin({
            iconDirs: [path.resolve(process.cwd(), 'src/icons/svy')],
            symbolId: 'icon-[dir]-[name]'
        })
    ]
})