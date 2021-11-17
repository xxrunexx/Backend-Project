package config

import (
	// Import GORM
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Global Variable for Database
var DB *gorm.DB

// Initialization Database
func InitDB() {
	var err error
	// For windows
	db, err := gorm.Open(mysql.Open("root:admin@/moviein?parseTime=true"), &gorm.Config{})
	// For Linux
	// db, err := gorm.Open(mysql.Open("root:admin@/moviein?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db
}
