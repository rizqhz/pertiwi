package crossover

type blend struct {
	a float64
}

func Blend(alpha float64) *blend {
	return &blend{alpha}
}

func (c *blend) Combine(p1, p2 Repr, p Rate, r *Rand) (Repr, Repr) {
	if r.Float64() > p {
		return p1.Clone(), p2.Clone()
	}

	c1, c2 := p1.Clone(), p2.Clone()
	n := p1.Len()

	for i := range n {
		x := p1.Get(i).(float64)
		y := p2.Get(i).(float64)

		if x > y {
			x, y = y, x
		}

		lo := x - (c.a * (y - x))
		hi := y + (c.a * (y - x))

		c1.Set(i, (lo + r.Float64()*(hi-lo)))
		c2.Set(i, (lo + r.Float64()*(hi-lo)))
	}

	return c1, c2
}
