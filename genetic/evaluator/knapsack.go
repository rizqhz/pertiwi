package evaluator

import (
	. "github.com/rizqhz/pertiwi/genetic/chromosome"
)

type knapsack struct {
	weights   []int
	values    []int
	threshold int
}

func Knapsack(weights, values []int, threshold int) *knapsack {
	return &knapsack{weights, values, threshold}
}

func (e *knapsack) Evaluate(c Repr) {
	genes := c.Genes().([]int)
	var w, v int
	for i, k := range genes {
		if k == 1 {
			w += e.weights[i]
			v += e.values[i]
		}
	}
	var k float64
	if w <= e.threshold {
		k = float64(v)
	} else {
		k = 0
	}
	c.Update(k)
}
