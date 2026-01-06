package crossover

type shuffle struct{}

func Shuffle() *shuffle {
	return &shuffle{}
}

func (c *shuffle) Combine(p1, p2 Repr, p Rate, r *Rand) (Repr, Repr) {
	if r.Float64() > p {
		return p1.Clone(), p2.Clone()
	}

	c1, c2 := p1.Clone(), p2.Clone()
	n := p1.Len()

	point := r.IntN(n-1) + 1
	index := r.Perm(n)

	for k := point; k < n; k++ {
		i := index[k]
		c1.Set(i, p2.Get(i))
		c2.Set(i, p1.Get(i))
	}

	return c1, c2
}
