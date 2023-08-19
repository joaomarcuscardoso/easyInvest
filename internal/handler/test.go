package handler

import (
	"github.com/gin-gonic/gin"
)

func TestHandler(ctx *gin.Context) {
	logger.Infof("TestHandler called")
}
