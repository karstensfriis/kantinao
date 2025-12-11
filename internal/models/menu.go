package models

import "github.com/google/uuid"

type WeekMenu struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Week uint      `json:"week"`

	DayMenuItems []DayMenuItem `json:"day_items"`
}

type DayMenuItem struct {
	ID        uuid.UUID `json:"id"`
	DayOfWeek string    `json:"day_of_week"`
	Likes     int       `json:"likes"`

	Dish Dish `json:"dish,omitempty"`
}
