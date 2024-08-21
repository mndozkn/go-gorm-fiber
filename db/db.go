package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	return gorm.Open(postgres.Open("postgres://root:root@localhost:5432/go-tutorial"), &gorm.Config{})
}
