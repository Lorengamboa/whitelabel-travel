import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  base: './',
  build: {
    assetsDir: './',
  },
  server: {
    proxy: {
      "/api": {
        target: "http://localhost:8080/",
        changeOrigin: true,
        secure: false,
        rewrite: (path) => path.replace(/^\/api/, ""),
      }
    }
  },
  plugins: [react()],
  resolve: {
    alias: {
      src: "/src",
      "@/store": "/src/store",
      "@/services": "/src/services",
      "@/config": "/src/config",
      "@/containers": "/src/containers",
      "@/utils": "/src/utils",
    }
  }
})
