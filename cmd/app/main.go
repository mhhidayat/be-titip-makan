package main

import (
	"be-titip-makan/configs"
	"be-titip-makan/internal/db"
	"be-titip-makan/internal/feature/auth"
	"be-titip-makan/internal/feature/order"
	"be-titip-makan/internal/feature/user"
	"be-titip-makan/internal/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func main() {

	conf := configs.Get()
	var validate = validator.New()

	dbConnection := db.GetDatabase(conf.Database)

	app := fiber.New()

	authRepository := auth.NewAuthRepository(dbConnection)
	authService := auth.NewAuthService(authRepository)

	apiGroup := app.Group("/api")

	auth.NewAuth(apiGroup, authService, conf.Auth, validate)

	protected := apiGroup.Group("/", func(c *fiber.Ctx) error {
		return middleware.JWTProtected(c, &conf.Auth)
	})

	user.NewUser(protected, conf.Auth)

	orderRepository := order.NewOrderRepository(dbConnection)
	orderService := order.NewOrderService(orderRepository)

	order.NewOrder(protected, orderService, validate)

	app.Listen(":3000")
}
