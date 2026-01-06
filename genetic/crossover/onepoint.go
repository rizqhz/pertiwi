package crossover

type onepoint struct{}

func OnePoint() *onepoint {
	return &onepoint{}
}

func (c *onepoint) Combine(p1, p2 Repr, p Rate, r *Rand) (Repr, Repr) {
	if r.Float64() > p {
		return p1.Clone(), p2.Clone()
	}

	c1, c2 := p1.Clone(), p2.Clone()
	n := p1.Len()

	k := r.IntN(n-1) + 1

	for i := k; i < n; i++ {
		c1.Set(i, p2.Get(i))
		c2.Set(i, p1.Get(i))
	}

	return c1, c2
}
