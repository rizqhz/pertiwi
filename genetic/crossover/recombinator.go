package crossover

import (
	. "math/rand/v2"

	. "github.com/rizqhz/pertiwi/genetic/chromosome"
)

type Recombinator interface {
	Combine(p1, p2 Repr, rate float64) (c1, c2 Repr)
	Settings
}

type Settings interface {
	Random(*Rand)
}
