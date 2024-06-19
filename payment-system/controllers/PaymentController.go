package controllers

import (
	"example/payments/handlers"
	"example/payments/models"
	"example/payments/services"

	"github.com/gin-gonic/gin"
)

func FindAll(c *gin.Context) {
	var list []models.Payment
	err := services.GetAllPayments(&list)

	if err != nil {
		handlers.RespondJSON(c, 204, list)
	} else {
		handlers.RespondJSON(c, 200, list)
	}
}

func PostPayment(c *gin.Context) {
	var pay models.Payment
	c.BindJSON(&pay)
	err := services.PostPayment(&pay)

	if err != nil {
		handlers.RespondJSON(c, 422, pay)
	} else {
		handlers.RespondJSON(c, 201, pay)
	}
}
