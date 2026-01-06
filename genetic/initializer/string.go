package initializer

import (
	"github.com/rizqhz/pertiwi/genetic/chromosome"
)

type chars struct{}

func String() *chars {
	return &chars{}
}

func (p *chars) Populate(r *Rand, n Size, k Len) Set {
	s := make(Set, n)
	for i := range s {
		s[i] = chromosome.String(r, k)
	}
	return s
}

func (p *chars) Empty(r *Rand, n Size) Set {
	return make(Set, n)
}
