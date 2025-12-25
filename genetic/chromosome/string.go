package chromosome

import (
	"math/rand/v2"
)

type chars struct {
	genes []byte
	score float64
}

func String(r *rand.Rand, n int) *chars {
	new := &chars{
		genes: make([]byte, n),
		score: 0,
	}
	for i := range n {
		new.genes[i] = byte(r.IntN(95)) + 32
	}
	return new
}

func (repr *chars) Clone() Repr {
	dup := &chars{
		genes: make([]byte, repr.Len()),
		score: repr.score,
	}
	copy(dup.genes, repr.genes)
	return dup
}

func (repr *chars) Genes() any           { return repr.genes }
func (repr *chars) Get(i int) any        { return repr.genes[i] }
func (repr *chars) Set(i int, v any)     { repr.genes[i] = v.(byte) }
func (repr *chars) Score() float64       { return repr.score }
func (repr *chars) Update(score float64) { repr.score = score }
func (repr *chars) Len() int             { return len(repr.genes) }
func (repr *chars) String() string       { return string(repr.genes) }
