import {defineConfig} from 'vite';
import react from '@vitejs/plugin-react';
import path from 'path';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      src: '/src',
      shadcn: path.resolve(__dirname, 'src/@/components/ui'),
      containers: path.resolve(__dirname, '/src/containers'),
    },
  },
});
