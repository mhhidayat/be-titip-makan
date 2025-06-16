package dashboard

import (
	"be-titip-makan/internal/feature/category"
	"be-titip-makan/internal/feature/restaurant"
	"context"
)

type DashboardRepository interface {
	ListCategory(ctx context.Context) (*[]category.Model, error)
	ListRestaurantByCategory(ctx context.Context, categoryId string) (*[]restaurant.Model, error)
}
