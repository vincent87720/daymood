package routers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vincent87720/daymood/app/internal/settings"
	usecases "github.com/vincent87720/daymood/app/internal/usecases"
)

func SetupSettingsRouters(router *gin.RouterGroup, db *sql.DB, s *settings.Settings) (*gin.RouterGroup, error) {

	router.GET("/api/tradings", GetTradingsHandler(db, s))
	router.PUT("/api/tradings", PutTradingsHandler(db, s))

	return router, nil
}

func GetTradingsHandler(db *sql.DB, s *settings.Settings) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		trading := usecases.NewTradings(s)
		tradingData, err := trading.Read()
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"status": "FAIL",
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"trading": tradingData,
		})
		return
	}

	return gin.HandlerFunc(fn)
}

func PutTradingsHandler(db *sql.DB, s *settings.Settings) gin.HandlerFunc {
	fn := func(context *gin.Context) {

		var tradingData *settings.Trading

		err := context.BindYAML(&tradingData)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		trading := usecases.NewTradings(s)

		err = trading.Update(*tradingData)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"status": "FAIL",
			})
		}

		tradingData, err = trading.Read()
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"status": "FAIL",
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"trading": tradingData,
		})
		return
	}

	return gin.HandlerFunc(fn)
}
