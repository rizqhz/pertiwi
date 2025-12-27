package crossover

import (
	. "math/rand/v2"

	. "github.com/rizqhz/pertiwi/genetic/chromosome"
)

type arithmetic struct {
	rng *Rand
}

func Arithmetic() *arithmetic {
	return &arithmetic{}
}

func (c *arithmetic) Combine(p1, p2 Repr, rate float64) (c1, c2 Repr) {
	c1 = p1.Clone()
	c2 = p2.Clone()
	if n := p1.Len(); c.rng.Float64() < rate {
		alpha := c.rng.Float64()
		for i := range n {
			x := p1.Get(i).(float64)
			y := p2.Get(i).(float64)
			c1.Set(i, alpha*x+(1-alpha)*y)
			c2.Set(i, alpha*y+(1-alpha)*x)
		}
	}
	return c1, c2
}

func (c *arithmetic) Random(r *Rand) {
	c.rng = r
}
