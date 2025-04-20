package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("myapp.db"), &gorm.Config{})
}
