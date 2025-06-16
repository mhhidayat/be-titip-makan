package dashboard

import (
	"be-titip-makan/internal/feature/category"
	"be-titip-makan/internal/feature/restaurant"
	"context"
)

type DashboardService interface {
	ListCategory(ctx context.Context) (*[]category.CategoryData, error)
	ListRestaurantByCategory(ctx context.Context, categoryId string) (*[]restaurant.RestaurantData, error)
}
