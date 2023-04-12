package calculate

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"flights-test/internal/entity"
	"flights-test/internal/services/calculator"
)

type calculatorService interface {
	Calculate([]entity.Flight) (entity.Flight, error)
}

type Calculate struct {
	svc calculatorService
}

func NewCalculateController(svc calculatorService) *Calculate {
	return &Calculate{svc: svc}
}

func (c *Calculate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var req []entity.Flight
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()}); err != nil {
			log.Println(err)
		}
		return
	}

	resp, err := c.svc.Calculate(req)
	if err != nil {
		if errors.Is(err, calculator.ErrIncorrectPath) {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}

		if err := json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()}); err != nil {
			log.Println(err)
		}
		return
	}

	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		log.Println(err)
	}
}
