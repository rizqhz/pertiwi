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

func (r *real) Populate() chromosome.Set {
	s := make(chromosome.Set, r.size)
	for i := range s {
		s[i] = chromosome.Real(r.rng, r.len, r.min, r.max)
	}
	return s
}

func (r *real) Empty() chromosome.Set {
	return make(chromosome.Set, r.size)
}

func (f *real) Random(r *Rand) { f.rng = r }
func (r *real) Size(n int)     { r.size = n }
func (r *real) Length(n int)   { r.len = n }
