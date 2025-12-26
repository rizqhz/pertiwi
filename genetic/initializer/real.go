package initializer

import (
	. "math/rand/v2"

	"github.com/rizqhz/pertiwi/genetic/chromosome"
)

type real struct {
	rng  *Rand
	size int
	len  int
	min  float64
	max  float64
}

func Real(min, max float64) *real {
	return &real{
		min: min,
		max: max,
	}
}

func (p *real) Populate() chromosome.Set {
	s := make(chromosome.Set, p.size)
	for i := range s {
		s[i] = chromosome.Real(p.rng, p.len, p.min, p.max)
	}
	return s
}

func (p *real) Empty() chromosome.Set {
	return make(chromosome.Set, p.size)
}

func (p *real) Random(r *Rand) { p.rng = r }
func (p *real) Size(n int)     { p.size = n }
func (p *real) Length(n int)   { p.len = n }
