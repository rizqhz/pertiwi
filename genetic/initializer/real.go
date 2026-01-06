package initializer

import (
	"github.com/rizqhz/pertiwi/genetic/chromosome"
)

type real struct {
	a float64
	b float64
}

func Real(min, max float64) *real {
	return &real{min, max}
}

func (p *real) Populate(r *Rand, n Size, k Len) Set {
	s := make(Set, n)
	for i := range s {
		s[i] = chromosome.Real(r, k, p.a, p.b)
	}
	return s
}

func (p *real) Empty(r *Rand, n Size) Set {
	return make(Set, n)
}
