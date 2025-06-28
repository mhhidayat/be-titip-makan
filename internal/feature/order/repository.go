package order

import (
	"be-titip-makan/internal/feature/category"
	"be-titip-makan/internal/feature/menu"
	"be-titip-makan/internal/feature/restaurant"
	"be-titip-makan/internal/strutil"
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/doug-martin/goqu/v9"
)

type orderRepository struct {
	db *goqu.Database
}

func NewOrderRepository(con *sql.DB) OrderRepository {
	return &orderRepository{
		db: goqu.New("default", con),
	}
}

func (dr *orderRepository) ListCategory(ctx context.Context) (*[]category.Model, error) {
	var categories []category.Model

	dataset := dr.db.From("mst_categories")

	if err := dataset.ScanStructsContext(ctx, &categories); err != nil {
		return nil, err
	}

	return &categories, nil
}

func (dr *orderRepository) ListRestaurantByCategory(ctx context.Context, categoryId string) (*[]restaurant.Model, error) {
	var restaurant []restaurant.Model

	dataset := dr.db.From("mst_restaurants").Where(goqu.Ex{
		"category_id": categoryId,
	})

	if err := dataset.ScanStructsContext(ctx, &restaurant); err != nil {
		return nil, err
	}

	return &restaurant, nil
}

func (dr *orderRepository) ListMenuByRestaurant(ctx context.Context, restaurantId string) (*[]menu.Model, error) {
	var menus []menu.Model

	dataset := dr.db.From("mst_menus").Where(goqu.Ex{
		"restaurant_id": restaurantId,
	})

	if err := dataset.ScanStructsContext(ctx, &menus); err != nil {
		return nil, err
	}

	return &menus, nil
}

func (dr *orderRepository) Order(ctx context.Context, orderRequest *OrderRequest) (*CreateOrderData, error) {

	orderNumber := fmt.Sprintf("ORD%s%s", time.Now().Format("060102"), strutil.GenerateRandomString(3))

	createOrder := CreateOrder{
		OrderNumber:   orderNumber,
		UserID:        orderRequest.UserID,
		PaymentStatus: orderRequest.PaymentStatus,
		TotalAmount:   orderRequest.TotalAmount,
		PaymentMethod: orderRequest.PaymentMethod,
	}

	if _, err := dr.db.Insert("tr_orders").
		Rows(&createOrder).
		Executor().ExecContext(ctx); err != nil {
		return nil, err
	}

	createOrderDetail := make([]CreateOrderDetail, 0, len(orderRequest.Detail))

	for _, val := range orderRequest.Detail {
		createOrderDetail = append(createOrderDetail, CreateOrderDetail{
			OrderNumber: orderNumber,
			MenuID:      val.MenuID,
			Qty:         val.Qty,
			Price:       val.Price,
			Description: val.Description,
		})
	}

	if _, err := dr.db.Insert("tr_order_details").
		Rows(createOrderDetail).
		Executor().ExecContext(ctx); err != nil {
		return nil, err
	}

	return &CreateOrderData{
		CreateOrder: createOrder,
		Detail:      createOrderDetail,
	}, nil
}

func (dr *orderRepository) DeleteDetailOrder(ctx context.Context, deleteOrderDetail *DeleteDetailOrder) error {

	expresion := goqu.Ex{
		"order_number": deleteOrderDetail.OrderNumber,
		"menu_id":      deleteOrderDetail.MenuID,
	}

	if _, err := dr.db.Delete("tr_order_details").
		Where(expresion).
		Executor().
		ExecContext(ctx); err != nil {
		return err
	}

	return nil
}

func (dr *orderRepository) DeleteOrder(ctx context.Context, deleteOrder *DeleteOrder) error {

	expresion := goqu.Ex{
		"order_number": deleteOrder.OrderNumber,
	}

	if _, err := dr.db.Delete("tr_orders").
		Where(expresion).
		Executor().
		ExecContext(ctx); err != nil {
		return err
	}

	return nil
}
