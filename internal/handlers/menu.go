// internal/handlers/menuHandler.go

package handlers

import (
	"net/http"
	"strconv"

	"kantinao-api/internal/models"
	"kantinao-api/internal/services"

	"github.com/gin-gonic/gin"
)

type MenuHandler struct {
	Service services.MenuService
}

func (h *MenuHandler) CreateMenu(c *gin.Context) {
	var newMenu models.WeekMenu

	if err := c.ShouldBindJSON(&newMenu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdMenu, err := h.Service.CreateWeeklyMenu(&newMenu)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdMenu)
}

func (h *MenuHandler) GetMenu(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid menu ID"})
		return
	}

	menu, err := h.Service.GetWeeklyMenu(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, menu)
}
