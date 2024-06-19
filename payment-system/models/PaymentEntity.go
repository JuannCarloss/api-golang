package models

import "github.com/jinzhu/gorm"

type Payment struct {
	gorm.Model
	User_email string  `json:"email"`
	Amount     float32 `json:"amount"`
}
