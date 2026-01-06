package crossover

type cycle struct{}

func Cycle() *cycle {
	return &cycle{}
}

func (c *cycle) Combine(p1, p2 Repr, p Rate, r *Rand) (Repr, Repr) {
	if r.Float64() > p {
		return p1.Clone(), p2.Clone()
	}

	c1, c2 := p1.Clone(), p2.Clone()
	n := p1.Len()

	pos := make(map[any]int, n)
	for i := range n {
		pos[p1.Get(i)] = i
	}

	cycle := make([]bool, n)
	for i := 0; !cycle[i]; i = pos[p2.Get(i)] {
		cycle[i] = true
	}

	for i := range n {
		if !cycle[i] {
			c1.Set(i, p2.Get(i))
			c2.Set(i, p1.Get(i))
		}
	}

	return c1, c2
}
