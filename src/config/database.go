package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectDatabase() {
	dsn := "root:12345@tcp(localhost:3306)/mockbank?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("[Connection Database Error]", err)
		return
	}
	fmt.Println("Connect to mysql database")
}
