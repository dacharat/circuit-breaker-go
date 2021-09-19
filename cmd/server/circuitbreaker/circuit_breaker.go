package circuitbreaker

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sony/gobreaker"
)

var cb *gobreaker.CircuitBreaker

func Init() {
	var st gobreaker.Settings
	st.Name = "HTTP GET"
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && failureRatio >= 0.6
	}

	cb = gobreaker.NewCircuitBreaker(st)
}

// func NewCircuitBreaker() *gobreaker.CircuitBreaker {
// 	return gobreaker.NewCircuitBreaker(gobreaker.Settings{
// 		Name:    "HTTP Client",
// 		Timeout: time.Second * 30,
// 		ReadyToTrip: func(counts gobreaker.Counts) bool {
// 			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
// 			return counts.Requests >= 3 && failureRatio >= 0.6
// 		},
// 		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
// 			// do smth when circuit breaker trips.
// 		},
// 	})
// }

func Get(url string) ([]byte, error) {
	body, err := cb.Execute(func() (interface{}, error) {
		resp, err := http.Get(url)
		fmt.Println("resp: ", resp, err)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != http.StatusOK {
			return body, errors.New("internal server error")
		}

		return body, nil
	})

	if err != nil {
		return nil, err
	}

	return body.([]byte), nil
}
