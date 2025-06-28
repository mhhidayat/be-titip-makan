package order

import (
	"be-titip-makan/internal/feature/category"
	"be-titip-makan/internal/feature/menu"
	"be-titip-makan/internal/feature/restaurant"
	"context"
)

type OrderRepository interface {
	ListCategory(ctx context.Context) (*[]category.Model, error)
	ListRestaurantByCategory(ctx context.Context, categoryId string) (*[]restaurant.Model, error)
	ListMenuByRestaurant(ctx context.Context, restaurantId string) (*[]menu.Model, error)
	Order(ctx context.Context, orderRequest *OrderRequest) (*CreateOrderData, error)
	DeleteDetailOrder(ctx context.Context, deleteOrderDetail *DeleteDetailOrder) error
	DeleteOrder(ctx context.Context, deleteDetail *DeleteOrder) error
}
