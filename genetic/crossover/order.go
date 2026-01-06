package crossover

type order struct{}

func Order() *order {
	return &order{}
}

func (c *order) Combine(p1, p2 Repr, p Rate, r *Rand) (Repr, Repr) {
	if r.Float64() > p {
		return p1.Clone(), p2.Clone()
	}

	c1, c2 := p1.Clone(), p2.Clone()
	n := p1.Len()

	a, b := r.IntN(n), r.IntN(n)
	if a > b {
		a, b = b, a
	}

	fill := func(child, parent, donor Repr) {
		used := make(map[any]bool)
		for i := a; i <= b; i++ {
			val := parent.Get(i)
			child.Set(i, val)
			used[val] = true
		}

		pos := (b + 1) % n
		for i := range n {
			k := (b + 1 + i) % n
			gene := donor.Get(k)
			if !used[gene] {
				child.Set(pos, gene)
				pos = (pos + 1) % n
			}
		}
	}

	fill(c1, p1, p2)
	fill(c2, p2, p1)

	return c1, c2
}
