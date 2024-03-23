package database

import (
	"fmt"
	"gin-gonic-gorm/config"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect() {
	connectionString := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", config.DbUser, config.DbPassword, config.DbServer, config.DbPort, config.DbName)
	Instance, dbError = gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
}
