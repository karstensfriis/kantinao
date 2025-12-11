package models

type WeekMenu struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Week uint   `json:"week"`

	DayMenuItems []DayMenuItem `json:"day_items"`
}

type DayMenuItem struct {
	ID        uint   `json:"id"`
	DayOfWeek string `json:"day_of_week"`
	Likes     int    `json:"likes"`

	Dish Dish `json:"dish,omitempty"`
}
