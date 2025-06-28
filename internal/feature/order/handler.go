package order

import (
	"be-titip-makan/internal/feature/menu"
	"be-titip-makan/internal/feature/restaurant"
	"be-titip-makan/internal/jsonutil"
	"context"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type orderHandler struct {
	orderService OrderService
	validate     *validator.Validate
}

func NewOrder(router fiber.Router, orderService OrderService, validate *validator.Validate) {

	oh := orderHandler{
		orderService: orderService,
		validate:     validate,
	}

	router.Get("/categories", oh.ListCategory)
	router.Get("/restaurants", oh.ListRestaurant)
	router.Get("/menus", oh.Menus)
	router.Post("/order", oh.Order)
	router.Delete("/delete-detail-order", oh.DeleteDetailOrder)
	router.Delete("/delete-order", oh.DeleteOrder)
}

func (dh *orderHandler) ListCategory(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	categories, err := dh.orderService.ListCategory(ctx)
	if err != nil {
		return c.Status(http.StatusInternalServerError).
			JSON(jsonutil.ErrorResponse("Failed to fetch categories: " + err.Error()))
	}

	return c.Status(http.StatusOK).
		JSON(jsonutil.SuccessResponse("Categories fetched successfully", fiber.Map{
			"categories": categories,
		}))
}

func (dh *orderHandler) ListRestaurant(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	req := restaurant.RestaurantByCategoryRequest{}

	c.BodyParser(&req)

	err := dh.validate.Struct(req)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		mappingErros := jsonutil.MappingErrors(validationErrors)
		return c.Status(http.StatusBadRequest).JSON(jsonutil.ValidationErrorResponse(fiber.Map{
			"errors": mappingErros,
		}))
	}

	restaurants, err := dh.orderService.ListRestaurantByCategory(ctx, req.CategoryID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).
			JSON(jsonutil.ErrorResponse("Failed to fetch restaurants: " + err.Error()))
	}

	return c.Status(http.StatusOK).
		JSON(jsonutil.SuccessResponse("Restaurants fetched successfully", fiber.Map{
			"restaurants": restaurants,
		}))
}

func (dh *orderHandler) Menus(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	req := menu.MenusByRestaurantRequest{}

	c.BodyParser(&req)

	err := dh.validate.Struct(req)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		mappingErros := jsonutil.MappingErrors(validationErrors)
		return c.Status(http.StatusBadRequest).JSON(jsonutil.ValidationErrorResponse(fiber.Map{
			"errors": mappingErros,
		}))
	}

	menus, err := dh.orderService.ListMenuByRestaurant(ctx, req.RestaurantID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).
			JSON(jsonutil.ErrorResponse("Failed to fetch menus: " + err.Error()))
	}

	return c.Status(http.StatusOK).
		JSON(jsonutil.SuccessResponse("Menus fetched successfully", fiber.Map{
			"menus": menus,
		}))
}

func (dh *orderHandler) Order(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	req := OrderRequest{}

	c.BodyParser(&req)

	err := dh.validate.Struct(req)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		mappingErros := jsonutil.MappingErrors(validationErrors)
		return c.Status(http.StatusBadRequest).JSON(jsonutil.ValidationErrorResponse(fiber.Map{
			"errors": mappingErros,
		}))
	}

	createOrders, err := dh.orderService.Order(ctx, &req)

	if err != nil {
		return c.Status(http.StatusInternalServerError).
			JSON(jsonutil.ErrorResponse("Failed to store order: " + err.Error()))
	}

	return c.Status(http.StatusOK).
		JSON(jsonutil.SuccessResponse("Order stored successfully", fiber.Map{
			"orders": createOrders,
		}))
}

func (dh *orderHandler) DeleteDetailOrder(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	req := DeleteDetailOrder{}

	c.BodyParser(&req)

	err := dh.validate.Struct(req)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		mappingErros := jsonutil.MappingErrors(validationErrors)
		return c.Status(http.StatusBadRequest).JSON(jsonutil.ValidationErrorResponse(fiber.Map{
			"errors": mappingErros,
		}))
	}

	err = dh.orderService.DeleteDetailOrder(ctx, &req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).
			JSON(jsonutil.ErrorResponse("Failed to delete detail order"))
	}

	return c.Status(http.StatusOK).
		JSON(jsonutil.SuccessResponse("Delete detail order successfully", ""))
}

func (dh *orderHandler) DeleteOrder(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	req := DeleteOrder{}

	c.BodyParser(&req)

	err := dh.validate.Struct(req)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		mappingErros := jsonutil.MappingErrors(validationErrors)
		return c.Status(http.StatusBadRequest).JSON(jsonutil.ValidationErrorResponse(fiber.Map{
			"errors": mappingErros,
		}))
	}

	err = dh.orderService.DeleteOrder(ctx, &req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).
			JSON(jsonutil.ErrorResponse("Failed to delete order"))
	}

	return c.Status(http.StatusOK).
		JSON(jsonutil.SuccessResponse("Delete order successfully", ""))
}
