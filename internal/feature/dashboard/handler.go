package dashboard

import (
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
}

func (dh *dashboardHandler) ListCategory(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	categories, err := dh.dashboardService.ListCategory(ctx)
	if err != nil {
		return c.Status(http.StatusBadRequest).
			JSON(jsonutil.ErrorResponse("Failed to fetch categories: " + err.Error()))
	}

	return c.Status(http.StatusOK).
		JSON(jsonutil.SuccessResponse("Categories fetched successfully", categories))
}
