package handler

import "github.com/joaomarcuscardoso/easyInvest/internal/config"

var (
	logger *config.Logger
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
}
