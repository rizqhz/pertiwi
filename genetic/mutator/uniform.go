package mutator

import (
	. "math/rand/v2"

	. "github.com/rizqhz/pertiwi/genetic/chromosome"
)

type uniform struct {
	rng *Rand
	min float64
	max float64
}

func Uniform(min, max float64) *uniform {
	return &uniform{
		min: min,
		max: max,
	}
}

func (m *uniform) Mutate(c Repr, rate float64) {
	genes := c.Genes().([]float64)
	for i := range genes {
		if m.rng.Float64() < rate {
			genes[i] = m.min + m.rng.Float64()*(m.max-m.min)
		}
	}
}

func (m *uniform) Random(r *Rand) {
	m.rng = r
}
