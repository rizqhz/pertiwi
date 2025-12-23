package random

import (
	"iter"
	"math/rand/v2"
)

type integer struct {
	r *rand.Rand
	n int
}

func Integer(r *rand.Rand, n int) *integer {
	return &integer{r, n}
}

func (z *integer) Next() int {
	return z.r.IntN(z.n)
}

func (z *integer) Fill(dst []int) {
	for i := range dst {
		dst[i] = z.Next()
	}
}

func (z *integer) Iter() iter.Seq[int] {
	return func(yield func(int) bool) {
		for {
			if !yield(z.Next()) {
				return
			}
		}
	}
}

func (z *integer) IterN(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for range n {
			if !yield(z.Next()) {
				return
			}
		}
	}
}
