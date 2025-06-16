package dashboard

import (
	"be-titip-makan/internal/feature/category"
	"be-titip-makan/internal/feature/menu"
	"be-titip-makan/internal/feature/restaurant"
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

func (dr *dashboardRepository) ListRestaurantByCategory(ctx context.Context, categoryId string) (*[]restaurant.Model, error) {
	var restaurant []restaurant.Model

	dataset := dr.db.From("mst_restaurants").Where(goqu.Ex{
		"category_id": categoryId,
	})

	if err := dataset.ScanStructsContext(ctx, &restaurant); err != nil {
		return nil, err
	}

	return &restaurant, nil
}

func (dr *dashboardRepository) ListMenuByRestaurant(ctx context.Context, restaurantId string) (*[]menu.Model, error) {
	var menus []menu.Model

	dataset := dr.db.From("mst_menus").Where(goqu.Ex{
		"restaurant_id": restaurantId,
	})

	if err := dataset.ScanStructsContext(ctx, &menus); err != nil {
		return nil, err
	}

	return &menus, nil
}
