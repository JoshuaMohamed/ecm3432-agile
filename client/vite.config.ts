import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

// https://vite.dev/config/
export default defineConfig({
  server: {
    port: 3000,
    strictPort: true,
    watch: {
      usePolling: true,
      interval: 300,
    },
  },
  plugins: [react()],
});
