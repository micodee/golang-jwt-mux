package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(mysql.Open("admin:admin@tcp(localhost:3306)/golang_jwt_mux"))
	if err != nil {
		fmt.Println("gagal koneksi database")
	}

	db.AutoMigrate(
		&User{},
	)

	DB = db

	fmt.Println("connected to database")
}

