package server

import (
	"mini-wallet-julo/controllers"
	"mini-wallet-julo/middlewares"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)
	router.Use(middlewares.AuthMiddleware())

	v1 := router.Group("/api")
	{
		v1Group := v1.Group("/v1")
		{
			v1Group.GET("/init", controllers.InitWallet)
		}
	}
	return router

}
