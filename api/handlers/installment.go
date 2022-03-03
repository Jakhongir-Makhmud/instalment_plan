package handlers

import (
	"instalment_plan/installment"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h handlers) Installment(c *gin.Context) {

	body := installment.Installment{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	m := make(map[string]int)
	m["phone"] = 0
	m["laptop"] = 1
	m["tv"] = 2

	category := m[body.Category]

	totalDebt, err := body.CalculateInstallment(category, body.Duration, body.ProductPrice)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	id := uuid.New()
	additionPrice := totalDebt - body.ProductPrice
	err = h.db.NewProductInstallment(id.String(), body.ProductName, body.CustomerPhone, body.ProductPrice, additionPrice, category, body.Duration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Created",
	})
}
