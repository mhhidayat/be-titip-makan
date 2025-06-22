package order

type OrderRequest struct {
	UserID        string `json:"user_id" validate:"required"`
	PaymentStatus string `json:"payment_status" validate:"required"`
	TotalAmount   string `json:"total_amount" validate:"required"`
	PaymentMethod string `json:"payment_method" validate:"required"`
}

type CreateOrder struct {
	OrderNumber   string `db:"order_number"`
	UserID        string `db:"user_id"`
	PaymentStatus string `db:"payment_status"`
	TotalAmount   string `db:"total_amount"`
	PaymentMethod string `db:"payment_method"`
}
