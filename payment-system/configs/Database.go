package configs

import (
	"example/payments/models"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	DB, err = gorm.Open("postgres", "postgresql://postgres:123@localhost/payment-db?sslmode=disable")

	if err != nil {
		fmt.Println("status: ", err)
	}

	defer DB.Close()
	DB.AutoMigrate(&models.Payment{})
}
