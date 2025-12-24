package random

import (
	"iter"
	"math/rand/v2"
)

type point struct {
	r *rand.Rand
	n int
}

func Point(r *rand.Rand, n int) *point {
	return &point{r, n}
}

func (p *point) Next() []float64 {
	k := make([]float64, p.n)
	for i := range k {
		k[i] = p.r.Float64()*2 - 1
	}
	return k
}

func (p *point) Fill(dst []float64) {
	for i := range dst {
		dst[i] = p.r.Float64()*2 - 1
	}
}

func (p *point) Iter() iter.Seq[[]float64] {
	return func(yield func([]float64) bool) {
		for {
			if !yield(p.Next()) {
				return
			}
		}
	}
}

func (p *point) IterN(n int) iter.Seq[[]float64] {
	return func(yield func([]float64) bool) {
		for range n {
			if !yield(p.Next()) {
				return
			}
		}
	}
}
