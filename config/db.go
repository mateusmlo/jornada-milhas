package config

import (
	"fmt"

	"github.com/mateusmlo/jornada-milhas/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDBConnection connects to database driver
func NewDBConnection(env Env, logger Logger) *gorm.DB {
	var err error

	host := env.DBHost
	username := env.DBUsername
	password := env.DBPassword
	dbName := env.DBName
	port := env.DBPort

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo", host, username, password, dbName, port)

	DB, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{
			FullSaveAssociations: true,
			Logger:               logger.GetGormLogger(),
		},
	)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}

	DB.AutoMigrate(models.Review{}, models.User{})

	return DB
}
