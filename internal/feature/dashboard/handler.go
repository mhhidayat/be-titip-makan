package dashboard

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

type dashboardHandler struct {
	dashboardService DashboardService
	validate         *validator.Validate
}

func NewDashboard(router fiber.Router, dashboardService DashboardService, validate *validator.Validate) {

	ua := dashboardHandler{
		dashboardService: dashboardService,
		validate:         validate,
	}

	router.Get("/categories", ua.ListCategory)
	router.Get("/restaurants", ua.ListRestaurant)
	router.Get("/menus", ua.Menus)
}

func (dh *dashboardHandler) ListCategory(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	categories, err := dh.dashboardService.ListCategory(ctx)
	if err != nil {
		return c.Status(http.StatusInternalServerError).
			JSON(jsonutil.ErrorResponse("Failed to fetch categories: " + err.Error()))
	}

	return c.Status(http.StatusOK).
		JSON(jsonutil.SuccessResponse("Categories fetched successfully", fiber.Map{
			"categories": categories,
		}))
}

func (dh *dashboardHandler) ListRestaurant(c *fiber.Ctx) error {
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

	restaurants, err := dh.dashboardService.ListRestaurantByCategory(ctx, req.CategoryID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).
			JSON(jsonutil.ErrorResponse("Failed to fetch restaurants: " + err.Error()))
	}

	return c.Status(http.StatusOK).
		JSON(jsonutil.SuccessResponse("Restaurants fetched successfully", fiber.Map{
			"restaurants": restaurants,
		}))
}

func (dh *dashboardHandler) Menus(c *fiber.Ctx) error {
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

	menus, err := dh.dashboardService.ListMenuByRestaurant(ctx, req.RestaurantID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).
			JSON(jsonutil.ErrorResponse("Failed to fetch menus: " + err.Error()))
	}

	return c.Status(http.StatusOK).
		JSON(jsonutil.SuccessResponse("Menus fetched successfully", fiber.Map{
			"menus": menus,
		}))
}
