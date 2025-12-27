package selector

import (
	. "math/rand/v2"

	. "github.com/rizqhz/pertiwi/genetic/chromosome"
)

type tournament struct {
	rng *Rand
	k   int
}

func Tournament(k int) *tournament {
	return &tournament{
		k: k,
	}
}

func (s *tournament) Select(p Set) Repr {
	var b Repr
	for range s.k {
		c := p[s.rng.IntN(len(p))]
		if b == nil || c.Score() > b.Score() {
			b = c
		}
	}
	return b
}

func (s *tournament) Random(r *Rand) {
	s.rng = r
}
