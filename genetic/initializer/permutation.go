package initializer

import (
	"github.com/rizqhz/pertiwi/genetic/chromosome"
)

type permutation struct{}

func Permutation() *permutation {
	return &permutation{}
}

func (p *permutation) Populate(r *Rand, n Size, k Len) Set {
	s := make(Set, n)
	for i := range s {
		s[i] = chromosome.Permutation(r, k)
	}
	return s
}

func (p *permutation) Empty(r *Rand, n Size) Set {
	return make(Set, n)
}
