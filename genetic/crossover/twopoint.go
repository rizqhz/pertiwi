package crossover

type twopoint struct{}

func TwoPoint() *twopoint {
	return &twopoint{}
}

func (c *twopoint) Combine(p1, p2 Repr, p Rate, r *Rand) (Repr, Repr) {
	if r.Float64() > p {
		return p1.Clone(), p2.Clone()
	}

	c1, c2 := p1.Clone(), p2.Clone()
	n := p1.Len()

	a := (r.IntN(n-0) + 0) % n
	b := (r.IntN(n-1) + a + 1) % n
	if a > b {
		a, b = b, a
	}

	for i := a; i < b; i++ {
		c1.Set(i, p2.Get(i))
		c2.Set(i, p1.Get(i))
	}

	return c1, c2
}
