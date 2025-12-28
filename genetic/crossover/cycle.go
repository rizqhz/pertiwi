package crossover

import (
	. "math/rand/v2"

	. "github.com/rizqhz/pertiwi/genetic/chromosome"
)

type cycle struct {
	rng *Rand
}

func Cycle() *cycle {
	return &cycle{}
}

func (c *cycle) Combine(p1, p2 Repr, rate float64) (c1, c2 Repr) {
	c1 = p1.Clone()
	c2 = p2.Clone()
	if n := p1.Len(); c.rng.Float64() < rate {
		pos := make(map[any]int, n)
		for i := range n {
			pos[p1.Get(i)] = i
		}

		cycle := make([]bool, n)
		curr := 0
		for !cycle[curr] {
			cycle[curr] = true
			val := p2.Get(curr)
			curr = pos[val]
		}

		for i := range n {
			if cycle[i] {
				c1.Set(i, p1.Get(i))
				c2.Set(i, p2.Get(i))
			} else {
				c1.Set(i, p2.Get(i))
				c2.Set(i, p1.Get(i))
			}
		}
	}
	return c1, c2
}

func (c *cycle) Random(r *Rand) {
	c.rng = r
}
