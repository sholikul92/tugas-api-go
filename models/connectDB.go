package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/mahasiswa_unipi"))
	// Jika koneksi ke database gagal maka akan mereturn error
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Users{})

	DB = db
}
