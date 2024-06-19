package main

import (
	"example/payments/configs"
	"example/payments/models"
	"example/payments/routers"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var err error

func main() {
	configs.DB, err = gorm.Open("postgres", "postgresql://postgres:123@localhost/payment-db?sslmode=disable")

	if err != nil {
		fmt.Println("status: ", err)
	}

	defer configs.DB.Close()
	configs.DB.AutoMigrate(&models.Payment{})

	r := routers.SetupRouters()

	r.Run()
}
