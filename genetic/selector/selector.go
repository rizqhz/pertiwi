package selector

import (
	"math/rand/v2"

	"github.com/rizqhz/pertiwi/genetic/chromosome"
)

type (
	Rand = rand.Rand
	Repr = chromosome.Repr
	Set  = chromosome.Set
)

type Selector interface {
	Select(Set, *Rand) Repr
}
