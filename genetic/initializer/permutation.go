package initializer

import (
	. "math/rand/v2"

	"github.com/rizqhz/pertiwi/genetic/chromosome"
)

type permutation struct {
	rng  *Rand
	size int
	len  int
}

func Permutation() *permutation {
	return &permutation{}
}

func (p *permutation) Populate() chromosome.Set {
	s := make(chromosome.Set, p.size)
	for i := range s {
		s[i] = chromosome.Permutation(p.rng, p.len)
	}
	return s
}

func (p *permutation) Empty() chromosome.Set {
	return make(chromosome.Set, p.size)
}

func (p *permutation) Random(r *Rand) { p.rng = r }
func (p *permutation) Size(n int)     { p.size = n }
func (p *permutation) Length(n int)   { p.len = n }
