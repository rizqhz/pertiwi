package initializer

import (
	. "math/rand/v2"

	"github.com/rizqhz/pertiwi/genetic/chromosome"
)

type Initializer interface {
	Populate() chromosome.Set
	Empty() chromosome.Set
	Settings
}

type Settings interface {
	Random(*Rand)
	Size(int)
	Length(int)
}
