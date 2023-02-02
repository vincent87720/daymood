package routers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/vincent87720/daymood/app/internal/model"
	"github.com/vincent87720/daymood/app/internal/settings"
)

func SetupRouters(db *sql.DB, s settings.Settings) (*gin.Engine, error) {

	router := gin.Default()
	// router.Use(CORSMiddleware())
	router.Use(SetSession(s))
	routerGroup := router.Group("")
	routerGroup.Use(SetSession(s))
	routerGroup.Use(AuthSession())

	SetupSupplierRouters(routerGroup, db, s)
	SetupPurchaseRouters(routerGroup, db, s)
	SetupPurchaseDetailRouters(routerGroup, db, s)
	SetupProductRouters(routerGroup, db, s)
	SetupDeliveryOrderRouters(routerGroup, db)
	SetupDeliveryOrderDetailRouters(routerGroup, db, s)
	SetupDiscountRouters(routerGroup, db, s)
	SetupReportRouters(routerGroup, db)
	SetupUserRouters(routerGroup, db, s)
	SetupSettingsRouters(routerGroup, db, &s)
	SetupSystemConfigRouters(routerGroup, s)
	SetupAuthRouters(router, db)

	// exePath := s.GetExeFilePath()
	// router.Static("/daymood", exePath+"/daymoodui")
	// router.NoRoute(func(c *gin.Context) {
	// 	c.File("./daymoodui/index.html")
	// })

	return router, nil
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8001")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Cookie")
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
var returnError = func(err error) gin.H {
	return gin.H{
		"status":  "FAIL",
		"role":    "router",
		"code":    3,
		"message": err.Error(),
	}
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
var validationError = func() gin.H {
	return gin.H{
		"status":  "FAIL",
		"role":    "router",
		"code":    4,
		"message": "validation error",
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
