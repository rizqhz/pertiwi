package crossover

import (
	. "math/rand/v2"

	. "github.com/rizqhz/pertiwi/genetic/chromosome"
)

type onepoint struct {
	rng *Rand
}

func OnePoint() *onepoint {
	return &onepoint{}
}

func (c *onepoint) Combine(p1, p2 Repr, rate float64) (c1, c2 Repr) {
	c1 = p1.Clone()
	c2 = p2.Clone()
	if n := p1.Len(); c.rng.Float64() < rate {
		point := c.rng.IntN(n-1) + 1
		for i := point; i < n; i++ {
			c1.Set(i, p2.Get(i))
			c2.Set(i, p1.Get(i))
		}
	}
	return c1, c2
}

func (c *onepoint) Random(r *Rand) {
	c.rng = r
}
