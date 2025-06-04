package auth

import (
	"be-titip-makan/configs"
	"be-titip-makan/internal/jsonutil"
	"be-titip-makan/internal/jwtutil"
	"be-titip-makan/internal/user"
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	authService AuthService
	configAuth  configs.Auth
}

func NewAuth(router fiber.Router, authService AuthService, configAuth configs.Auth) {

	ua := authHandler{
		authService: authService,
		configAuth:  configAuth,
	}

	router.Post("login", ua.Login)

}

func (ua authHandler) Login(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	req := user.LoginRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).
			JSON(jsonutil.ErrorResponse("Invalid request format"))
	}

	if req.PhoneNumber == "" || req.Password == "" {
		return c.Status(http.StatusBadRequest).
			JSON(jsonutil.ErrorResponse("Phone Number or password should not be empty"))
	}

	userData, err := ua.authService.Login(ctx, req.PhoneNumber, req.Password)

	if err != nil || userData == nil {
		return c.Status(http.StatusUnauthorized).
			JSON(jsonutil.ErrorResponse("Invalid credentials, please check your username and password"))
	}

	tokenAuth, err := jwtutil.GenerateToken(userData.ID, userData.Name, userData.PhoneNumber, ua.configAuth)

	if err != nil {
		return c.Status(http.StatusInternalServerError).
			JSON(jsonutil.ErrorResponse("Failed to generate token"))
	}

	responseData := map[string]any{
		"users": userData,
		"token": tokenAuth,
	}

	return c.Status(http.StatusOK).
		JSON(jsonutil.SuccessResponse("Login successful", responseData))

}
