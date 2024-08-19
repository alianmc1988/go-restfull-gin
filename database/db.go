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
	config:= base.Config
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", config["DB_HOST"], config["DB_USER"], config["DB_PASSWORD"], config["DB_NAME"], config["DB_PORT"], config["DB_SSL_MODE"])
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