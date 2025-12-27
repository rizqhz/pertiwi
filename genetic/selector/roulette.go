package selector

import (
	. "math/rand/v2"

	. "github.com/rizqhz/pertiwi/genetic/chromosome"
)

type roulette struct {
	rng *Rand
}

func Roulette() *roulette {
	return &roulette{}
}

func (s *roulette) Select(p Set) Repr {
	var sum float64
	for _, c := range p {
		sum += c.Score()
	}
	pick := s.rng.Float64() * sum
	var curr float64
	for _, c := range p {
		curr += c.Score()
		if curr >= pick {
			return c
		}
	}
	return p[len(p)-1]
}

func (s *roulette) Random(r *Rand) {
	s.rng = r
}
