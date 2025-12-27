package chromosome

import (
	"cmp"
	"slices"
)

type Repr interface {
	Clone() Repr
	Genes() any
	Get(i int) any
	Set(i int, v any)
	Score() float64
	Update(score float64)
	Len() int
}

type Set []Repr

func (s Set) Best(n int) Set {
	x := slices.Clone(s)
	x.SortDesc()
	return x[:min(n, len(x))]
}

func (s Set) SortAsc() {
	slices.SortFunc(s, func(a, b Repr) int {
		return cmp.Compare(a.Score(), b.Score())
	})
}

func (s Set) SortDesc() {
	slices.SortFunc(s, func(a, b Repr) int {
		return cmp.Compare(b.Score(), a.Score())
	})
}
