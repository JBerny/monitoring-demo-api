package generator

import (
	"math"
	"time"
)

// LoggedOnCustomers demo generator, simulates a trend of customers
// implements Generator interface
type LoggedOnCustomers struct {
	sin Sin
	rand Rand
}

// NewLoggedOnCustomers creates a LoggedOnCustomers struct
// initialized with a generator.Sin and a generator.Rand 
func NewLoggedOnCustomers() LoggedOnCustomers {
	return LoggedOnCustomers{
		sin: *NewSin(15*time.Minute, 50),
		rand: Rand{ Max: 20 },
	}
}

// NextVal returns next value 
func (r LoggedOnCustomers) NextVal() float64 {
	val := r.sin.NextVal() + 15
	offset := r.rand.NextVal()
	if val <= 0 {
		return offset
	} 
	return math.Round(val + offset)
}

// APIRequestDuration demo generator, 
// represents the total time in seconds that it takes to the api to fulfill a request
// implements Generator interface
type APIRequestDuration struct {
	count   float64
}

// NextVal returns next value 
func (a APIRequestDuration) NextVal() float64 {
	return 32 + math.Floor(100*math.Cos(a.count*0.11))/10
}

// ServiceRequestDuration demo generator, 
// represents the total time in seconds that it takes to the api to fulfill a request
// implements Generator interface
type ServiceRequestDuration struct {
	count   float64
}

// NextVal returns next value 
func (a ServiceRequestDuration) NextVal() float64 {
	a.count = a.count + 1.0
	return 30 + math.Floor(120*math.Sin(a.count*0.1))/10
}
