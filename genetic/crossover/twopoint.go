package crossover

import (
	. "math/rand/v2"

	. "github.com/rizqhz/pertiwi/genetic/chromosome"
)

type twopoint struct {
	rng *Rand
}

func TwoPoint() *twopoint {
	return &twopoint{}
}

func (c *twopoint) Combine(p1, p2 Repr, rate float64) (c1, c2 Repr) {
	c1 = p1.Clone()
	c2 = p2.Clone()
	if n := p1.Len(); c.rng.Float64() < rate {
		a := c.rng.IntN(n)
		b := (c.rng.IntN(n-1) + a + 1) % n
		if a > b {
			a, b = b, a
		}
		for i := a; i < b; i++ {
			c1.Set(i, p2.Get(i))
			c2.Set(i, p1.Get(i))
		}
	}
	return c1, c2
}

func (c *twopoint) Random(r *Rand) {
	c.rng = r
}
