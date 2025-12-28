package mutator

import (
	. "math/rand/v2"

	. "github.com/rizqhz/pertiwi/genetic/chromosome"
)

type swap struct {
	rng *Rand
}

func Swap() *swap {
	return &swap{}
}

func (m *swap) Mutate(c Repr, rate float64) {
	if n := c.Len(); m.rng.Float64() < rate {
		i := m.rng.IntN(n)
		j := m.rng.IntN(n)
		v1, v2 := c.Get(i), c.Get(j)
		c.Set(i, v2)
		c.Set(j, v1)
	}
}

func (m *swap) Random(r *Rand) {
	m.rng = r
}
