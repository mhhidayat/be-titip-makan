package dashboard

import (
	"be-titip-makan/internal/feature/category"
	"be-titip-makan/internal/feature/restaurant"
	"context"
)

type dashboardService struct {
	dashboardRepository DashboardRepository
}

func NewDashboardService(dashboardRepository DashboardRepository) DashboardService {
	return &dashboardService{
		dashboardRepository: dashboardRepository,
	}
}
func (ds *dashboardService) ListCategory(ctx context.Context) (*[]category.CategoryData, error) {
	categories, err := ds.dashboardRepository.ListCategory(ctx)
	if err != nil {
		return nil, err
	}

	var categoriesData []category.CategoryData
	for _, v := range *categories {
		categoriesData = append(categoriesData, category.CategoryData{
			ID:   v.ID,
			Name: v.Name,
		})
	}

	return &categoriesData, nil
}

func (ds *dashboardService) ListRestaurantByCategory(ctx context.Context, categoryId string) (*[]restaurant.RestaurantData, error) {
	restaurants, err := ds.dashboardRepository.ListRestaurantByCategory(ctx, categoryId)
	if err != nil {
		return nil, err
	}

	var restaurantData []restaurant.RestaurantData
	for _, v := range *restaurants {
		restaurantData = append(restaurantData, restaurant.RestaurantData{
			ID:         v.ID,
			Name:       v.Name,
			CategoryID: v.CategoryID,
		})
	}

	return &restaurantData, nil
}
