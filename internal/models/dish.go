package models

type Dish struct {
	ID   uint   `json:"id"`
	Name string `json:"name" binding:"required"`

	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Allergens   string  `json:"allergens"`
	SpiceLevel  string  `json:"spice_level"`
}
