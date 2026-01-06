package crossover

type heuristic struct{}

func Heuristic() *heuristic {
	return &heuristic{}
}

func (c *heuristic) Combine(p1, p2 Repr, p Rate, r *Rand) (Repr, Repr) {
	if r.Float64() > p {
		return p1.Clone(), p2.Clone()
	}

	c1, c2 := p1.Clone(), p2.Clone()
	n := p1.Len()

	a, b := p1, p2
	if b.Score() > a.Score() {
		a, b = p2, p1
	}

	for i := range n {
		x := a.Get(i).(float64)
		y := b.Get(i).(float64)

		c1.Set(i, (x + r.Float64()*(x-y)))
		c2.Set(i, (x + r.Float64()*(x-y)))
	}

	return c1, c2
}
