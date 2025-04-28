import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';

export default defineConfig({
  plugins: [vue()],
  server: {
    port: 5173,
    open: true,
    proxy: {
      '/login': 'http://localhost:8080',
      '/register': 'http://localhost:8080',
      '/message': 'http://localhost:8080',
    }
  }
});
