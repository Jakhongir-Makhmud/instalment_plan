package handlers

import (
	shopinformative "instalment_plan/shop_informative"
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h handlers) InformateShop(c *gin.Context) {

	body := shopinformative.ShopInformative{}

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	totalPaid,err := body.GetInfo(h.db)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"total paid":totalPaid,
	})
}