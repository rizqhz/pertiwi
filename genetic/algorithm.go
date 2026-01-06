package genetic

import (
	"runtime"
	"sync"

	"github.com/rizqhz/pertiwi/genetic/chromosome"
	"github.com/rizqhz/pertiwi/genetic/strategy"
	"github.com/rizqhz/pertiwi/random"
)

func (e *Engine) Evolve(proc chan chromosome.Repr) chromosome.Repr {
	r := random.From(e.RandomSeed)

	curr := e.initializer.Populate(r.PRNG(), e.PopulationSize, e.GenesLength)
	next := e.initializer.Empty(r.PRNG(), e.PopulationSize)

	e.Compute(curr)
	best := curr.Best(e.KeepElitism)

	for iter := 0; iter < e.MaxGeneration; iter++ {
		for i := 0; i < e.KeepElitism; i++ {
			next[i] = best[i].Clone()
		}

		e.Repopulate(curr, next, r)
		curr, next = next, curr

		e.Compute(curr)
		best = curr.Best(e.KeepElitism)
		proc <- best[0].Clone()

		e.strategy.Adapt(strategy.Metrics{
			CurGen:    iter,
			MaxGen:    e.MaxGeneration,
			BestScore: best[0].Score(),
			Rate:      e.MutationRate,
		})
	}

	close(proc)
	return best[0]
}

func (e *Engine) Repopulate(curr, next chromosome.Set, r random.Source) {
	var wg sync.WaitGroup

	jobs := make(chan int, (e.PopulationSize-e.KeepElitism+1)/2)
	worker := func(seed uint64) {
		r := random.From(seed)
		for i := range jobs {
			p1 := e.selector.Select(curr, r.PRNG())
			p2 := e.selector.Select(curr, r.PRNG())

			c1, c2 := e.recombinator.Combine(p1, p2, e.CrossoverRate, r.PRNG())

			rate := e.strategy.Rate()
			e.mutator.Mutate(c1, rate, r.PRNG())
			e.mutator.Mutate(c2, rate, r.PRNG())

			next[i] = c1
			if k := i + 1; k < e.PopulationSize {
				next[k] = c2
			}
		}
		wg.Done()
	}

	n := runtime.GOMAXPROCS(0)
	wg.Add(n)
	for range n {
		go worker(r.Uint64())
	}

	for i := e.KeepElitism; i < e.PopulationSize; i += 2 {
		jobs <- i
	}

	close(jobs)
	wg.Wait()
}
