package mutator

import (
	. "math/rand/v2"

	. "github.com/rizqhz/pertiwi/genetic/chromosome"
)

type scramble struct {
	rng *Rand
}

func Scramble() *scramble {
	return &scramble{}
}

func (m *scramble) Mutate(c Repr, rate float64) {
	if n := c.Len(); m.rng.Float64() < rate {
		a := m.rng.IntN(n)
		b := m.rng.IntN(n)
		if a > b {
			a, b = b, a
		}
		if a == b {
			return
		}

		interval := b - a + 1
		idx := make([]int, interval)
		val := make([]any, interval)

		for i := range interval {
			idx[i] = a + i
			val[i] = c.Get(a + i)
		}

		m.rng.Shuffle(interval, func(i, j int) {
			val[i], val[j] = val[j], val[i]
		})

		for i := range interval {
			c.Set(idx[i], val[i])
		}
	}
}

func (m *scramble) Random(r *Rand) {
	m.rng = r
}
