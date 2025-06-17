package user

import (
	"be-titip-makan/configs"
	"be-titip-makan/internal/jsonutil"
	"be-titip-makan/internal/jwtutil"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	configAuth configs.Auth
}

func NewUser(router fiber.Router, configAuth configs.Auth) {

	ua := userHandler{
		configAuth: configAuth,
	}

	router.Get("/user", ua.GetUserDetail)

}

func (ua userHandler) GetUserDetail(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	dataClaims, err := jwtutil.ExtractClaims(tokenString, ua.configAuth)

	if err != nil {
		return c.Status(http.StatusUnauthorized).
			JSON(jsonutil.ErrorResponse("Invalid token"))
	}

	usersData := UsersData{
		ID:          fmt.Sprintf("%v", dataClaims["id"]),
		Username:    fmt.Sprintf("%v", dataClaims["username"]),
		Name:        fmt.Sprintf("%v", dataClaims["name"]),
		PhoneNumber: fmt.Sprintf("%v", dataClaims["phone_number"]),
	}

	responseData := map[string]UsersData{
		"user": usersData,
	}

	return c.Status(http.StatusOK).
		JSON(jsonutil.SuccessResponse("Succes get user data", responseData))
}
