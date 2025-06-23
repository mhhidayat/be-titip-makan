package order

type OrderRequest struct {
	UserID        string               `json:"user_id" validate:"required"`
	PaymentStatus string               `json:"payment_status" validate:"required"`
	TotalAmount   string               `json:"total_amount" validate:"required"`
	PaymentMethod string               `json:"payment_method" validate:"required"`
	Detail        []OrderDetailRequest `json:"detail" validate:"required,dive"`
}

type OrderDetailRequest struct {
	MenuID      string `json:"menu_id" validate:"required"`
	Qty         string `json:"qty" validate:"required"`
	Price       string `json:"price" validate:"required"`
	Description string `json:"description"`
}

type CreateOrder struct {
	OrderNumber   string `db:"order_number" json:"order_number"`
	UserID        string `db:"user_id" json:"user_id"`
	PaymentStatus string `db:"payment_status" json:"payment_status"`
	TotalAmount   string `db:"total_amount" json:"total_amount"`
	PaymentMethod string `db:"payment_method" json:"payment_method"`
}

type CreateOrderDetail struct {
	OrderNumber string `db:"order_number" json:"order_number"`
	MenuID      string `db:"menu_id" json:"menu_id"`
	Qty         string `db:"qty" json:"qty"`
	Price       string `db:"price" json:"price"`
	Description string `db:"description" json:"description"`
}

type CreateOrderData struct {
	CreateOrder CreateOrder         `json:"create_order"`
	Detail      []CreateOrderDetail `json:"detail"`
}

type DeleteDetailOrder struct {
	OrderNumber string `json:"order_number" validate:"required"`
	MenuID      string `json:"menu_id" validate:"required"`
}
