package routers

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vincent87720/daymood/app/internal/settings"
)

func SetupSystemConfigRouters(router *gin.Engine, s settings.Settings) (*gin.Engine, error) {

	router.GET("/systemConfigs", GetSystemConfigsHandler(s))

	return router, nil
}

func GetSystemConfigsHandler(s settings.Settings) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		sysConf, err := ioutil.ReadFile(s.GetExeFilePath() + "/assets/systemConfigs.json")
		if err != nil {
			log.Fatalf("ERROR: %v", err)
			return
		}
		context.Data(http.StatusOK, "application/json", sysConf)
		return
	}

	return gin.HandlerFunc(fn)
}
