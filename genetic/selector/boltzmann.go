package selector

import (
	. "math"
	. "math/rand/v2"

	. "github.com/rizqhz/pertiwi/genetic/chromosome"
)

type boltzmann struct {
	rng         *Rand
	temperature float64
	decay       float64
	interval    int
	step        int
}

func Boltzmann(temperature, decay float64, interval int) *boltzmann {
	return &boltzmann{
		temperature: temperature,
		decay:       decay,
		interval:    interval,
	}
}

func (s *boltzmann) Select(p Set) Repr {
	defer func() {
		s.step += 1
		if s.step%s.interval == 0 {
			s.temperature *= s.decay
		}
	}()

	var max float64 = Inf(-1)
	for _, c := range p {
		max = Max(c.Score(), max)
	}

	weights := make([]float64, len(p))
	var sum float64
	for i, c := range p {
		weights[i] = Exp((c.Score() - max) / s.temperature)
		sum += weights[i]
	}

	pick := s.rng.Float64() * sum
	var curr float64
	for i, c := range p {
		curr += weights[i]
		if curr >= pick {
			return c
		}
	}

	return p[len(p)-1]
}

func (s *boltzmann) Random(r *Rand) {
	s.rng = r
}
