package routers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vincent87720/daymood/app/internal/model"
	"github.com/vincent87720/daymood/app/internal/settings"
)

func SetupReportRouters(router *gin.Engine, db *sql.DB, s settings.Settings) (*gin.Engine, error) {

	router.GET("/reports/balances", GetBalancesHandler(db))

	return router, nil
}

func GetBalancesHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		balanceXi, modelErr := model.GetBalances(db)
		if modelErr != nil {
			context.JSON(http.StatusBadRequest, modelError(modelErr))
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"records": balanceXi,
		})
		return
	}

	return gin.HandlerFunc(fn)
}
