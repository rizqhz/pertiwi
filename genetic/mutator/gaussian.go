package mutator

type gaussian struct {
	sigma float64
}

func Gaussian(stddev float64) *gaussian {
	return &gaussian{stddev}
}

func (m *gaussian) Mutate(c Repr, p Rate, r *Rand) {
	genes := c.Genes().([]float64)
	for i := range genes {
		if r.Float64() < p {
			genes[i] += m.sigma * r.NormFloat64()
		}
	}
}
