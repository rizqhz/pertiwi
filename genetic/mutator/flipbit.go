package mutator

type flipbit struct{}

func FlipBit() *flipbit {
	return &flipbit{}
}

func (m *flipbit) Mutate(c Repr, p Rate, r *Rand) {
	genes := c.Genes().([]int)
	for i := range genes {
		if r.Float64() < p {
			genes[i] ^= 1
		}
	}
}
