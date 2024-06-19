package services

import (
	"example/payments/configs"
	"example/payments/models"
)

func GetAllPayments(p *[]models.Payment) (err error) {
	if err = configs.DB.Find(p).Error; err != nil {
		return err
	}

	return nil
}

func PostPayment(p *models.Payment) (err error) {
	if err = configs.DB.Create(p).Error; err != nil {
		return err
	}

	return nil
}
