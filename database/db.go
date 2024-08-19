package database

import (
	"fmt"

	"github.com/alianmc1988/go-restfull-gin/base"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB *gorm.DB
var err error

func buildDsn() string {
	config:= base.GetConfig()
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", config.DbHost, config.DbUser, config.DbPassword, config.DbName, config.DbPort, config.DbSSLMode)
}

func DBConnect() *gorm.DB {
	dsn:= buildDsn()
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	return DB
}