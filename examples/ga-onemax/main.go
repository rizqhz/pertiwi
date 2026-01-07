package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/rizqhz/pertiwi/genetic"
	"github.com/rizqhz/pertiwi/genetic/chromosome"
	"github.com/rizqhz/pertiwi/genetic/crossover"
	"github.com/rizqhz/pertiwi/genetic/evaluator"
	"github.com/rizqhz/pertiwi/genetic/initializer"
	"github.com/rizqhz/pertiwi/genetic/mutator"
	"github.com/rizqhz/pertiwi/genetic/selector"
	"github.com/rizqhz/pertiwi/genetic/strategy"
)

var (
	MAX_GENERATION  = 600
	POPULATION_SIZE = 100
	GENES_LENGTH    = 2048
	CROSSOVER_RATE  = 0.9
	MUTATION_RATE   = 0.001
	KEEP_ELITISM    = 4
	RANDOM_SEED     = rand.Uint64()
)

var (
	INITIALIZER  = initializer.Bitset()
	EVALUATOR    = evaluator.OneMax()
	SELECTOR     = selector.Tournament(3)
	RECOMBINATOR = crossover.Uniform()
	MUTATOR      = mutator.FlipBit()
	STRATEGY     = strategy.Stagnant(100, MUTATION_RATE)
)

var parameter = genetic.NewParameter(
	genetic.WithMaxGeneration(MAX_GENERATION),
	genetic.WithPopulationSize(POPULATION_SIZE),
	genetic.WithGenesLength(GENES_LENGTH),
	genetic.WithCrossoverRate(CROSSOVER_RATE),
	genetic.WithMutationRate(MUTATION_RATE),
	genetic.WithElitism(KEEP_ELITISM),
	genetic.WithRandomSeed(RANDOM_SEED),
)

var component = genetic.NewComponent(
	genetic.WithInitializer(INITIALIZER),
	genetic.WithEvaluator(EVALUATOR),
	genetic.WithSelector(SELECTOR),
	genetic.WithRecombinator(RECOMBINATOR),
	genetic.WithMutator(MUTATOR),
	genetic.WithStrategy(STRATEGY),
)

var (
	process = make(chan chromosome.Repr, POPULATION_SIZE)
	output  = make(chan chromosome.Repr)
)

func main() {
	engine := genetic.NewEngine(parameter, component)

	go func() {
		output <- engine.Evolve(process)
		close(output)
	}()

	for {
		select {
		case result := <-process:
			_ = result
		case result := <-output:
			fmt.Printf("C: %s F: %f\n", result, result.Score())
			return
		}
	}
}
