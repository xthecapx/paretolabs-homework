import { defineConfig, loadEnv } from "vite";
import react from "@vitejs/plugin-react";
import viteTsconfigPaths from "vite-tsconfig-paths";
import type { ConfigEnv } from "vite";
import { NodeGlobalsPolyfillPlugin } from "@esbuild-plugins/node-globals-polyfill";
import { NodeModulesPolyfillPlugin } from "@esbuild-plugins/node-modules-polyfill";
import rollupNodePolyFill from "rollup-plugin-node-polyfills";
import { splitVendorChunkPlugin } from "vite";

// https://vitejs.dev/config/
export default (configEnv: ConfigEnv) => {
  process.env = {
    ...process.env,
    ...loadEnv(configEnv.mode, process.cwd(), ""),
  };

  return defineConfig({
    plugins: [
      react(),
      viteTsconfigPaths(),
      splitVendorChunkPlugin(),
    ],
    build: {
      outDir: "build",
      rollupOptions: {
        plugins: [
          // Enable rollup polyfills plugin
          // used during production bundling
          rollupNodePolyFill(),
        ],
      },
    },
    server: {
      open: false,
      port: 3000,
      host: true,
      proxy: {
        "/apiURL": {
          target: 'http://localhost:1234',
          changeOrigin: true,
          secure: false,
          rewrite: (path) => path.replace(/^\/apiURL/, ""),
        },
      },
    },
    optimizeDeps: {
      esbuildOptions: {
        // Node.js global to browser globalThis
        define: {
          global: "globalThis",
        },
        // Enable esbuild polyfill plugins
        plugins: [
          NodeGlobalsPolyfillPlugin({
            process: true,
            buffer: true,
          }),
          NodeModulesPolyfillPlugin(),
        ],
      },
    },
  });
};
