package mutator

import (
	"math/rand/v2"

	"github.com/rizqhz/pertiwi/genetic/chromosome"
)

type (
	Rand = rand.Rand
	Repr = chromosome.Repr
	Rate = float64
)

type Mutator interface {
	Mutate(Repr, Rate, *Rand)
}
