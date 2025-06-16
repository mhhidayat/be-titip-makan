package dashboard

import (
	"be-titip-makan/internal/feature/category"
	"context"
)

type DashboardRepository interface {
	ListCategory(ctx context.Context) (*[]category.Model, error)
}
