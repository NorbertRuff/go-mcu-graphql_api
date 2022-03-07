package database

import (
	"fmt"
	"github.com/norbertruff/go-graphql/graphql/models"
	"github.com/norbertruff/go-graphql/internal/pkg/db/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB             *gorm.DB
	PostgresConfig configs.PostgresConfig
)

func InitDB() {
	// open database
	PostgresConfig = configs.GetPostgresConfig()
	psqlConnString := PostgresConfig.GetPostgresConnectionInfo()

	db, err := gorm.Open(postgres.Open(psqlConnString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connected!")
	DB = db
}

func MigrateData() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("failed to migrate database")
	}
	fmt.Println("Database Migrated")
}
