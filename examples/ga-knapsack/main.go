package main

import (
	"fmt"
	"math/rand/v2"
	"time"

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
	weights = []int{
		7, 0, 30, 22, 80, 94, 11, 81, 70, 64, 59, 18, 0, 36, 3, 8, 15, 42, 9, 0, 42, 47,
		52, 32, 26, 48, 55, 6, 29, 84, 2, 4, 18, 56, 7, 29, 93, 44, 71, 3, 86, 66, 31,
		65, 0, 79, 20, 65, 52, 13,
	}
	values = []int{
		360, 83, 59, 130, 431, 67, 230, 52, 93, 125, 670, 892, 600, 38, 48, 147, 78, 256,
		63, 17, 120, 164, 432, 35, 92, 110, 22, 42, 50, 323, 514, 28, 87, 73, 78, 15, 26,
		78, 210, 36, 85, 189, 274, 43, 33, 10, 19, 389, 276, 312,
	}
	threshold = 850
)

var (
	MAX_GENERATION  = 1000
	POPULATION_SIZE = 200
	GENES_LENGTH    = 50
	CROSSOVER_RATE  = 0.9
	MUTATION_RATE   = 0.01
	KEEP_ELITISM    = 3
	RANDOM_SEED     = rand.Uint64()
)

var (
	INITIALIZER  = initializer.Bitset()
	EVALUATOR    = evaluator.Knapsack(weights, values, threshold)
	SELECTOR     = selector.Tournament(3)
	RECOMBINATOR = crossover.Uniform()
	MUTATOR      = mutator.FlipBit()
	STRATEGY     = strategy.Cosine(MUTATION_RATE, MUTATION_RATE*35, 100)
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

func Knapsack(c chromosome.Repr) {
	genes := c.Genes().([]int)
	var w, v int
	for i, k := range genes {
		if k == 1 {
			w += weights[i]
			v += values[i]
		}
	}
	fmt.Printf("w: %v\n", w)
	fmt.Printf("v: %v\n", v)
}

func main() {
	engine := genetic.NewEngine(parameter, component)
	go func() {
		output <- engine.Evolve(process)
		close(output)
	}()
	t := time.Now()
	for {
		select {
		case result, ok := <-process:
			if ok {
				fmt.Printf("%s (%f)\n", result, result.Score())
				time.Sleep(time.Millisecond * 50)
			}
		case result, ok := <-output:
			if ok {
				fmt.Printf("%s (%f)\n", result, result.Score())
			} else {
				ms := time.Since(t).Milliseconds()
				fmt.Printf("%vms\n", ms)
				return
			}
		}
	}
}
