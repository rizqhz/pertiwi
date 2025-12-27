package selector

import (
	. "math/rand/v2"
	. "slices"

	. "github.com/rizqhz/pertiwi/genetic/chromosome"
)

type rank struct {
	rng *Rand
}

func Rank() *rank {
	return &rank{}
}

func (s *rank) Select(p Set) Repr {
	set, n := Clone(p), len(p)
	set.SortAsc()
	pick := s.rng.Float64() * float64(n*(n+1)/2)
	var curr int
	for i, c := range set {
		curr += i + 1
		if float64(curr) >= pick {
			return c
		}
	}
	return set[n-1]
}

func (s *rank) Random(r *Rand) {
	s.rng = r
}
