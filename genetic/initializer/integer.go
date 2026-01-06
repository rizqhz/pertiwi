package initializer

import (
	"github.com/rizqhz/pertiwi/genetic/chromosome"
)

type integer struct {
	n int
}

func Integer(n int) *integer {
	return &integer{n}
}

func (p *integer) Populate(r *Rand, n Size, k Len) Set {
	s := make(Set, n)
	for i := range s {
		s[i] = chromosome.Integer(r, k, p.n)
	}
	return s
}

func (p *integer) Empty(r *Rand, n Size) Set {
	return make(Set, n)
}
