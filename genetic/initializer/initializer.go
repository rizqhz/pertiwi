package initializer

import (
	"math/rand/v2"

	"github.com/rizqhz/pertiwi/genetic/chromosome"
)

type (
	Rand = rand.Rand
	Set  = chromosome.Set
	Size = int
	Len  = int
)

type Initializer interface {
	Populate(*Rand, Size, Len) Set
	Empty(*Rand, Size) Set
}
