package initializer

import (
	. "math/rand/v2"

	"github.com/rizqhz/pertiwi/genetic/chromosome"
)

type integer struct {
	rng  *Rand
	size int
	len  int
	k    int
}

func Integer(k int) *integer {
	return &integer{
		k: k,
	}
}

func (p *integer) Populate() chromosome.Set {
	s := make(chromosome.Set, p.size)
	for i := range s {
		s[i] = chromosome.Integer(p.rng, p.len, p.k)
	}
	return s
}

func (p *integer) Empty() chromosome.Set {
	return make(chromosome.Set, p.size)
}

func (p *integer) Random(r *Rand) { p.rng = r }
func (p *integer) Size(n int)     { p.size = n }
func (p *integer) Length(n int)   { p.len = n }
