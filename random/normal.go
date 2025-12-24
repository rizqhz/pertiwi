package random

import (
	"iter"
	"math/rand/v2"
)

type normal struct {
	r      *rand.Rand
	mean   float64
	stddev float64
}

func Normal(r *rand.Rand, mean, stddev float64) *normal {
	return &normal{r, mean, stddev}
}

func (d *normal) Next() float64 {
	return d.mean + d.r.NormFloat64()*d.stddev
}

func (d *normal) Fill(dst []float64) {
	for i := range dst {
		dst[i] = d.Next()
	}
}

func (d *normal) Iter() iter.Seq[float64] {
	return func(yield func(float64) bool) {
		for {
			if !yield(d.Next()) {
				return
			}
		}
	}
}

func (d *normal) IterN(n int) iter.Seq[float64] {
	return func(yield func(float64) bool) {
		for range n {
			if !yield(d.Next()) {
				return
			}
		}
	}
}
