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

	fileExtention := []string{".jpg", ".png", ".PNG"}
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

func HadndleRemoveFile(ctx *gin.Context) {
	fileName := ctx.Param("fileName")
	if fileName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message1": "fileName is required.",
		})
		return
	}

	err := utils.RemoveFile(fileName)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message2": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message3": "file deleted.",
	})
}
