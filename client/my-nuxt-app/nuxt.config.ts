export default {
  // その他の設定
  router: {
    extendRoutes(routes: Array<any>, resolve: Function) {
      routes.push({
        name: 'home',
        path: '/',
        component: resolve(__dirname, 'pages/Toppage.vue')
      })
    }
  }
}