package mutator

type swap struct{}

func Swap() *swap {
	return &swap{}
}

func (m *swap) Mutate(c Repr, p float64, r *Rand) {
	if r.Float64() > p {
		return
	}

	n := c.Len()

	i := r.IntN(n)
	j := r.IntN(n)

	x, y := c.Get(i), c.Get(j)

	c.Set(i, y)
	c.Set(j, x)
}
