package middleware

import (
	"be-titip-makan/domain/dto"
	"be-titip-makan/internal/config"
	"be-titip-makan/internal/util"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func JWTProtected(ctx *fiber.Ctx, configAuth *config.Auth) error {
	tokenString := ctx.Get("Authorization")

	if tokenString == "" {
		return ctx.Status(http.StatusBadRequest).
			JSON(dto.ErrorResponse("Missing authorization"))
	}

	err := util.VerifyToken(tokenString, *configAuth)

	if err != nil {
		return ctx.Status(http.StatusUnauthorized).
			JSON(dto.ErrorResponse("Invalid token"))
	}

	return ctx.Next()

}
