package routers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vincent87720/daymood/app/internal/settings"
)

func SetupSettingsRouters(router *gin.Engine, db *sql.DB, s *settings.Settings) (*gin.Engine, error) {

	router.GET("/tradings", GetTradingsHandler(db, s))
	router.PUT("/tradings", PutTradingsHandler(db, s))

	return router, nil
}

func GetTradingsHandler(db *sql.DB, s *settings.Settings) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		trading, err := s.GetTradingSettings()
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"status": "FAIL",
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"trading": trading,
		})
		return
	}

	return gin.HandlerFunc(fn)
}

func PutTradingsHandler(db *sql.DB, s *settings.Settings) gin.HandlerFunc {
	fn := func(context *gin.Context) {

		var trading settings.Trading

		err := context.BindYAML(&trading)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		err = s.SetTradingSettings(trading)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"status": "FAIL",
			})
			return
		}

		err = s.MarshalSettings()
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"status": "FAIL",
			})
			return
		}

		err = s.WriteFile()
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"status": "FAIL",
			})
			return
		}

		trading, err = s.GetTradingSettings()
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"status": "FAIL",
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"trading": trading,
		})
		return
	}

	return gin.HandlerFunc(fn)
}
