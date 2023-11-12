package config

import (
	"contactform/api"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func CreateDatabase() *gorm.DB {

	// Create db instance with gorm
	db, err := gorm.Open(sqlite.Open("/formtmp.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	// migrate our model for artist
	db.AutoMigrate(&api.ContactUser{})
	db.AutoMigrate(&api.ContactUserEndereco{})
	db.AutoMigrate(&api.ContactUserVeiculo{})

	return db

}
