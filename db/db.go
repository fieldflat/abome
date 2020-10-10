package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Use PostgreSQL in gorm
)

var (
	db  *gorm.DB
	err error
)

// Init is initialize db from main function
func Init() {
	db, err = gorm.Open("postgres", "host=0.0.0.0 port=5432 user=abome dbname=abome password=abome sslmode=disable")
	if err != nil {
		panic(err)
	}
	autoMigration()
	log.Println("[OK] DB Init done!")
}

// GetDB is called in models
func GetDB() *gorm.DB {
	log.Println("[*] call GetDB")
	return db
}

// Close is closing db
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
	log.Println("[OK] DB Close done!")
}

func autoMigration() {
	db.AutoMigrate(&User{})
	log.Println("[OK] DB AutoMigration done!")
}
