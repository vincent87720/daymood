package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vincent87720/daymood/app/internal/settings"
)

func SetupSystemConfigRouters(router *gin.RouterGroup, s *settings.Settings) *gin.RouterGroup {

	router.GET("/api/systemConfigs", GetSystemConfigsHandler(s))

	return router
}

func GetSystemConfigsHandler(s *settings.Settings) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		sysConf := s.GetSystemConfigs()
		context.Data(http.StatusOK, "application/json", sysConf)
		return
	}

	return gin.HandlerFunc(fn)
}
