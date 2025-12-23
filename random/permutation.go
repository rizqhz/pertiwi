package random

import (
	"iter"
	"math/rand/v2"
)

type permutation struct {
	r *rand.Rand
	n int
}

func Permutation(r *rand.Rand, n int) *permutation {
	return &permutation{r, n}
}

func (p *permutation) Next() []int {
	return p.r.Perm(p.n)
}

func (p *permutation) Fill(dst []int) {
	for i := range dst {
		dst[i] = i
	}
	p.r.Shuffle(len(dst), func(i, j int) {
		dst[i], dst[j] = dst[j], dst[i]
	})
}

func (p *permutation) Iter() iter.Seq[[]int] {
	return func(yield func([]int) bool) {
		for {
			if !yield(p.Next()) {
				return
			}
		}
	}
}

func (p *permutation) IterN(n int) iter.Seq[[]int] {
	return func(yield func([]int) bool) {
		for range n {
			if !yield(p.Next()) {
				return
			}
		}
	}
}
