package selector

type tournament struct {
	k int
}

func Tournament(k int) *tournament {
	return &tournament{k}
}

func (s *tournament) Select(p Set, r *Rand) Repr {
	var b Repr
	for range s.k {
		c := p[r.IntN(len(p))]
		if b == nil || c.Score() > b.Score() {
			b = c
		}
	}
	return b
}
