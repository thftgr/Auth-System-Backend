package database

import (
	"Auth-System-Backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:@song4622@tcp(127.0.0.1:4622)/backend?charset=utf8"), &gorm.Config{})

	if err != nil {
		panic("could not connect to yhe database")
	}

	DB = connection
	connection.AutoMigrate(&models.User{})
}
