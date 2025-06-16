package main

import (
	"be-titip-makan/configs"
	"be-titip-makan/internal/db"
	"be-titip-makan/internal/feature/auth"
	"be-titip-makan/internal/feature/dashboard"
	"be-titip-makan/internal/feature/user"
	"be-titip-makan/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func main() {

	conf := configs.Get()

	dbConnection := db.GetDatabase(conf.Database)

	app := fiber.New()

	authRepository := auth.NewAuthRepository(dbConnection)
	authService := auth.NewAuthService(authRepository)

	apiGroup := app.Group("/api")

	auth.NewAuth(apiGroup, authService, conf.Auth)

	protected := apiGroup.Group("/", func(c *fiber.Ctx) error {
		return middleware.JWTProtected(c, &conf.Auth)
	})

	user.NewUser(protected, conf.Auth)

	dashboardRepository := dashboard.NewDashboardRepository(dbConnection)
	dashboardService := dashboard.NewDashboardService(dashboardRepository)

	dashboard.NewDashboard(protected, dashboardService)

	app.Listen(":3000")
}
