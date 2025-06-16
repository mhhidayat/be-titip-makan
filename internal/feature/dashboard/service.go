package dashboard

import (
	"be-titip-makan/internal/feature/category"
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
