package selector

import (
	. "math/rand/v2"

	. "github.com/rizqhz/pertiwi/genetic/chromosome"
)

type Selector interface {
	Select(Set) Repr
	Settings
}

type Settings interface {
	Random(*Rand)
}
