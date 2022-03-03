package api

import (
	"instalment_plan/config"
	"instalment_plan/storage/repo"

	"github.com/gin-gonic/gin"
)

// Option ...
type Option struct {
	Conf config.Config
	db   repo.DatabaseRepo
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
