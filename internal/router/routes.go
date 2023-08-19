package router

import (
	"github.com/gin-gonic/gin"
	"github.com/joaomarcuscardoso/easyInvest/internal/handler"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/test", handler.TestHandler)
	}
}
