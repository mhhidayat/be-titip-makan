package api

import (
	"be-titip-makan/domain"
	"be-titip-makan/domain/dto"
	"be-titip-makan/domain/dto/request"
	"be-titip-makan/internal/config"
	"be-titip-makan/internal/util"
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type authApi struct {
	usersService domain.UserService
	configAuth   config.Auth
}

func NewAuth(router fiber.Router, usersService domain.UserService, configAuth config.Auth) {

	ua := authApi{
		usersService: usersService,
		configAuth:   configAuth,
	}

	router.Post("login", ua.Login)

}

func (ua authApi) Login(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	req := request.LoginRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(dto.ErrorResponse("Invalid request format"))
	}

	if req.PhoneNumber == "" || req.Password == "" {
		return ctx.Status(http.StatusBadRequest).
			JSON(dto.ErrorResponse("Phone Number or password should not be empty"))
	}

	userData, err := ua.usersService.Login(c, req.PhoneNumber, req.Password)

	if err != nil || userData == nil {
		return ctx.Status(http.StatusUnauthorized).
			JSON(dto.ErrorResponse("Invalid credentials, please check your username and password"))
	}

	tokenAuth, err := util.GenerateToken(userData.ID, userData.Name, userData.PhoneNumber, ua.configAuth)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.ErrorResponse("Failed to generate token"))
	}

	responseData := map[string]any{
		"users": userData,
		"token": tokenAuth,
	}

	return ctx.Status(http.StatusOK).
		JSON(dto.SuccessResponse("Login successful", responseData))

}
