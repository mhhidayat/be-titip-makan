package router

import (
	"be-titip-makan/internal/api"
	"be-titip-makan/internal/config"
	"be-titip-makan/internal/middleware"
	"be-titip-makan/internal/repository"
	"be-titip-makan/internal/service"
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func NewRouterAPI(app *fiber.App, dbConnection *sql.DB, conf *config.Config) {

	// Repository and Service layer
	usersRepository := repository.NewUsers(dbConnection)
	userService := service.NewUsers(usersRepository)

	apiGroup := app.Group("/api")

	api.NewAuth(apiGroup, userService, conf.Auth)

	protected := apiGroup.Group("/", func(c *fiber.Ctx) error {
		return middleware.JWTProtected(c, &conf.Auth)
	})

	api.NewUsers(protected, userService, conf.Auth)

}
