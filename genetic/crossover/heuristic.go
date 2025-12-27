package crossover

import (
	. "math/rand/v2"

	. "github.com/rizqhz/pertiwi/genetic/chromosome"
)

type heuristic struct {
	rng *Rand
}

func Heuristic() *heuristic {
	return &heuristic{}
}

func (c *heuristic) Combine(p1, p2 Repr, rate float64) (c1, c2 Repr) {
	c1 = p1.Clone()
	c2 = p2.Clone()
	if n := p1.Len(); c.rng.Float64() < rate {
		pBest, pWorst := p1, p2
		if pWorst.Score() > pBest.Score() {
			pBest, pWorst = p2, p1
		}
		for i := range n {
			vBest := pBest.Get(i).(float64)
			vWorst := pWorst.Get(i).(float64)
			c1.Set(i, vBest+c.rng.Float64()*(vBest-vWorst))
			c2.Set(i, vBest+c.rng.Float64()*(vBest-vWorst))
		}
	}
	return c1, c2
}

func (c *heuristic) Random(r *Rand) {
	c.rng = r
}
