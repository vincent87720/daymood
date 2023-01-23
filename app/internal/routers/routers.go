package routers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/vincent87720/daymood/app/internal/model"
	"github.com/vincent87720/daymood/app/internal/settings"
)

func SetupRouters(db *sql.DB, s settings.Settings) (*gin.Engine, error) {

	router := gin.Default()
	router.Use(CORSMiddleware())

	SetupSupplierRouters(router, db, s)
	SetupPurchaseRouters(router, db, s)
	SetupPurchaseDetailRouters(router, db, s)
	SetupProductRouters(router, db, s)
	SetupDeliveryOrderRouters(router, db)
	SetupDeliveryOrderDetailRouters(router, db, s)
	SetupDiscountRouters(router, db, s)
	SetupReportRouters(router, db)
	SetupSettingsRouters(router, db, &s)
	SetupSystemConfigRouters(router, s)

	exePath := s.GetExeFilePath()
	router.Static("/daymood", exePath+"/daymoodui")

	return router, nil
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func checkEmpty(s string) bool {
	if s == "" || len(s) <= 0 {
		return true
	}
	return false
}

var generalError = gin.H{
	"status": "FAIL",
	"role":   "router",
	"code":   0,
}
var emptyError = func(varName string) gin.H {
	return gin.H{
		"status":  "FAIL",
		"role":    "router",
		"code":    1,
		"message": varName + " field should not be empty",
	}
}
var typeError = func(varName string) gin.H {
	return gin.H{
		"status":  "FAIL",
		"role":    "router",
		"code":    2,
		"message": "Invalid type: " + varName,
	}
}
var modelError = func(modelErr *model.ModelError) gin.H {
	return gin.H{
		"status":  "FAIL",
		"role":    "model",
		"model":   modelErr.Model,
		"code":    modelErr.Code,
		"message": modelErr.Message,
	}
}
