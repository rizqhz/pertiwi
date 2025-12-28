package mutator

import (
	. "math/rand/v2"

	. "github.com/rizqhz/pertiwi/genetic/chromosome"
)

type flipbit struct {
	rng *Rand
}

func FlipBit() *flipbit {
	return &flipbit{}
}

func (m *flipbit) Mutate(c Repr, rate float64) {
	genes := c.Genes().([]int)
	for i := range genes {
		if m.rng.Float64() < rate {
			genes[i] ^= 1
		}
	}
}

func (m *flipbit) Random(r *Rand) {
	m.rng = r
}
