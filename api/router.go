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
	Sms  sms.Sms
	DB   repo.DatabaseRepo
}

// New ...
func New(option Option) *gin.Engine {
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	

	return router
}
