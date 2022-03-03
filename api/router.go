package api

import (
	"instalment_plan/config"
	"instalment_plan/sms"
	"instalment_plan/storage/repo"

	"github.com/gin-gonic/gin"
)

// Option ...
type Option struct {
	Conf config.Config
	DB  repo.DatabaseRepo
	Sms sms.Sms
}

// New ...
func New(option Option) *gin.Engine {
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	
	// handler := handlers.NewHandler(option.db)

	// router.GET("/",)

	return router
}
