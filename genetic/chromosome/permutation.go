package chromosome

import (
	"math/rand/v2"
	"strconv"
	"strings"
)

type permutation struct {
	genes []int
	score float64
}

func Permutation(r *rand.Rand, n int) *permutation {
	return &permutation{
		genes: r.Perm(n),
		score: 0,
	}
}

func (repr *permutation) Clone() Repr {
	dup := &permutation{
		genes: make([]int, repr.Len()),
		score: repr.score,
	}
	copy(dup.genes, repr.genes)
	return dup
}

func (repr *permutation) Genes() any           { return repr.genes }
func (repr *permutation) Get(i int) any        { return repr.genes[i] }
func (repr *permutation) Set(i int, v any)     { repr.genes[i] = v.(int) }
func (repr *permutation) Score() float64       { return repr.score }
func (repr *permutation) Update(score float64) { repr.score = score }
func (repr *permutation) Len() int             { return len(repr.genes) }

func (repr *permutation) String() string {
	var b strings.Builder
	b.Grow(repr.Len())
	for i := range repr.genes {
		b.WriteString(strconv.Itoa(repr.genes[i]))
		b.WriteRune(' ')
	}
	return b.String()
}
