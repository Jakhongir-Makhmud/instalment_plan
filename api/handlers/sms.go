package handlers

import (
	"fmt"
	"instalment_plan/storage/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h handlers) Pay(c *gin.Context) {

	body := models.Pay{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	remaining,err := h.db.Pay(body.ProductName,body.CustomerPhone,body.InstallmentDuration,body.PayingAmount)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	msg := fmt.Sprintf(
		`Paid amount: %v,
		Remaining debt: %v
		`, body.PayingAmount, remaining,
	)

	if err = h.sms.SendSms(body.CustomerPhone,msg); err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"cannot send sms" + err.Error(),
		})
	}

	c.JSON(http.StatusOK,gin.H{
		"message":"ok",
	})

}