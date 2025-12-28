package evaluator

import (
	. "github.com/rizqhz/pertiwi/genetic/chromosome"
)

type onemax struct{}

func OneMax() *onemax {
	return &onemax{}
}

func (e *onemax) Evaluate(c Repr) {
	genes := c.Genes().([]int)
	var k float64
	for _, v := range genes {
		if v == 1 {
			k++
		}
	}
	c.Update(k)
}
