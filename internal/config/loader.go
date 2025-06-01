package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Get() *Config {

	err := godotenv.Load()

	if err != nil {
		panic("Error loader configuration" + err.Error())
	}

	return &Config{
		Server: Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Database: Database{
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			Name: os.Getenv("DB_NAME"),
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
			Tz:   os.Getenv("DB_TZ"),
		},
		Auth: Auth{
			JwtScret: os.Getenv("JWT_SECRET"),
			JwtET:    os.Getenv("JWT_ET"),
		},
	}

}
