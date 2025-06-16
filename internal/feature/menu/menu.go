package menu

import "database/sql"

type Model struct {
	ID           string       `db:"id"`
	Name         string       `db:"name"`
	Price        string       `db:"price"`
	RestaurantID string       `db:"restaurant_id"`
	CreatedAt    sql.NullTime `db:"created_at"`
	UpdatedAt    sql.NullTime `db:"updated_at"`
}
