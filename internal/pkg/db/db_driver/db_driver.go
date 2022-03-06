package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "marveluniverse"
)

var (
	DB *gorm.DB
)

func InitDB() {

	// open database
	psqlConnString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(psqlConnString), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connected!")

	//err = db.AutoMigrate(&model.NewUser{})
	//err = db.AutoMigrate(&model.User{})
	//if err != nil {
	//	panic("failed to migrate database")
	//}
	//fmt.Println("Database Migrated")
	DB = db
}
