package handlers

import (
	"kantinao-api/internal/models"
	"kantinao-api/internal/services"
	"net/http"

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

func (h *MenuHandler) GetAllMenus(c *gin.Context) {
	menus, err := h.Service.GetAllMenus()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, menus)
}

func (h *MenuHandler) GetMenu(c *gin.Context) {
	id := c.Param("id")

	menu, err := h.Service.GetSingleMenu(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, menu)
}

func (h *MenuHandler) DeleteMenu(c *gin.Context) {
	id := c.Param("id")

	if err := h.Service.DeleteMenu(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "menu deleted successfully"})
}
