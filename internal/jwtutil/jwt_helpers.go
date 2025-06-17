package jwtutil

import (
	"be-titip-makan/configs"
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID, name string, phoneNumber string, username string, configAuth configs.Auth) (string, error) {
	ET, _ := strconv.Atoi(configAuth.JwtET)
	expirationTime := time.Now().Add(time.Second * time.Duration(ET))

	type Claims struct {
		UserID      string `json:"id"`
		Username    string `json:"username"`
		Name        string `json:"name"`
		PhoneNumber string `json:"phone_number"`
		jwt.RegisteredClaims
	}

	claims := &Claims{
		UserID:      userID,
		Username:    username,
		Name:        name,
		PhoneNumber: phoneNumber,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(configAuth.JwtScret))
}

func VerifyToken(tokenString string, configAuth configs.Auth) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(configAuth.JwtScret), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("token is not valid")
	}

	return nil
}

func ExtractClaims(tokenStr string, configAuth configs.Auth) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(configAuth.JwtScret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid jwt token")
	}
}
