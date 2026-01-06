package selector

import "math"

type boltzmann struct {
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

func (s *boltzmann) Select(p Set, r *Rand) Repr {
	defer func() {
		s.step += 1
		if s.step%s.interval == 0 {
			s.temperature *= s.decay
		}
	}()

	max := math.Inf(-1)
	for _, c := range p {
		max = math.Max(c.Score(), max)
	}

	weights := make([]float64, len(p))
	var sum float64
	for i, c := range p {
		weights[i] = math.Exp((c.Score() - max) / s.temperature)
		sum += weights[i]
	}

	pick := r.Float64() * sum
	var curr float64
	for i, c := range p {
		curr += weights[i]
		if curr >= pick {
			return c
		}
	}

	return p[len(p)-1]
}
