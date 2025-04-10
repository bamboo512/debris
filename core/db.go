package core

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func GetDB(path string) (db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
