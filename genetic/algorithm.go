package genetic

import (
	"runtime"
	"sync"

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
		e.Repopulate(curr, next)
		curr, next = next, curr
		e.Compute(curr)
		best = curr.Best(e.KeepElitism)
		e.strategy.Adapt(strategy.Metrics{
			CurGen:    iter,
			MaxGen:    e.MaxGeneration,
			BestScore: best[0].Score(),
			Rate:      e.MutationRate,
		})
		proc <- best[0].Clone()
	}
	close(proc)
	return best[0]
}

func (e *Engine) Repopulate(curr, next Set) {
	var wg sync.WaitGroup
	jobs := make(chan int, len(curr)/2-e.KeepElitism)
	worker := func() {
		defer wg.Done()
		for i := range jobs {
			p1 := e.selector.Select(curr)
			p2 := e.selector.Select(curr)
			c1, c2 := e.recombinator.Combine(p1, p2, e.CrossoverRate)
			rate := e.strategy.Rate()
			e.mutator.Mutate(c1, rate)
			e.mutator.Mutate(c2, rate)
			next[i] = c1
			if k := i + 1; k < e.PopulationSize {
				next[k] = c2
			}
		}
	}
	n := runtime.NumCPU()
	wg.Add(n)
	for range n {
		go worker()
	}
	for i := e.KeepElitism; i < e.PopulationSize; i += 2 {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}
