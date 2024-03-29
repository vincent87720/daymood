package routers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vincent87720/daymood/app/internal/model"
	"github.com/vincent87720/daymood/app/internal/usecases"
)

func SetupAuthRouters(router *gin.Engine, db *sql.DB) *gin.Engine {

	router.POST("/api/login", UserLoginHandler(db))
	router.POST("/api/logout", UserLogoutHandler(db))

	return router
}

func UserLoginHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {

		userModel := &model.User{}

		err := context.BindJSON(&userModel)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		checkList := []Field{
			{Key: "Username", Val: userModel.Username},
			{Key: "Password", Val: userModel.Password},
		}
		err = checkEmpty(checkList)
		if err != nil {
			context.JSON(http.StatusBadRequest, emptyError(err))
			return
		}

		user := usecases.NewUser(userModel)
		valid, err := usecases.Login(user, db)
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

func UserLogoutHandler(db *sql.DB) gin.HandlerFunc {
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
