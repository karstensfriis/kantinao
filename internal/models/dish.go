package models

import "github.com/google/uuid"

type Dish struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name" binding:"required"`

	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Allergens   string  `json:"allergens"`
	SpiceLevel  string  `json:"spice_level"`
}
