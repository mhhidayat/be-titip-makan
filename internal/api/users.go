package api

import (
	"be-titip-makan/domain/user"
	"be-titip-makan/internal/config"
	"be-titip-makan/internal/util"
	"be-titip-makan/internal/util/response"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type usersApi struct {
	usersService user.UserService
	configAuth   config.Auth
}

func NewUsers(router fiber.Router, userService user.UserService, configAuth config.Auth) {

	ua := usersApi{
		usersService: userService,
		configAuth:   configAuth,
	}

	router.Get("/users", ua.FindUsersDetailByPhoneNumber)

}

func (ua usersApi) FindUsersDetailByPhoneNumber(ctx *fiber.Ctx) error {
	tokenString := ctx.Get("Authorization")
	dataClaims, err := util.ExtractClaims(tokenString, ua.configAuth)

	if err != nil {
		return ctx.Status(http.StatusUnauthorized).
			JSON(response.ErrorResponse("Invalid token"))
	}

	usersData := user.UsersData{
		ID:          fmt.Sprintf("%v", dataClaims["id"]),
		Name:        fmt.Sprintf("%v", dataClaims["name"]),
		PhoneNumber: fmt.Sprintf("%v", dataClaims["phone_number"]),
	}

	responseData := map[string]user.UsersData{
		"users": usersData,
	}

	return ctx.Status(http.StatusOK).
		JSON(response.SuccessResponse("Succes get users data", responseData))
}
