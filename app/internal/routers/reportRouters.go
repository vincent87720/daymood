package routers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	usecases "github.com/vincent87720/daymood/app/internal/usecases"
)

func SetupReportRouters(router *gin.RouterGroup, db *sql.DB) (*gin.RouterGroup, error) {

	router.GET("/api/reports/balances", GetBalancesHandler(db))

	return router, nil
}

func GetBalancesHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		balance := usecases.NewBalance()
		balanceXi, modelErr := balance.Read(db)
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
