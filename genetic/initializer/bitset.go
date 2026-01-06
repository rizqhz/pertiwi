package initializer

import (
	"github.com/rizqhz/pertiwi/genetic/chromosome"
)

type bitset struct{}

func Bitset() *bitset {
	return &bitset{}
}

func (p *bitset) Populate(r *Rand, n Size, k Len) Set {
	s := make(Set, n)
	for i := range s {
		s[i] = chromosome.Bitset(r, k)
	}
	return s
}

func (p *bitset) Empty(r *Rand, n Size) Set {
	return make(Set, n)
}
