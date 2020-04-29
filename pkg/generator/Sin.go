package generator

import (
	"math"
	"time"
)

// Sin implements Generator
type Sin struct {
	ts time.Time
	Freq float64
}

// NewSin creates a new Sin generator with initial timestamp
func NewSin(frequency float64) *Sin {
	return &Sin{
		ts: time.Now(),
		Freq: frequency,
	}
}

// NextVal generates values from a sin function
func (s *Sin) NextVal() float64 {
	return math.Sin(s.radians())
}

func (s *Sin) radians() float64 {
	now := time.Now()
	dur := now.Sub(s.ts)
	radSec := s.Freq * math.Pi * 2.0
	rad := radSec * dur.Seconds()
	s.ts = now
	return rad
}