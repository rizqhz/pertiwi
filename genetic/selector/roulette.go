package selector

type roulette struct{}

func Roulette() *roulette {
	return &roulette{}
}

func (s *roulette) Select(p Set, r *Rand) Repr {
	var sum float64
	for _, c := range p {
		sum += c.Score()
	}

	pick := r.Float64() * sum

	var curr float64
	for _, c := range p {
		curr += c.Score()
		if curr >= pick {
			return c
		}
	}

	return p[len(p)-1]
}
