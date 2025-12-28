package crossover

import (
	. "math/rand/v2"

	. "github.com/rizqhz/pertiwi/genetic/chromosome"
)

type pmx struct {
	rng *Rand
}

func PartiallyMapped() *pmx {
	return &pmx{}
}

func (c *pmx) Combine(p1, p2 Repr, rate float64) (c1, c2 Repr) {
	c1 = p1.Clone()
	c2 = p2.Clone()
	if n := p1.Len(); c.rng.Float64() < rate {
		a := c.rng.IntN(n)
		b := (c.rng.IntN(n-1) + a + 1) % n
		if a > b {
			a, b = b, a
		}

		m1 := make(map[any]any, b-a)
		m2 := make(map[any]any, b-a)

		for i := a; i < b; i++ {
			v1 := p1.Get(i)
			v2 := p2.Get(i)

			c1.Set(i, v2)
			c2.Set(i, v1)

			m1[v2] = v1
			m2[v1] = v2
		}

		for i := range n {
			if i >= a && i < b {
				continue
			}

			v1 := c1.Get(i)
			for {
				if replacement, ok := m1[v1]; ok {
					v1 = replacement
				} else {
					break
				}
			}
			c1.Set(i, v1)

			v2 := c2.Get(i)
			for {
				if replacement, ok := m2[v2]; ok {
					v2 = replacement
				} else {
					break
				}
			}
			c2.Set(i, v2)
		}
	}
	return c1, c2
}

func (c *pmx) Random(r *Rand) {
	c.rng = r
}
