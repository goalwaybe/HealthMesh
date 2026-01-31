// src/main.ts
import { createSSRApp } from 'vue'
import App from './App.vue'
import * as medi from './api/medi'

// 类型声明：扩展 Vue 实例的  $ api 属性
declare module 'vue' {
  interface ComponentCustomProperties {
     $api: typeof medi
  }
}

export function createApp() {
  const app = createSSRApp(App)

  // 注入 API（类型安全）
  app.config.globalProperties.$api = medi

  return {
    app
  }
}