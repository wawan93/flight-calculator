package calculate

import "flights-test/internal/entity"

type Request []entity.Flight

type ErrorResponse struct {
	Error string `json:"error"`
}
