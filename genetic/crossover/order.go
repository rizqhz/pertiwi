package crossover

import (
	. "math/rand/v2"

	. "github.com/rizqhz/pertiwi/genetic/chromosome"
)

type order struct {
	rng *Rand
}

func Order() *order {
	return &order{}
}

func (c *order) Combine(p1, p2 Repr, rate float64) (c1, c2 Repr) {
	c1 = p1.Clone()
	c2 = p2.Clone()
	if n := p1.Len(); c.rng.Float64() < rate {
		a := c.rng.IntN(n)
		b := c.rng.IntN(n)
		if a > b {
			a, b = b, a
		}

		fill := func(child, parent, donor Repr) {
			used := make(map[any]bool)
			for i := a; i <= b; i++ {
				val := parent.Get(i)
				child.Set(i, val)
				used[val] = true
			}

			pos := (b + 1) % n
			for i := range n {
				idx := (b + 1 + i) % n
				gene := donor.Get(idx)
				if !used[gene] {
					child.Set(pos, gene)
					pos = (pos + 1) % n
				}
			}
		}

		fill(c1, p1, p2)
		fill(c2, p2, p1)
	}
	return c1, c2
}

func (c *order) Random(r *Rand) {
	c.rng = r
}
