// vite.config.ts
import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
const path = require("path");


// https://vitejs.dev/config/
export default defineConfig({

  plugins: [vue()],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "src"),
      "comps": path.resolve(__dirname, "src/components"),
      "assets": path.resolve(__dirname, "src/assets"),
      "apis": path.resolve(__dirname, "src/apis"),
      "views": path.resolve(__dirname, "src/views"),
      "utils": path.resolve(__dirname, "src/utils"),
      "style": path.resolve(__dirname, "src/style"),
      "routers": path.resolve(__dirname, "src/routers"),
      "store": path.resolve(__dirname, "src/store"),
    },
  },
});