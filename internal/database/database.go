package database

import "gorm.io/gorm"

type Database interface {
	Getdb() *gorm.DB
}
