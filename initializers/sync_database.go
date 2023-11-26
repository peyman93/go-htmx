package initializers

import "github.com/peyman93/go-jwt/models"

func SyncDatabase() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("Not able to migrate database")
	}
}
