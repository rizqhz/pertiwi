package initializer

import (
	. "math/rand/v2"

	"github.com/rizqhz/pertiwi/genetic/chromosome"
)

type bitset struct {
	rng  *Rand
	size int
	len  int
}

func Bitset() *bitset {
	return &bitset{}
}

func (b *bitset) Populate() chromosome.Set {
	s := make(chromosome.Set, b.size)
	for i := range s {
		s[i] = chromosome.Bitset(b.rng, b.len)
	}
	return s
}

func (b *bitset) Empty() chromosome.Set {
	return make(chromosome.Set, b.size)
}

func (b *bitset) Random(r *Rand) { b.rng = r }
func (b *bitset) Size(n int)     { b.size = n }
func (b *bitset) Length(n int)   { b.len = n }
