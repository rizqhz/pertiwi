package genetic

import (
	"runtime"
	"sync"

	"github.com/rizqhz/pertiwi/genetic/chromosome"
)

type Engine struct {
	*Parameter
	*Component
}

func NewEngine(p *Parameter, c *Component) *Engine {
	if p == nil || c == nil {
		panic("parameter or component must be set.")
	}
	return &Engine{p, c}
}

func (e *Engine) Compute(s chromosome.Set) {
	var wg sync.WaitGroup

	jobs := make(chan chromosome.Repr, len(s))
	worker := func() {
		for c := range jobs {
			e.evaluator.Evaluate(c)
		}
		wg.Done()
	}

	n := runtime.GOMAXPROCS(0)
	wg.Add(n)
	for range n {
		go worker()
	}

	for _, c := range s {
		jobs <- c
	}

	close(jobs)
	wg.Wait()
}
