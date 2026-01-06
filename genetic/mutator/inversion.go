package mutator

type inversion struct{}

func Inversion() *inversion {
	return &inversion{}
}

func (m *inversion) Mutate(c Repr, p Rate, r *Rand) {
	if r.Float64() > p {
		return
	}

	n := c.Len()

	a := r.IntN(n)
	b := r.IntN(n)

	if a > b {
		a, b = b, a
	}

	for a < b {
		x := c.Get(a)
		y := c.Get(b)
		c.Set(a, y)
		c.Set(b, x)
		a += 1
		b -= 1
	}
}
