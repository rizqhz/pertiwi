package chromosome

import (
	"math/rand/v2"
	"strings"
)

type bitset struct {
	genes []int
	score float64
}

func Bitset(r *rand.Rand, n int) *bitset {
	new := &bitset{
		genes: make([]int, n),
		score: 0,
	}
	for i := range n {
		new.genes[i] = r.Int() & 1
	}
	return new
}

func (repr *bitset) Clone() Repr {
	dup := &bitset{
		genes: make([]int, repr.Len()),
		score: repr.score,
	}
	copy(dup.genes, repr.genes)
	return dup
}

func (repr *bitset) Genes() any           { return repr.genes }
func (repr *bitset) Get(i int) any        { return repr.genes[i] }
func (repr *bitset) Set(i int, v any)     { repr.genes[i] = v.(int) }
func (repr *bitset) Score() float64       { return repr.score }
func (repr *bitset) Update(score float64) { repr.score = score }
func (repr *bitset) Len() int             { return len(repr.genes) }

func (repr *bitset) String() string {
	var b strings.Builder
	b.Grow(repr.Len())
	for i := range repr.genes {
		chr := '0' + rune(repr.genes[i])
		b.WriteRune(chr)
	}
	return b.String()
}
