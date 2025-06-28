package order

import (
	"be-titip-makan/internal/feature/category"
	"be-titip-makan/internal/feature/menu"
	"be-titip-makan/internal/feature/restaurant"
	"context"
)

type OrderService interface {
	ListCategory(ctx context.Context) (*[]category.CategoryData, error)
	ListRestaurantByCategory(ctx context.Context, categoryId string) (*[]restaurant.RestaurantData, error)
	ListMenuByRestaurant(ctx context.Context, restaurantId string) (*[]menu.MenuData, error)
	Order(ctx context.Context, orderRequest *OrderRequest) (*CreateOrderData, error)
	DeleteDetailOrder(ctx context.Context, deleteOrderDetail *DeleteDetailOrder) error
	DeleteOrder(ctx context.Context, deleteOrder *DeleteOrder) error
}
