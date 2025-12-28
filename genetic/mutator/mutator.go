package mutator

import (
	. "math/rand/v2"

	. "github.com/rizqhz/pertiwi/genetic/chromosome"
)

type Mutator interface {
	Mutate(c Repr, rate float64)
	Settings
}

type Settings interface {
	Random(*Rand)
}
