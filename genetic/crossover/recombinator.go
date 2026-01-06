package crossover

import (
	"math/rand/v2"

	"github.com/rizqhz/pertiwi/genetic/chromosome"
)

type (
	Rand = rand.Rand
	Repr = chromosome.Repr
	Rate = float64
)

type Recombinator interface {
	Combine(Repr, Repr, Rate, *Rand) (Repr, Repr)
}
