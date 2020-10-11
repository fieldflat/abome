package db

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Use PostgreSQL in gorm

	"github.com/fieldflat/abome/entity"
)

var (
	db  *gorm.DB
	err error
)

// Init is initialize db from main function
func Init() {
	// db, err = gorm.Open("postgres", "host=0.0.0.0 port=5432 user=abome dbname=abome password=abome sslmode=disable")
	databaseURL := os.Getenv("DATABASE_URL") // DATABASE_URL="postgres://abome:abome@0.0.0.0:5432/abome?sslmode=disable"
	db, err = gorm.Open("postgres", databaseURL)
	if err != nil {
		panic(err)
	}

	autoMigration()
}

// GetDB is called in models
func GetDB() *gorm.DB {
	return db
}

// Close is closing db
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigration() {
	db.AutoMigrate(&entity.User{})
}
