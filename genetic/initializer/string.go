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

func (p *chars) Populate() chromosome.Set {
	s := make(chromosome.Set, p.size)
	for i := range s {
		s[i] = chromosome.String(p.rng, p.len)
	}
	return s
}

func (p *chars) Empty() chromosome.Set {
	return make(chromosome.Set, p.size)
}

func (p *chars) Random(r *Rand) { p.rng = r }
func (p *chars) Size(n int)     { p.size = n }
func (p *chars) Length(n int)   { p.len = n }
