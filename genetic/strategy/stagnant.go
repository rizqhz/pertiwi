package strategy

type stagnant struct {
	rate  float64
	base  float64
	score float64
	gen   int
	limit int
	slop  float64
}

func Stagnant(threshold int, base float64) *stagnant {
	return &stagnant{
		rate:  base,
		base:  base,
		limit: threshold,
		slop:  0.001,
	}
}

func (s *stagnant) Rate() float64 {
	return s.rate
}

func (s *stagnant) Adapt(m Metrics) {
	if m.BestScore > s.score {
		s.rate = s.base
		s.score = m.BestScore
		s.gen = 0
		return
	}
	s.gen += 1
	if s.gen >= s.limit {
		excess := 1 + float64(s.gen-s.limit)
		s.rate = min(s.base+excess*s.slop, 0.1)
	}
}
