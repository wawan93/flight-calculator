package calculator

import (
	"encoding/csv"
	"errors"
	"math/rand"
	"os"
	"reflect"
	"testing"
	"time"

	"flights-test/internal/entity"
)

func TestCalculator_Calculate(t *testing.T) {
	testCases := map[string]struct {
		data           []entity.Flight
		expectedError  error
		expectedResult entity.Flight
	}{
		"success": {
			[]entity.Flight{
				{"IND", "EWR"},
				{"SFO", "ATL"},
				{"ATL", "GSO"},
				{"GSO", "IND"},
			},
			nil,
			entity.Flight{"SFO", "EWR"},
		},
		"incorrect path": {
			[]entity.Flight{
				{"SFO", "ATL"},
				{"ATL", "GSO"},
				{"GSO", "IND"},
				{"GSO", "IND"},
				{"IND", "EWR"},
			},
			ErrIncorrectPath,
			entity.Flight{},
		},
	}

	rand.NewSource(time.Now().UnixNano())

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			svc := Calculator{}
			rand.Shuffle(len(tc.data), func(i, j int) {
				tc.data[i], tc.data[j] = tc.data[j], tc.data[i]
			})

			res, err := svc.Calculate(tc.data)
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("unexpected error: want %v, got %v", tc.expectedError, err)
			}
			if !reflect.DeepEqual(res, tc.expectedResult) {
				t.Errorf("unexpected result: want %v, got %v", tc.expectedResult, res)
			}
		})
	}
}

func BenchmarkCalculator_Calculate(b *testing.B) {
	f, err := os.Open("airports.csv")
	if err != nil {
		b.Fatal(err)
	}
	fileReader := csv.NewReader(f)
	airportsCSV, err := fileReader.ReadAll()
	if err != nil {
		b.Fatal(err)
	}

	rand.NewSource(time.Now().UnixNano())

	// The max possible number of flights:
	// for someone who born in some airport, travel every day for the whole 120-years life without any breaks,
	// it's avg flight takes 30 minutes and 30 minutes waiting in airport (1 hour per flight in total)
	const MAX_PATHS = 24 * 365 * 120

	start := airportsCSV[rand.Intn(len(airportsCSV))][0]

	flights := make([]entity.Flight, 0, MAX_PATHS)
	for i := 0; i < MAX_PATHS; i++ {
		arrive := airportsCSV[rand.Intn(len(airportsCSV))][0]
		for arrive == start {
			arrive = airportsCSV[rand.Intn(len(airportsCSV))][0]
		}
		flights = append(flights, entity.Flight{
			start,
			arrive,
		})
		start = arrive
	}

	rand.Shuffle(len(flights), func(i, j int) {
		flights[i], flights[j] = flights[j], flights[i]
	})

	calc := Calculator{}

	for i := 0; i < b.N; i++ {
		_, err := calc.Calculate(flights)
		if err != nil {
			b.Fatal(err)
		}
	}

}
