package main

import (
	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	username string `json: "username" gorm: "username"`
	password string `json: "password" gorm: "password"`
	email    string `json: "email" gorm: "email"`
}

var data = []Account{
	{gorm.Model{}, "haarunrasyid", "admin", "rasyid.id3@gmail.com"},
	{gorm.Model{}, "ranggadharma", "admin", "rangga.dharma@hotmail.com"},
}

var DB *gorm.DB

func initDB() {
	var err error
	db, err := gorm.Open(mysql.Open("root:admin@localhost:3306/moviein"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db
	DB.AutoMigrate(&Account{})
}

func main() {
	// Initiate DB
	initDB()

	// Initiate Echo
	e := echo.New()
	e.GET("/")
}

// func
