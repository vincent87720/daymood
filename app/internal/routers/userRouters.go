package routers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vincent87720/daymood/app/internal/model"
	"github.com/vincent87720/daymood/app/internal/settings"
	"github.com/vincent87720/daymood/app/internal/usecases"
)

type PatchUserPasswordReqBody struct {
	ID          int64
	Username    string
	OldPassword string
	NewPassword string
}

func SetupUserRouters(router *gin.RouterGroup, db *sql.DB, s settings.Settings) (*gin.RouterGroup, error) {

	router.GET("/api/users", GetUsersHandler(db))
	router.GET("/api/users/:id", GetUserHandler(db))
	router.POST("/api/users", PostUserHandler(db))
	router.PUT("/api/users/:id", PutUserHandler(db))
	router.PATCH("/api/users/:id/passwords", PatchUserPasswordHandler(db))
	router.DELETE("/api/users/:id", DeleteUserHandler(db))

	return router, nil
}

func GetUsersHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		userModel := &model.User{}

		user := usecases.NewUser(userModel)
		userXi, modelErr := usecases.ReadAll(user, db)
		if modelErr != nil {
			context.JSON(http.StatusBadRequest, modelError(modelErr))
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"records": userXi,
		})
		return
	}

	return gin.HandlerFunc(fn)
}

func GetUserHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		userID := context.Param("id")

		if checkEmpty(userID) == true {
			context.JSON(http.StatusBadRequest, emptyError("id"))
			return
		}

		userIDVal, err := strconv.ParseInt(userID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError("id"))
			return
		}
		userModel := &model.User{
			ID: userIDVal,
		}

		user := usecases.NewUser(userModel)
		userXi, modelErr := usecases.Read(user, db)
		if modelErr != nil {
			context.JSON(http.StatusBadRequest, modelError(modelErr))
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"records": userXi,
		})
		return
	}

	return gin.HandlerFunc(fn)
}

func PostUserHandler(db *sql.DB) gin.HandlerFunc {
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
		modelErr := usecases.Create(user, db)
		if modelErr != nil {
			context.JSON(http.StatusBadRequest, modelError(modelErr))
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
		return
	}

	return gin.HandlerFunc(fn)
}

func PutUserHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		userID := context.Param("id")

		if checkEmpty(userID) == true {
			context.JSON(http.StatusBadRequest, emptyError("id"))
			return
		}

		userIDVal, err := strconv.ParseInt(userID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError("id"))
			return
		}

		userModel := &model.User{}

		err = context.BindJSON(&userModel)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		if checkEmpty(userModel.Name) == true {
			context.JSON(http.StatusBadRequest, emptyError("name"))
			return
		}

		userModel.ID = userIDVal

		user := usecases.NewUser(userModel)
		modelErr := usecases.Update(user, db)
		if modelErr != nil {
			context.JSON(http.StatusBadRequest, modelError(modelErr))
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
		return
	}

	return gin.HandlerFunc(fn)
}

func PatchUserPasswordHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		userID := context.Param("id")

		if checkEmpty(userID) == true {
			context.JSON(http.StatusBadRequest, emptyError("id"))
			return
		}

		userIDVal, err := strconv.ParseInt(userID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError("id"))
			return
		}

		passwordReq := &PatchUserPasswordReqBody{}

		err = context.BindJSON(&passwordReq)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError(err.Error()))
			return
		}

		userModel := &model.User{
			ID:       userIDVal,
			Username: passwordReq.Username,
		}

		user := usecases.NewUser(userModel)
		err = user.UpdatePassword(db, passwordReq.OldPassword, passwordReq.NewPassword)
		if err != nil {
			context.JSON(http.StatusBadRequest, validationError())
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
		return
	}

	return gin.HandlerFunc(fn)
}

func DeleteUserHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {

		userID := context.Param("id")

		if checkEmpty(userID) == true {
			context.JSON(http.StatusBadRequest, emptyError("id"))
			return
		}

		userIDVal, err := strconv.ParseInt(userID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, typeError("id"))
			return
		}

		userModel := &model.User{
			ID: userIDVal,
		}

		user := usecases.NewUser(userModel)
		modelErr := usecases.Delete(user, db)
		if modelErr != nil {
			context.JSON(http.StatusBadRequest, modelError(modelErr))
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
		return
	}

	return gin.HandlerFunc(fn)
}
