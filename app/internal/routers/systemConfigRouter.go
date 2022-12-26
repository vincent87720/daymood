package routers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/vincent87720/daymood.backend/internal/settings"
)

func SetupSystemConfigRouters(router *gin.Engine, db *sql.DB, s *settings.Settings) (*gin.Engine, error) {

	router.GET("/systemConfigs", GetSystemConfigsHandler())

	return router, nil
}

func GetSystemConfigsHandler() gin.HandlerFunc {
	fn := func(context *gin.Context) {
		// y, err := s.GetTradingSettings()
		// if err != nil {
		// 	fmt.Println(err)
		// 	context.JSON(http.StatusBadRequest, gin.H{
		// 		"status": "FAIL",
		// 	})
		// 	return
		// }

		// context.JSON(http.StatusOK, gin.H{
		// 	"status":  "OK",
		// 	"trading": y.Trading,
		// })
		// return
	}

	return gin.HandlerFunc(fn)
}
