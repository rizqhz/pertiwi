package strategy

type fixed struct {
	rate float64
}

func Fixed(rate float64) *fixed {
	return &fixed{rate}
}

func (s *fixed) Rate() float64   { return s.rate }
func (s *fixed) Adapt(m Metrics) { s.rate = m.Rate }
