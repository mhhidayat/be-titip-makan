package jsonutil

type response[T any] struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func ErrorResponse(message string) response[string] {
	return response[string]{
		Status:  false,
		Message: message,
		Data:    "",
	}
}

func SuccessResponse[T any](message string, data T) response[T] {
	return response[T]{
		Status:  true,
		Message: message,
		Data:    data,
	}
}
