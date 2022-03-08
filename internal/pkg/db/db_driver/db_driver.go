package database

import (
	"github.com/norbertruff/go-graphql/graphql/model"
	"github.com/norbertruff/go-graphql/internal/pkg/db/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	DB             *gorm.DB
	PostgresConfig configs.PostgresConfig
)

// InitDB initializes connection to db based on postgres config
func InitDB() {
	// open database
	PostgresConfig = configs.GetPostgresConfig()
	psqlConnString := PostgresConfig.GetPostgresConnectionInfo()

	db, err := gorm.Open(postgres.Open(psqlConnString), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	log.Println("Connected!")
	DB = db
}

// MigrateData migrates model in db
func MigrateData() {
	err := DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Panic("failed to migrate database")
	}
	log.Println("Database Migrated")
}
