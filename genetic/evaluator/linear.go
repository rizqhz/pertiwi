package evaluator

import (
	. "math"

	. "github.com/rizqhz/pertiwi/genetic/chromosome"
)

type linear struct {
	inputs []float64
	target float64
}

func Linear(y float64, x ...float64) *linear {
	return &linear{x, y}
}

func (e *linear) Evaluate(c Repr) {
	genes := c.Genes().([]float64)
	var k float64
	for i, w := range genes {
		k += w * e.inputs[i]
	}
	k = 1.0 / ((Abs(k - e.target)) + 1e-6)
	c.Update(k)
}
