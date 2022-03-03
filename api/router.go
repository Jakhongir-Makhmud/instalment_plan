package api

import (
	"instalment_plan/api/handlers"
	"instalment_plan/config"
	"instalment_plan/sms"
	"instalment_plan/storage/repo"

	"github.com/gin-gonic/gin"
)

// Option ...
type Option struct {
	Conf config.Config
	Sms  sms.Sms
	DB   repo.DatabaseRepo
}

// New ...
func New(option Option) *gin.Engine {
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handler := handlers.NewHandler(option.Conf,option.DB,option.Sms)

	router.POST("/pay",handler.Pay)
	router.GET("/shop/inform",handler.InformateShop)
	router.POST("/buy/installment",handler.Installment)
	return router
}
