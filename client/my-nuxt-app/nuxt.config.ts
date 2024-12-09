// nuxt.config.js
export default defineNuxtConfig({
  compatibilityDate: '2024-04-03',
  devtools: { enabled: true },
  modules: [
    '@pinia/nuxt',
  ],
  plugins: [
    '~/plugins/firebase.ts',
  ],
  build: {
    transpile: ['firebase', 'firebaseui']
  },
  app: {
    head: {
      // CDN からの読み込みを削除
      script: [],
      link: []
    }
  }
})
