package db

import (
	"be-titip-makan/configs"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetDatabase(conf configs.Database) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Name, conf.Tz)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to the database: %v", err))
	}

	if err = db.Ping(); err != nil {
		panic(fmt.Sprintf("Failed to ping the database: %v", err))
	}

	return db
}
