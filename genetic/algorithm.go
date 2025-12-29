package genetic

import (
	. "github.com/rizqhz/pertiwi/genetic/chromosome"
	"github.com/rizqhz/pertiwi/genetic/strategy"
)

func (e *Engine) Evolve(proc chan Repr) Repr {
	curr := e.initializer.Populate()
	next := e.initializer.Empty()
	e.Compute(curr)
	best := curr.Best(e.KeepElitism)
	for iter := 0; iter < e.MaxGeneration; iter++ {
		for i := 0; i < e.KeepElitism; i++ {
			next[i] = best[i].Clone()
		}
		for i := e.KeepElitism; i < e.PopulationSize; i += 2 {
			p1 := e.selector.Select(curr)
			p2 := e.selector.Select(curr)
			c1, c2 := e.recombinator.Combine(p1, p2, e.CrossoverRate)
			e.mutator.Mutate(c1, e.strategy.Rate())
			e.mutator.Mutate(c2, e.strategy.Rate())
			next[i] = c1
			if k := i + 1; k < e.PopulationSize {
				next[k] = c2
			}
		}
		curr, next = next, curr
		e.Compute(curr)
		best = curr.Best(e.KeepElitism)
		e.strategy.Adapt(strategy.Metrics{
			CurGen:    iter,
			MaxGen:    e.MaxGeneration,
			BestScore: best[0].Score(),
			Rate:      e.MutationRate,
		})
		// proc <- best[0]
	}
	// close(proc)
	return best[0]
}
