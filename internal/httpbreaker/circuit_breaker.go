package httpbreaker

import (
	"time"

	"github.com/sony/gobreaker"
)

func NewCircuitBreaker() *gobreaker.CircuitBreaker {
	return gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "HTTP Client",
		Timeout: time.Second * 30,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return counts.Requests >= 3 && failureRatio >= 0.6
		},
		// OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
		// 	do smth when circuit breaker trips.
		// },
	})
}
