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
	}
}
