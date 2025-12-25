package chromosome

import (
	"slices"
	"sort"
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

func (s Set) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Set) Less(i, j int) bool { return s[i].Score() > s[j].Score() }
func (s Set) Size() int          { return len(s) }
func (s Set) Clone() Set         { return slices.Clone(s) }

func (s Set) Best(n int) Set {
	x := s.Clone()
	sort.Slice(x, x.Less)
	return x[:n]
}
