package database

import (
	"dumbmerch/models"
	"dumbmerch/pkg/mysql"
	"fmt"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Profile{},
		&models.Category{},
		&models.Transaction{},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Migration success cuy")
}
