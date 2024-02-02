import { defineConfig } from 'vite';
import solidPlugin from 'vite-plugin-solid';

export default defineConfig({
  plugins: [
    solidPlugin(),
  ],
  server: {
    port: 3000,
    watch: {
      usePolling: true,
      interval: 100
    }
  },
  build: {
    target: 'esnext',
    outDir: '../dist/src',
    assetsDir: './assets',
    emptyOutDir: true,
  },
});
