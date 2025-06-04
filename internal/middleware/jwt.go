package middleware

import (
	"be-titip-makan/configs"
	"be-titip-makan/internal/jsonutil"
	"be-titip-makan/internal/jwtutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func JWTProtected(ctx *fiber.Ctx, configAuth *configs.Auth) error {
	tokenString := ctx.Get("Authorization")

	if tokenString == "" {
		return ctx.Status(http.StatusBadRequest).
			JSON(jsonutil.ErrorResponse("Missing authorization"))
	}

	err := jwtutil.VerifyToken(tokenString, *configAuth)

	if err != nil {
		return ctx.Status(http.StatusUnauthorized).
			JSON(jsonutil.ErrorResponse("Invalid token"))
	}

	return ctx.Next()

}
