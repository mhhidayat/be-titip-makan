package dashboard

import (
	"be-titip-makan/internal/feature/category"
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

type dashboardRepository struct {
	db *goqu.Database
}

func NewDashboardRepository(con *sql.DB) DashboardRepository {
	return &dashboardRepository{
		db: goqu.New("default", con),
	}
}
func (dr *dashboardRepository) ListCategory(ctx context.Context) (*[]category.Model, error) {
	var categories []category.Model

	dataset := dr.db.From("mst_categories").Select("*")

	if err := dataset.ScanStructsContext(ctx, &categories); err != nil {
		return nil, err
	}

	return &categories, nil
}
