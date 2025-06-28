package order

import (
	"be-titip-makan/internal/feature/category"
	"be-titip-makan/internal/feature/menu"
	"be-titip-makan/internal/feature/restaurant"
	"context"
)

type orderService struct {
	orderRepository OrderRepository
}

func NewOrderService(orderRepository OrderRepository) OrderService {
	return &orderService{
		orderRepository: orderRepository,
	}
}
func (os *orderService) ListCategory(ctx context.Context) (*[]category.CategoryData, error) {
	categories, err := os.orderRepository.ListCategory(ctx)
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

func (os *orderService) ListRestaurantByCategory(ctx context.Context, categoryId string) (*[]restaurant.RestaurantData, error) {
	restaurants, err := os.orderRepository.ListRestaurantByCategory(ctx, categoryId)
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

func (os *orderService) ListMenuByRestaurant(ctx context.Context, restaurantId string) (*[]menu.MenuData, error) {
	menus, err := os.orderRepository.ListMenuByRestaurant(ctx, restaurantId)
	if err != nil {
		return nil, err
	}

	var restaurantData []menu.MenuData
	for _, v := range *menus {
		restaurantData = append(restaurantData, menu.MenuData{
			ID:           v.ID,
			Name:         v.Name,
			Price:        v.Price,
			RestaurantID: v.RestaurantID,
		})
	}

	return &restaurantData, nil
}

func (os *orderService) Order(ctx context.Context, orderRequest *OrderRequest) (*CreateOrderData, error) {
	createOrdersData, err := os.orderRepository.Order(ctx, orderRequest)
	if err != nil {
		return nil, err
	}
	return createOrdersData, nil
}

func (os *orderService) DeleteDetailOrder(ctx context.Context, deleteOrderDetail *DeleteDetailOrder) error {
	err := os.orderRepository.DeleteDetailOrder(ctx, deleteOrderDetail)
	if err != nil {
		return err
	}
	return nil
}

func (os *orderService) DeleteOrder(ctx context.Context, deleteOrder *DeleteOrder) error {
	err := os.orderRepository.DeleteOrder(ctx, deleteOrder)
	if err != nil {
		return err
	}
	return nil
}
