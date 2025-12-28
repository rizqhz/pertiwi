package evaluator

import (
	. "github.com/rizqhz/pertiwi/genetic/chromosome"
)

type Evaluator interface {
	Evaluate(Repr)
}
