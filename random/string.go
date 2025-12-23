package random

import (
	"iter"
	"math/rand/v2"
)

type chars struct {
	r *rand.Rand
	n int
}

func String(r *rand.Rand, n int) *chars {
	return &chars{r, n}
}

func (s *chars) Next() string {
	str := make([]byte, s.n)
	for i := range str {
		str[i] = byte(s.r.IntN(95)) + 32
	}
	return string(str)
}

func (s *chars) Fill(dst []byte) {
	for i := range dst {
		dst[i] = byte(s.r.IntN(95)) + 32
	}
}

func (s *chars) Iter() iter.Seq[string] {
	return func(yield func(string) bool) {
		for {
			if !yield(s.Next()) {
				return
			}
		}
	}
}

func (s *chars) IterN(n int) iter.Seq[string] {
	return func(yield func(string) bool) {
		for range n {
			if !yield(s.Next()) {
				return
			}
		}
	}
}
