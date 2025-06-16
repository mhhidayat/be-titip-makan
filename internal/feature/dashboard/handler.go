package dashboard

import (
	"be-titip-makan/internal/feature/restaurant"
	"be-titip-makan/internal/jsonutil"
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type dashboardHandler struct {
	dashboardService DashboardService
}

func NewDashboard(router fiber.Router, dashboardService DashboardService) {

	ua := dashboardHandler{
		dashboardService: dashboardService,
	}

	router.Get("/categories", ua.ListCategory)
	router.Get("/restaurants", ua.ListRestaurant)
}

func (dh *dashboardHandler) ListCategory(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	categories, err := dh.dashboardService.ListCategory(ctx)
	if err != nil {
		return c.Status(http.StatusInternalServerError).
			JSON(jsonutil.ErrorResponse("Failed to fetch categories: " + err.Error()))
	}

	responseData := map[string]any{
		"categories": categories,
	}

	return c.Status(http.StatusOK).
		JSON(jsonutil.SuccessResponse("Categories fetched successfully", responseData))
}

func (dh *dashboardHandler) ListRestaurant(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	req := restaurant.RestaurantByCategoryRequest{}

	if err := c.BodyParser(&req); err != nil || req.CategoryID == "" {
		return c.Status(http.StatusBadRequest).
			JSON(jsonutil.ErrorResponse("Category ID is required"))
	}

	restaurants, err := dh.dashboardService.ListRestaurantByCategory(ctx, req.CategoryID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).
			JSON(jsonutil.ErrorResponse("Failed to fetch restaurants: " + err.Error()))
	}

	responseData := map[string]any{
		"restaurants": restaurants,
	}

	return c.Status(http.StatusOK).
		JSON(jsonutil.SuccessResponse("Restaurants fetched successfully", responseData))
}
