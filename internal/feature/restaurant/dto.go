package restaurant

type RestaurantData struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	CategoryID string `json:"category_id"`
}

type RestaurantByCategoryRequest struct {
	CategoryID string `json:"category_id" validate:"required"`
}
