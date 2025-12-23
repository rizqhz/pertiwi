package random

import (
	"iter"
	"math/rand/v2"
)

type binary struct {
	r *rand.Rand
}

func Binary(r *rand.Rand) *binary {
	return &binary{r}
}

func (b *binary) Next() int {
	return b.r.Int() & 1
}

func (b *binary) Fill(dst []int) {
	for i := range dst {
		dst[i] = b.Next()
	}
}

func (b *binary) Iter() iter.Seq[int] {
	return func(yield func(int) bool) {
		for {
			if !yield(b.Next()) {
				return
			}
		}
	}
}

func (b *binary) IterN(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for range n {
			if !yield(b.Next()) {
				return
			}
		}
	}
}
