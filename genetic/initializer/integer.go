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

func (z *integer) Populate() chromosome.Set {
	s := make(chromosome.Set, z.size)
	for i := range s {
		s[i] = chromosome.Integer(z.rng, z.len, z.k)
	}
	return s
}

func (z *integer) Empty() chromosome.Set {
	return make(chromosome.Set, z.size)
}

func (z *integer) Random(r *Rand) { z.rng = r }
func (z *integer) Size(n int)     { z.size = n }
func (z *integer) Length(n int)   { z.len = n }
