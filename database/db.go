package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB *gorm.DB
var dsn = "host=localhost user=postgres password=postgres dbname=gopostgres port=5432"
var err error

func DBConnect() *gorm.DB {
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	return DB
}