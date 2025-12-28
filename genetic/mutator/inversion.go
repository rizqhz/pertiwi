package mutator

import (
	. "math/rand/v2"

	. "github.com/rizqhz/pertiwi/genetic/chromosome"
)

type inversion struct {
	rng *Rand
}

func Inversion() *inversion {
	return &inversion{}
}

func (m *inversion) Mutate(c Repr, rate float64) {
	if n := c.Len(); m.rng.Float64() < rate {
		a := m.rng.IntN(n)
		b := m.rng.IntN(n)
		if a > b {
			a, b = b, a
		}
		for a < b {
			v1 := c.Get(a)
			v2 := c.Get(b)
			c.Set(a, v2)
			c.Set(b, v1)
			a += 1
			b -= 1
		}
	}
}

func (m *inversion) Random(r *Rand) {
	m.rng = r
}
