package controllers

import (
	"fmt"
	"net/http"

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

	//	file, errFile := fileHeader.Open()
	//	if errFile != nil {
	//		panic(errFile)
	//	}

	//	extentionFile := filepath.Ext(fileHeader.Filename)
	//	filename :=
	errUpload := ctx.SaveUploadedFile(fileHeader, fmt.Sprintf("./public/files/%s", fileHeader.Filename))
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
