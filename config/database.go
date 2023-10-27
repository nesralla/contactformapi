package config

import (
	"contactform/api"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func CreateDatabase() *gorm.DB {

	// Create db instance with gorm
	db, err := gorm.Open("sqlite3", "/ksronrvtmp.db")
	if err != nil {
		panic("Failed to connect to database!")
	}

	// migrate our model for artist
	db.AutoMigrate(&api.ContactUser{})

	return db

}
