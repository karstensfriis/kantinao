package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"kantinao-api/internal/models"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type MenuService interface {
	CreateWeeklyMenu(menu *models.WeekMenu) (*models.WeekMenu, error)
	GetAllMenus() ([]*models.WeekMenu, error)
	GetSingleMenu(menuID string) (*models.WeekMenu, error)
	DeleteMenu(menuID string) error
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
	menu.ID = uuid.New()

	menuKey := fmt.Sprintf("menu:%s", menu.ID)

	jsonData, err := json.Marshal(menu)
	if err != nil {
		return nil, err
	}

	err = s.rdb.Set(s.ctx, menuKey, jsonData, 0).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to store menu: %w", err)
	}

	s.rdb.SAdd(s.ctx, "menus:all_ids", menu.ID)

	return menu, nil
}

func (s *menuService) GetAllMenus() ([]*models.WeekMenu, error) {
	var menus []*models.WeekMenu

	iter := s.rdb.Scan(s.ctx, 0, "menu:*", 50).Iterator()
	for iter.Next(s.ctx) {
		key := iter.Val()

		val, err := s.rdb.Get(s.ctx, key).Result()
		if err != nil {
			return nil, fmt.Errorf("failed to read menu (%s): %w", key, err)
		}

		var menu models.WeekMenu
		if err := json.Unmarshal([]byte(val), &menu); err != nil {
			return nil, fmt.Errorf("failed to decode menu (%s): %w", key, err)
		}

		menus = append(menus, &menu)
	}

	if err := iter.Err(); err != nil {
		return nil, fmt.Errorf("redis scan error: %w", err)
	}

	return menus, nil
}

func (s *menuService) GetSingleMenu(menuID string) (*models.WeekMenu, error) {
	key := fmt.Sprintf("menu:%s", menuID)

	val, err := s.rdb.Get(s.ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, fmt.Errorf("menu with ID %s not found", menuID)
		}
		return nil, fmt.Errorf("redis error: %w", err)
	}

	var menu models.WeekMenu
	if err := json.Unmarshal([]byte(val), &menu); err != nil {
		return nil, fmt.Errorf("invalid menu format: %w", err)
	}

	return &menu, nil
}

func (s *menuService) DeleteMenu(menuID string) error {
	key := fmt.Sprintf("menu:%s", menuID)

	deleted, err := s.rdb.Del(s.ctx, key).Result()
	if err != nil {
		return fmt.Errorf("failed to delete menu %s: %w", menuID, err)
	}

	if deleted == 0 {
		return fmt.Errorf("menu %s not found", menuID)
	}

	s.rdb.SRem(s.ctx, "menus:all_ids", menuID)

	return nil
}
