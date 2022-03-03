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

	

}