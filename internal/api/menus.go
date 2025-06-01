package api

import "github.com/gofiber/fiber/v2"

func NewMenus(router fiber.Router) {
	router.Post("/cek-token", func(ctx *fiber.Ctx) error {
		return ctx.SendString("SUCCESS")
	})
}
