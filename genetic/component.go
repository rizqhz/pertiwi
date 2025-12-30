package genetic

import (
	"github.com/rizqhz/pertiwi/genetic/crossover"
	"github.com/rizqhz/pertiwi/genetic/evaluator"
	"github.com/rizqhz/pertiwi/genetic/initializer"
	"github.com/rizqhz/pertiwi/genetic/mutator"
	"github.com/rizqhz/pertiwi/genetic/selector"
	"github.com/rizqhz/pertiwi/genetic/strategy"
)

type Component struct {
	initializer  initializer.Initializer
	evaluator    evaluator.Evaluator
	selector     selector.Selector
	recombinator crossover.Recombinator
	mutator      mutator.Mutator
	strategy     strategy.Strategy
}

type ComponentOption func(c *Component)

func NewComponent(opts ...ComponentOption) *Component {
	var component = &Component{
		selector:     selector.Tournament(3),
		recombinator: crossover.Uniform(),
		strategy:     strategy.Cosine(1e-3, 1e-1, 100),
	}
	for _, opt := range opts {
		opt(component)
	}
	if component.initializer == nil {
		panic("initializer must be set.")
	}
	if component.evaluator == nil {
		panic("evaluator must be set.")
	}
	if component.selector == nil {
		panic("selector must be set.")
	}
	if component.recombinator == nil {
		panic("recombinator must be set.")
	}
	if component.mutator == nil {
		panic("mutator must be set.")
	}
	return component
}

func WithInitializer(i initializer.Initializer) ComponentOption {
	return func(c *Component) {
		c.initializer = i
	}
}

func WithEvaluator(e evaluator.Evaluator) ComponentOption {
	return func(c *Component) {
		c.evaluator = e
	}
}

func WithSelector(s selector.Selector) ComponentOption {
	return func(c *Component) {
		c.selector = s
	}
}

func WithRecombinator(r crossover.Recombinator) ComponentOption {
	return func(c *Component) {
		c.recombinator = r
	}
}

func WithMutator(m mutator.Mutator) ComponentOption {
	return func(c *Component) {
		c.mutator = m
	}
}

func WithStrategy(s strategy.Strategy) ComponentOption {
	return func(c *Component) {
		c.strategy = s
	}
}
