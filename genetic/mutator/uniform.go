package mutator

type uniform struct {
	a float64
	b float64
}

func Uniform(min, max float64) *uniform {
	return &uniform{
		a: min,
		b: max,
	}
}

func (m *uniform) Mutate(c Repr, p Rate, r *Rand) {
	genes := c.Genes().([]float64)
	for i := range genes {
		if r.Float64() < p {
			genes[i] = m.a + r.Float64()*(m.b-m.a)
		}
	}
}
