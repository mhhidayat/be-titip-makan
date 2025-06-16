package dashboard

import (
	"be-titip-makan/internal/feature/category"
	"context"
)

type DashboardService interface {
	ListCategory(ctx context.Context) (*[]category.CategoryData, error)
}
