package controllers

import (
	"fmt"
	"gin-gonic-gorm/utils"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func HandleUploadFile(ctx *gin.Context) {

	fileHeader, _ := ctx.FormFile("file")
	if fileHeader == nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "file is required.",
		})
		return
	}

	// file, errFile := fileHeader.Open()
	// if errFile != nil {
	// 	panic(errFile)
	// }

	extentionFile := filepath.Ext(fileHeader.Filename)
	fmt.Println("extentionFile : " + extentionFile)
	currentTime := time.Now().UTC().Format("20060101")
	filename := fmt.Sprintf("%s-%s%s", currentTime, utils.RandomString(5), extentionFile)

	errUpload := ctx.SaveUploadedFile(fileHeader, fmt.Sprintf("./public/files/%s", filename))
	if errUpload != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": errUpload.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "file uploaded.",
	})
}
