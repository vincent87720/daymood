package routers

import (
	"database/sql"
	"fmt"
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
		y, err := s.GetTradingSettings()
		if err != nil {
			fmt.Println(err)
			context.JSON(http.StatusBadRequest, gin.H{
				"status": "FAIL",
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"trading": y.Trading,
		})
		return
	}

	return gin.HandlerFunc(fn)
}

func PutTradingsHandler(db *sql.DB, s *settings.Settings) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		ajeossi := context.PostForm("ajeossi")
		shippingFee := context.PostForm("shippingFee")
		exchangeRate := context.PostForm("exchangeRate")
		tariff := context.PostForm("tariff")
		markup := context.PostForm("markup")

		y, err := s.GetTradingSettings()
		if err != nil {
			fmt.Println(err)
			context.JSON(http.StatusBadRequest, gin.H{
				"status": "FAIL",
			})
			return
		}

		ajeossiVal, err := checkNum(context, ajeossi)
		if err != nil {
			return
		}
		y.Trading.Ajeossi = ajeossiVal

		shippingFeeVal, err := checkNum(context, shippingFee)
		if err != nil {
			return
		}
		y.Trading.ShippingFee = shippingFeeVal

		exchangeRateVal, err := checkNum(context, exchangeRate)
		if err != nil {
			return
		}
		y.Trading.ExchangeRate = exchangeRateVal

		tariffVal, err := checkNum(context, tariff)
		if err != nil {
			return
		}
		y.Trading.Tariff = tariffVal

		markupVal, err := checkNum(context, markup)
		if err != nil {
			return
		}
		y.Trading.Markup = markupVal

		err = s.SetTradingSettings(y)
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

		context.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
		return
	}

	return gin.HandlerFunc(fn)
}
