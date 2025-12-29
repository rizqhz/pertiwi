package genetic

import (
	"encoding/json"
	"math/rand/v2"
)

type Parameter struct {
	MaxGeneration  int     `json:"max_generation"`
	PopulationSize int     `json:"population_size"`
	GenesLength    int     `json:"genes_length"`
	CrossoverRate  float64 `json:"crossover_rate"`
	MutationRate   float64 `json:"mutation_rate"`
	KeepElitism    int     `json:"keep_elitism"`
	RandomSeed     uint64  `json:"random_seed"`
}

func (p *Parameter) Encode() []byte {
	b, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	return b
}

func (p *Parameter) Decode(b []byte) {
	err := json.Unmarshal(b, p)
	if err != nil {
		panic(err)
	}
}

type ParameterOption func(p *Parameter)

func NewParameter(opts ...ParameterOption) *Parameter {
	var parameter = &Parameter{
		MaxGeneration:  500,
		PopulationSize: 100,
		GenesLength:    1,
		CrossoverRate:  0.900,
		MutationRate:   0.001,
		KeepElitism:    0,
		RandomSeed:     rand.Uint64(),
	}
	for _, opt := range opts {
		opt(parameter)
	}
	return parameter
}

func WithMaxGeneration(max int) ParameterOption {
	if max < 50 {
		panic("maximum number of generation must be greater than or equal to 50")
	}
	return func(p *Parameter) {
		p.MaxGeneration = max
	}
}

func WithPopulationSize(size int) ParameterOption {
	if size < 10 {
		panic("population size must be greater than or equal to 10")
	}
	return func(p *Parameter) {
		p.PopulationSize = size
	}
}

func WithGenesLength(length int) ParameterOption {
	if length < 1 {
		panic("length of chromosome genes must not be 0")
	}
	return func(p *Parameter) {
		p.GenesLength = length
	}
}

func WithCrossoverRate(rate float64) ParameterOption {
	if rate < 1e-6 {
		panic("crossover rate must be greater than 0")
	}
	return func(p *Parameter) {
		p.CrossoverRate = rate
	}
}

func WithMutationRate(rate float64) ParameterOption {
	if rate < 1e-6 {
		panic("mutation rate must be greater than 0")
	}
	return func(p *Parameter) {
		p.MutationRate = rate
	}
}

func WithElitism(n int) ParameterOption {
	if n < 1 {
		panic("number of elitist must not be 0")
	}
	return func(p *Parameter) {
		p.KeepElitism = n
	}
}

func WithRandomSeed(seed uint64) ParameterOption {
	return func(p *Parameter) {
		p.RandomSeed = seed
	}
}
