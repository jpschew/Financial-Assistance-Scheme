package db

import (
	"FinancialAssistanceScheme/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var dbConn *gorm.DB

// InitDatabase initializes the database connection using GORM
func InitDatabase(appConfig config.AppConfig) {
	dsn := fmt.Sprintf("%s:%s@(%s:%v)/%s?charset=utf8mb4&parseTime=true",
		appConfig.Database.User, appConfig.Database.Password, appConfig.Database.Host, appConfig.Database.Port, appConfig.Database.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
		//return nil, err
		return
	}
	log.Println("Connected to the database")
	dbConn = db
}

func GetDBConn() *gorm.DB {
	return dbConn
}
