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

func (p *bitset) Populate() chromosome.Set {
	s := make(chromosome.Set, p.size)
	for i := range s {
		s[i] = chromosome.Bitset(p.rng, p.len)
	}
	return s
}

func (p *bitset) Empty() chromosome.Set {
	return make(chromosome.Set, p.size)
}

func (p *bitset) Random(r *Rand) { p.rng = r }
func (p *bitset) Size(n int)     { p.size = n }
func (p *bitset) Length(n int)   { p.len = n }
