package generator

import (
	"math"
	"time"
)

// LoggedOnCustomers demo generator, simulates a trend of customers
type LoggedOnCustomers struct {
	sin Sin
	rand Rand
}

func NewLoggedOnCustomers() LoggedOnCustomers {
	return LoggedOnCustomers{
		sin: *NewSin(15*time.Minute, 50),
		rand: Rand{ Max: 20 },
	}
}

func (r LoggedOnCustomers) NextVal() float64 {
	val := r.sin.NextVal() + 15
	offset := r.rand.NextVal()
	if val <= 0 {
		return offset
	} 
	return math.Round(val + offset)
}