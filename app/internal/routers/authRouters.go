package routers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vincent87720/daymood/app/internal/model"
	"github.com/vincent87720/daymood/app/internal/settings"
	"github.com/vincent87720/daymood/app/internal/usecases"
)

func SetupAuthRouters(router *gin.Engine, db *sql.DB, s settings.Settings) (*gin.Engine, error) {

	router.POST("/api/login", UserLoginHandler(db, s))
	router.POST("/api/logout", UserLogoutHandler(db, s))

	return router, nil
}

func UserLoginHandler(db *sql.DB, s settings.Settings) gin.HandlerFunc {
	fn := func(context *gin.Context) {

		userModel := &model.User{}

		err := context.BindJSON(&userModel)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		if checkEmpty(userModel.Username) == true {
			context.JSON(http.StatusBadRequest, emptyError("username"))
			return
		}

		if checkEmpty(userModel.Password) == true {
			context.JSON(http.StatusBadRequest, emptyError("password"))
			return
		}

		user := usecases.NewUser(userModel)
		valid, err := usecases.Login(user, db, s)
		if err != nil || valid == false {
			context.JSON(http.StatusBadRequest, validationError())
			return
		}
		SaveSession(context, user.Model.ID)

		context.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"isLogin": true,
		})
		return
	}

	return gin.HandlerFunc(fn)
}

func UserLogoutHandler(db *sql.DB, s settings.Settings) gin.HandlerFunc {
	fn := func(context *gin.Context) {

		ClearSession(context)

		context.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"isLogin": false,
		})
		return
	}

	return gin.HandlerFunc(fn)
}
