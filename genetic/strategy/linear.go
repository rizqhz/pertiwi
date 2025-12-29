package strategy

// Linear Decay
// y: rate_start - (curr_gen / max_gen) * (rate_start - rate_end)

type linear struct {
	rate  float64
	start float64
	end   float64
}

func Linear(start, end float64) *linear {
	return &linear{
		rate:  start,
		start: start,
		end:   end,
	}
}

func (s *linear) Rate() float64 {
	return s.rate
}

func (s *linear) Adapt(m Metrics) {
	progress := float64(m.CurGen) / float64(m.MaxGen)
	s.rate = s.start - progress*(s.start-s.end)
	s.rate = max(s.rate, s.end)
}
