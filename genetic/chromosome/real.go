package chromosome

import (
	"math/rand/v2"
	"strconv"
	"strings"
)

type real struct {
	genes []float64
	score float64
}

func Real(r *rand.Rand, n int, min, max float64) *real {
	new := &real{
		genes: make([]float64, n),
		score: 0,
	}
	for i := range n {
		new.genes[i] = min + (max+min)*r.Float64()
	}
	return new
}

func (repr *real) Clone() Repr {
	dup := &real{
		genes: make([]float64, repr.Len()),
		score: repr.score,
	}
	copy(dup.genes, repr.genes)
	return dup
}

func (repr *real) Genes() any           { return repr.genes }
func (repr *real) Get(i int) any        { return repr.genes[i] }
func (repr *real) Set(i int, v any)     { repr.genes[i] = v.(float64) }
func (repr *real) Score() float64       { return repr.score }
func (repr *real) Update(score float64) { repr.score = score }
func (repr *real) Len() int             { return len(repr.genes) }

func (repr *real) String() string {
	var b strings.Builder
	b.Grow(repr.Len())
	for i := range repr.genes {
		b.WriteString(strconv.FormatFloat(repr.genes[i], 'f', 5, 64))
		b.WriteRune(' ')
	}
	return b.String()
}
