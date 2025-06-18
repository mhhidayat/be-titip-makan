package menu

type MenuData struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Price        string `json:"price"`
	RestaurantID string `json:"restaurant_id"`
}

type MenusByRestaurantRequest struct {
	RestaurantID string `json:"restaurant_id" validate:"required"`
}
