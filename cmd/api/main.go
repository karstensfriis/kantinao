package main

import (
	"kantinao-api/internal/handlers"
	"kantinao-api/internal/platform"
	"kantinao-api/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	rdb := platform.ConnectRedis()

	router := gin.Default()
	menuHandler := handlers.MenuHandler{Service: services.NewMenuService(rdb)}

	router.POST("/menus", menuHandler.CreateMenu)
	router.GET("/menus/:id", menuHandler.GetMenu)
}
