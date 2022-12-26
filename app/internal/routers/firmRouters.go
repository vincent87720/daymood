package routers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vincent87720/daymood.backend/internal/model"
	"github.com/vincent87720/daymood.backend/internal/settings"
)

func SetupFirmRouters(router *gin.Engine, db *sql.DB, s settings.Settings) (*gin.Engine, error) {

	router.GET("/firm", GetFirmHandler(db))
	router.POST("/firm", PostFirmHandler(db))
	router.PUT("/firm/:id", PutFirmHandler(db))
	router.DELETE("/firm/:id", DeleteFirmHandler(db))
	router.GET("/firm/dumping", DumpFirmHandler(db, s))

	return router, nil
}

func GetFirmHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		firmXi, err := model.GetAllFirm(db)
		if err != nil {
			fmt.Println(err)
			context.JSON(http.StatusBadRequest, gin.H{
				"status": "FAIL",
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"firms":  firmXi,
		})
		return
	}

	return gin.HandlerFunc(fn)
}

func PostFirmHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {

		firmName := context.PostForm("name")
		firmAddress := context.PostForm("address")
		firmRemark := context.PostForm("remark")

		if checkEmpty(firmName) == true {
			err := fmt.Errorf("Invalid input")
			fmt.Println(err.Error())
			context.JSON(http.StatusBadRequest, gin.H{
				"status": "FAIL",
			})
			return
		}

		firm, err := model.NewFirm(firmName, firmAddress, firmRemark)
		if err != nil {
			fmt.Println(err.Error())
			context.JSON(http.StatusBadRequest, gin.H{
				"status": "FAIL",
			})
			return
		}

		err = firm.Create(db)
		if err != nil {
			fmt.Println(err.Error())
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

func PutFirmHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {

		firmID := context.Param("id")
		firmName := context.PostForm("name")
		firmAddress := context.PostForm("address")
		firmRemark := context.PostForm("remark")

		if checkEmpty(firmName) == true {
			err := fmt.Errorf("Invalid input")
			fmt.Println(err.Error())
			context.JSON(http.StatusBadRequest, gin.H{
				"status": "FAIL",
			})
			return
		}

		firmIDVal, err := strconv.ParseInt(firmID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"status":  "FAIL",
				"code":    "PERR2",
				"message": "Invalid input",
			})
			return
		}

		firm, err := model.NewFirm(firmName, firmAddress, firmRemark)
		if err != nil {
			fmt.Println(err.Error())
			context.JSON(http.StatusBadRequest, gin.H{
				"status": "FAIL",
			})
			return
		}
		firm.ID = firmIDVal

		errInfo, err := firm.Update(db)
		if err != nil {
			fmt.Println(err.Error())
			context.JSON(http.StatusBadRequest, gin.H{
				"status":  "FAIL",
				"code":    errInfo.Code,
				"message": errInfo.Message,
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

func DeleteFirmHandler(db *sql.DB) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		// context.Header("Access-Control-Allow-Origin", "*") //allow CORS actions

		firmID := context.Param("id")

		if checkEmpty(firmID) == true {
			err := fmt.Errorf("Invalid input")
			fmt.Println(err.Error())
			context.JSON(http.StatusBadRequest, gin.H{
				"status": "FAIL",
			})
			return
		}

		firmIDVal, err := strconv.ParseInt(firmID, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"status":  "FAIL",
				"code":    "PERR2",
				"message": "Invalid input",
			})
			return
		}

		firm := model.Firm{
			ID: firmIDVal,
		}

		errInfo, err := firm.Delete(db)
		if err != nil {
			fmt.Println(err.Error())
			context.JSON(http.StatusBadRequest, gin.H{
				"status":  "FAIL",
				"code":    errInfo.Code,
				"message": errInfo.Message,
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

func DumpFirmHandler(db *sql.DB, s settings.Settings) gin.HandlerFunc {
	fn := func(context *gin.Context) {
		firmXi, err := model.GetAllFirm(db)
		if err != nil {
			fmt.Println(err)
			context.JSON(http.StatusBadRequest, gin.H{
				"status": "FAIL",
			})
			return
		}

		var prepareCSV [][]string

		prepareCSV = append(prepareCSV, []string{"廠商名稱", "廠商地址", "備註"})

		for _, v := range firmXi {
			tmpXi := []string{}
			tmpXi = append(tmpXi, v.Name, v.Address, v.Remark)
			prepareCSV = append(prepareCSV, tmpXi)
		}

		err = s.WriteCSV(s.GetExeFilePath()+"/firms.csv", prepareCSV)
		if err != nil {
			fmt.Println(err)
			context.JSON(http.StatusBadRequest, gin.H{
				"status": "FAIL",
			})
			return
		}

		context.FileAttachment(s.GetExeFilePath()+"/firms.csv", "firms.csv")
		return
	}

	return gin.HandlerFunc(fn)
}
