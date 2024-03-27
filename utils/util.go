package utils

import (
	"fmt"
	"gin-gonic-gorm/constants"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

var charset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

func FileValidation(fileHeader *multipart.FileHeader, fileType []string) bool {
	contentType := fileHeader.Header.Get("Content-Type")
	result := false

	for _, typeFile := range fileType {
		if contentType == typeFile {
			result = true
			break
		}
	}

	return result
}

func FileValidationByExtention(fileHeader *multipart.FileHeader, fileExtention []string) (bool, string) {
	extention := filepath.Ext(fileHeader.Filename)
	result := false

	for _, extentionFile := range fileExtention {
		if extention == extentionFile {
			result = true
			break
		}
	}

	return result, extention
}

func RandomFileName(extentionFile string, prefix ...string) string {
	if prefix[0] == "" {
		prefix[0] = "file"
	}

	currentTime := time.Now().UTC().Format("20061206") // 현시간을 왜 못 가져오지?
	filename := fmt.Sprintf("%s-%s-%s%s", prefix[0], currentTime, RandomString(5), extentionFile)

	return filename
}

func SaveFile(ctx *gin.Context, fileHeader *multipart.FileHeader, filename string) bool {
	errUpload := ctx.SaveUploadedFile(fileHeader, fmt.Sprintf("./public/files/%s", filename))
	if errUpload != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": errUpload.Error(),
		})
		log.Println("Can't save file")
		return false
	} else {
		return true
	}
}

func RemoveFile(fileName string) error {
	err := os.Remove(constants.FILE_DIR + fileName)
	if err != nil {
		log.Println("Failed to remove file")
		return err
	}

	return nil
}
