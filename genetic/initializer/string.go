package initializer

import (
	. "math/rand/v2"

	"github.com/rizqhz/pertiwi/genetic/chromosome"
)

type chars struct {
	rng  *Rand
	size int
	len  int
}

func String() *chars {
	return &chars{}
}

func (s *chars) Populate() chromosome.Set {
	p := make(chromosome.Set, s.size)
	for i := range p {
		p[i] = chromosome.String(s.rng, s.len)
	}
	return p
}

func (s *chars) Empty() chromosome.Set {
	return make(chromosome.Set, s.size)
}

func (s *chars) Random(r *Rand) { s.rng = r }
func (s *chars) Size(n int)     { s.size = n }
func (s *chars) Length(n int)   { s.len = n }
