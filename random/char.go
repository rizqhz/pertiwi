package random

import (
	"iter"
	"math/rand/v2"
)

type char struct {
	r *rand.Rand
}

func Char(r *rand.Rand) *char {
	return &char{r}
}

func (c *char) Next() rune {
	return c.r.Int32N(95) + 32
}

func (c *char) Fill(dst []byte) {
	for i := range dst {
		dst[i] = byte(c.Next())
	}
}

func (c *char) Iter() iter.Seq[rune] {
	return func(yield func(rune) bool) {
		for {
			if !yield(c.Next()) {
				return
			}
		}
	}
}

func (c *char) IterN(n int) iter.Seq[rune] {
	return func(yield func(rune) bool) {
		for range n {
			if !yield(c.Next()) {
				return
			}
		}
	}
}
