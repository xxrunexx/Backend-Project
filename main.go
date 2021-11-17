package main

import (
	// standard Library

	// Import Framework Echo, GORM, & Middleware

	"gorm.io/gorm"

	// Import Files
	"movie-api/config"
	"movie-api/migrate"
	"movie-api/route"
)

// Global Variable For Table Account
type Account struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func main() {
	// Initiate DB from config/database.go
	config.InitDB()

	// Initiate Router from route/route.go
	e := route.Router()

	// Initiate Migrate from migrate/migrate.go
	migrate.AutoMigrate()

	// Starting the server
	e.Start(":8585")
}
