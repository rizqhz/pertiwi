package crossover

type uniform struct{}

func Uniform() *uniform {
	return &uniform{}
}

func (c *uniform) Combine(p1, p2 Repr, p Rate, r *Rand) (Repr, Repr) {
	if r.Float64() > p {
		return p1.Clone(), p2.Clone()
	}

	c1, c2 := p1.Clone(), p2.Clone()
	n := p1.Len()

	for i := range n {
		if r.Float64() < 0.5 {
			c1.Set(i, p2.Get(i))
			c2.Set(i, p1.Get(i))
		}
	}

	return c1, c2
}
