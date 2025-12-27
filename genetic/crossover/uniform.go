package crossover

import (
	. "math/rand/v2"

	. "github.com/rizqhz/pertiwi/genetic/chromosome"
)

type uniform struct {
	rng *Rand
}

func Uniform() *uniform {
	return &uniform{}
}

func (c *uniform) Combine(p1, p2 Repr, rate float64) (c1, c2 Repr) {
	c1 = p1.Clone()
	c2 = p2.Clone()
	if n := p1.Len(); c.rng.Float64() < rate {
		for i := range n {
			if c.rng.Float64() < 0.5 {
				c1.Set(i, p2.Get(i))
				c2.Set(i, p1.Get(i))
			}
		}
	}
	return c1, c2
}

func (c *uniform) Random(r *Rand) {
	c.rng = r
}
