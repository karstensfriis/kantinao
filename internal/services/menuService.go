// internal/services/menuService.go

package services

import (
	"context"
	"fmt"
	"strconv"

	"kantinao-api/internal/models"

	"github.com/redis/go-redis/v9"
)

type MenuService interface {
	CreateWeeklyMenu(menu *models.WeekMenu) (*models.WeekMenu, error)
	GetWeeklyMenu(menuID uint) (*models.WeekMenu, error)
}

type menuService struct {
	rdb *redis.Client
	ctx context.Context
}

func NewMenuService(rdb *redis.Client) MenuService {
	return &menuService{
		rdb: rdb,
		ctx: context.Background(),
	}
}

func (s *menuService) CreateWeeklyMenu(menu *models.WeekMenu) (*models.WeekMenu, error) {
	menuID, err := s.rdb.Incr(s.ctx, "menu:id_counter").Result()
	if err != nil {
		return nil, fmt.Errorf("failed to generate menu ID: %w", err)
	}
	menu.ID = uint(menuID)

	menuKey := fmt.Sprintf("menu:%d", menu.ID)

	err = s.rdb.HSet(s.ctx, menuKey, map[string]interface{}{
		"ID":   menu.ID,
		"Name": menu.Name,
	}).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to save menu to Redis: %w", err)
	}

	s.rdb.SAdd(s.ctx, "menus:all_ids", menuID)

	return menu, nil
}

func (s *menuService) GetWeeklyMenu(menuID uint) (*models.WeekMenu, error) {
	menuKey := fmt.Sprintf("menu:%d", menuID)
	data, err := s.rdb.HGetAll(s.ctx, menuKey).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get menu from Redis: %w", err)
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("menu with ID %d not found", menuID)
	}

	id, _ := strconv.ParseUint(data["ID"], 10, 64)
	menu := &models.WeekMenu{
		ID:   uint(id),
		Name: data["Name"],
	}

	return menu, nil
}
