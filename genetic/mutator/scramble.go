package mutator

type scramble struct{}

func Scramble() *scramble {
	return &scramble{}
}

func (m *scramble) Mutate(c Repr, p Rate, r *Rand) {
	if r.Float64() > p {
		return
	}

	n := c.Len()

	a := r.IntN(n)
	b := r.IntN(n)

	if a > b {
		a, b = b, a
	}

	if a == b {
		return
	}

	interval := b - a + 1
	idx := make([]int, interval)
	val := make([]any, interval)

	for i := range interval {
		idx[i] = a + i
		val[i] = c.Get(a + i)
	}

	r.Shuffle(interval, func(i, j int) {
		val[i], val[j] = val[j], val[i]
	})

	for i := range interval {
		c.Set(idx[i], val[i])
	}
}
