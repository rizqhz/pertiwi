package random

import (
	"iter"
	"math/rand/v2"
)

type float struct {
	r *rand.Rand
}

func Float(r *rand.Rand) *float {
	return &float{r}
}

func (f *float) Next() float64 {
	return f.r.Float64()
}

func (f *float) Fill(dst []float64) {
	for i := range dst {
		dst[i] = f.Next()
	}
}

func (f *float) Iter() iter.Seq[float64] {
	return func(yield func(float64) bool) {
		for {
			if !yield(f.Next()) {
				return
			}
		}
	}
}

func (f *float) IterN(n int) iter.Seq[float64] {
	return func(yield func(float64) bool) {
		for range n {
			if !yield(f.Next()) {
				return
			}
		}
	}
}
