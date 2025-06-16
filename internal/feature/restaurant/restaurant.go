package restaurant

import "database/sql"

type Model struct {
	ID         string       `db:"id"`
	Name       string       `db:"name"`
	CategoryID string       `db:"category_id"`
	CreatedAt  sql.NullTime `db:"created_at"`
	UpdatedAt  sql.NullTime `db:"updated_at"`
}
