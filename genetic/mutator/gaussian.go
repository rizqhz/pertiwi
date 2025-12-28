package mutator

import (
	. "math/rand/v2"

	. "github.com/rizqhz/pertiwi/genetic/chromosome"
)

type gaussian struct {
	rng    *Rand
	stddev float64
}

func Gaussian(stddev float64) *gaussian {
	return &gaussian{
		stddev: stddev,
	}
}

func (m *gaussian) Mutate(c Repr, rate float64) {
	genes := c.Genes().([]float64)
	for i := range genes {
		if m.rng.Float64() < rate {
			genes[i] += m.stddev * m.rng.NormFloat64()
		}
	}
}

func (m *gaussian) Random(r *Rand) {
	m.rng = r
}
