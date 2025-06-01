package main

import (
	"be-titip-makan/internal/config"
	"be-titip-makan/internal/connection"
	"be-titip-makan/internal/router"

	"github.com/gofiber/fiber/v2"
)

func main() {

	conf := config.Get()

	dbConnection := connection.GetDatabase(conf.Database)

	app := fiber.New()

	router.NewRouterAPI(app, dbConnection, conf)

	app.Listen(":3000")
}
