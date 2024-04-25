package database

import (
	"fmt"

	c "github.com/LaMoldy/TheBowlShop/config"
	"github.com/LaMoldy/TheBowlShop/internal/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateDatabaseConnection(config *c.Config) (*gorm.DB, error) {
	connectionString := fmt.Sprintf(
		"user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		config.Database.User,
		config.Database.Password,
		config.Database.Name,
		config.Database.Port,
		config.Database.SSLMode,
		config.Database.Timezone,
	)

	// Open the connection
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: connectionString,
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	if err != nil {
		return nil, err
	}

	// Migrate database models
	db.AutoMigrate(
		&models.User{},
		&models.Cart{},
		&models.Comment{},
		&models.Product{},
		&models.Rating{},
	)

	return db, nil
}
