package random

import (
	"math/rand/v2"
	"sync/atomic"
)

type Source interface {
	PRNG() *rand.Rand
	rand.Source
}

type impl struct {
	state atomic.Uint64
}

func (s *impl) Uint64() uint64 {
	return SplitMix64(&s.state)
}

func (s *impl) PRNG() *rand.Rand {
	s1 := s.Uint64()
	s2 := s1 | 1
	return rand.New(rand.NewPCG(s1, s2))
}

func New() *impl {
	s := &impl{}
	s.state.Store(rand.Uint64())
	return s
}

func From(seed uint64) *impl {
	s := &impl{}
	s.state.Store(seed)
	return s
}
