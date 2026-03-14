import path from 'path'
import { fileURLToPath } from 'url'
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import tailwindcss from '@tailwindcss/vite'

const __dirname = path.dirname(fileURLToPath(import.meta.url))

export default defineConfig({
  plugins: [vue(), tailwindcss()],
  optimizeDeps: {
    include: ['@heroicons/vue/24/outline'],
  },
  resolve: {
    alias: {
      '@ui': path.resolve(__dirname, '../packages/ui'),
      '@heroicons/vue': path.resolve(__dirname, 'node_modules/@heroicons/vue'),
      'markdown-it': path.resolve(__dirname, 'node_modules/markdown-it'),
    },
  },
  base: process.env.NODE_ENV === 'production' ? '/' : '/',
  server: {
    port: 5173,
    host: '0.0.0.0',
    origin: process.env.VITE_ORIGIN || 'http://localhost:5173',
    proxy: {
      '/api': { target: process.env.VITE_API_TARGET || 'http://localhost:8080', changeOrigin: true },
      '/storage': { target: process.env.VITE_API_TARGET || 'http://localhost:8080', changeOrigin: true },
    },
  },
})
