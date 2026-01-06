package crossover

type arithmetic struct{}

func Arithmetic() *arithmetic {
	return &arithmetic{}
}

func (c *arithmetic) Combine(p1, p2 Repr, p Rate, r *Rand) (Repr, Repr) {
	if r.Float64() > p {
		return p1.Clone(), p2.Clone()
	}

	c1, c2 := p1.Clone(), p2.Clone()
	n := p1.Len()

	alpha := r.Float64()
	compl := 1 - alpha

	for i := range n {
		x := p1.Get(i).(float64)
		y := p2.Get(i).(float64)

		c1.Set(i, (alpha*x + compl*y))
		c2.Set(i, (alpha*y + compl*x))
	}

	return c1, c2
}
