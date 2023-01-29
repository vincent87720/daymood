module.exports = {
  transpileDependencies: [
    'vuetify'
  ],
  publicPath: '/daymood/',
  outputDir: 'daymoodui',
  devServer: {
		proxy:{
      '/api': {
        target: 'http://app:8000/',
        changeOrigin: true,
      },
    }
	},
  pwa: {
    manifestOptions: {
      name: "Daymood",
      short_name: "Daymood",
      display: "standalone",
      theme_color: "#E2DDD3",
    },
    themeColor: "#E2DDD3",
  },
}
