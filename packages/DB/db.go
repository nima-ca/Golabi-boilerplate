package db

import "gorm.io/gorm"

// Return DB
func GetDB() *gorm.DB {
	return DB
}
