package genetic

import (
	"runtime"
	"sync"

	. "github.com/rizqhz/pertiwi/genetic/chromosome"
	"github.com/rizqhz/pertiwi/random"
)

type Engine struct {
	*Parameter
	*Component
}

func NewEngine(p *Parameter, c *Component) *Engine {
	var engine = &Engine{
		Parameter: p,
		Component: c,
	}
	RNG := random.From(engine.RandomSeed)
	engine.Setup(RNG)
	return engine
}

func (e *Engine) Setup(r random.Source) {
	e.initializer.Size(e.PopulationSize)
	e.initializer.Length(e.GenesLength)
	e.initializer.Random(r.PRNG())
	e.selector.Random(r.PRNG())
	e.recombinator.Random(r.PRNG())
	e.mutator.Random(r.PRNG())
}

func (e *Engine) Compute(p Set) {
	var wg sync.WaitGroup
	jobs := make(chan Repr, len(p))
	worker := func() {
		defer wg.Done()
		for c := range jobs {
			e.evaluator.Evaluate(c)
		}
	}
	n := runtime.NumCPU()
	wg.Add(n)
	for range n {
		go worker()
	}
	for _, c := range p {
		jobs <- c
	}
	close(jobs)
	wg.Wait()
}
