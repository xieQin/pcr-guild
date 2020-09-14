module.exports = {
  // 修改的配置
  // 将baseUrl: '/api',改为baseUrl: '/',
  // baseUrl: '/',
  devServer: {
    proxy: {
      '/api': {
        target: 'http://localhost:9099',
        changeOrigin: true,
        ws: true
      }
    }
  }
}