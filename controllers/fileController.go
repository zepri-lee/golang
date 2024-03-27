package controllers

import (
	"gin-gonic-gorm/utils"
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

	// file, errFile := fileHeader.Open()
	// if errFile != nil {
	// 	panic(errFile)
	// }

	fileExtention := []string{".jpg", ".png"}
	isFileExtentionValidated, extentionFile := utils.FileValidationByExtention(fileHeader, fileExtention)
	if !isFileExtentionValidated {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "fileExtention is not validated",
		})
		return
	}
	fileType := []string{"image/jpg", "image/png"}
	isFileValidated := utils.FileValidation(fileHeader, fileType)
	if !isFileValidated {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "filetype is not validated",
		})
		return
	}
	filename := utils.RandomFileName(extentionFile, "")
	isSaved := utils.SaveFile(ctx, fileHeader, filename)

	if isSaved {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "file uploaded.",
		})
	}
}
