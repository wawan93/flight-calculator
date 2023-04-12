package calculator

import (
	"errors"
	"fmt"

	"flights-test/internal/entity"
)

var ErrIncorrectPath = errors.New("incorrect path")
var ErrIncorrectAirportCode = errors.New("incorrect airport code")

type Destination bool

func (d Destination) String() string {
	if d == DEPART {
		return "DEPARTURE"
	}
	return "ARRIVAL"
}

const (
	DEPART Destination = true
	ARRIVE Destination = false
)

type Calculator struct{}

func (s *Calculator) Calculate(flights []entity.Flight) (entity.Flight, error) {
	airports := make(map[string]Destination)
	for _, flight := range flights {
		if !flight.Valid() {
			return entity.Flight{}, ErrIncorrectAirportCode
		}

		if _, ok := airports[flight.Depart()]; ok {
			delete(airports, flight.Depart())
		} else {
			airports[flight.Depart()] = DEPART
		}

		if _, ok := airports[flight.Arrive()]; ok {
			delete(airports, flight.Arrive())
		} else {
			airports[flight.Arrive()] = ARRIVE
		}
	}

	if len(airports) != 2 {
		return entity.Flight{}, fmt.Errorf("%w: %v", ErrIncorrectPath, airports)
	}

	var start, end string
	for key, direction := range airports {
		if direction == DEPART {
			start = key
		}
		if direction == ARRIVE {
			end = key
		}
	}

	return entity.Flight{start, end}, nil
}
