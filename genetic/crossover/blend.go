package crossover

import (
	. "math/rand/v2"

	. "github.com/rizqhz/pertiwi/genetic/chromosome"
)

type blx struct {
	rng   *Rand
	alpha float64
}

func Blend(alpha float64) *blx {
	return &blx{
		alpha: alpha,
	}
}

func (c *blx) Combine(p1, p2 Repr, rate float64) (c1, c2 Repr) {
	c1 = p1.Clone()
	c2 = p2.Clone()
	if n := p1.Len(); c.rng.Float64() < rate {
		for i := range n {
			x := p1.Get(i).(float64)
			y := p2.Get(i).(float64)
			if x > y {
				x, y = y, x
			}
			diff := y - x
			lo := x - (c.alpha * diff)
			hi := y + (c.alpha * diff)
			c1.Set(i, lo+c.rng.Float64()*(hi-lo))
			c2.Set(i, lo+c.rng.Float64()*(hi-lo))
		}
	}
	return c1, c2
}

func (c *blx) Random(r *Rand) {
	c.rng = r
}
