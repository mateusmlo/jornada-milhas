package config

import (
	"fmt"

	"github.com/mateusmlo/jornada-milhas/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDBConnection connects to database driver
func NewDBConnection(env *Env) *gorm.DB {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo", env.DBHost, env.DBUsername, env.DBPassword, env.DBName, env.DBPort)

	DB, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{
			FullSaveAssociations: true,
		},
	)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("[DB]            ✅ Successfully connected to the database")
	}

	DB.AutoMigrate(models.Review{}, models.User{})

	return DB
}
