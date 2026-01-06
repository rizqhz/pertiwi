package selector

import "slices"

type rank struct{}

func Rank() *rank {
	return &rank{}
}

func (s *rank) Select(p Set, r *Rand) Repr {
	dup := slices.Clone(p)
	dup.SortAsc()

	n := len(p)
	pick := r.Float64() * float64(n*(n+1)/2)

	var curr int
	for i, c := range dup {
		curr += i + 1
		if float64(curr) >= pick {
			return c
		}
	}

	return dup[n-1]
}
