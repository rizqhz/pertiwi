package chromosome

import (
	"math/rand/v2"
	"strconv"
	"strings"
)

type integer struct {
	genes []int
	score float64
}

func Integer(r *rand.Rand, n, k int) *integer {
	new := &integer{
		genes: make([]int, n),
		score: 0,
	}
	for i := range n {
		new.genes[i] = r.IntN(k)
	}
	return new
}

func (repr *integer) Clone() Repr {
	dup := &integer{
		genes: make([]int, repr.Len()),
		score: repr.score,
	}
	copy(dup.genes, repr.genes)
	return dup
}

func (repr *integer) Genes() any           { return repr.genes }
func (repr *integer) Get(i int) any        { return repr.genes[i] }
func (repr *integer) Set(i int, v any)     { repr.genes[i] = v.(int) }
func (repr *integer) Score() float64       { return repr.score }
func (repr *integer) Update(score float64) { repr.score = score }
func (repr *integer) Len() int             { return len(repr.genes) }

func (repr *integer) String() string {
	var b strings.Builder
	b.Grow(repr.Len())
	for i := range repr.genes {
		b.WriteString(strconv.Itoa(repr.genes[i]))
		b.WriteRune(' ')
	}
	return b.String()
}
