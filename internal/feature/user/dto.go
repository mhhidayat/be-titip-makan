package user

type UsersData struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
