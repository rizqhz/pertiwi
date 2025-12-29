package strategy

import "math"

// Cosine Annealing
// y: min_rate + 0.5 * (max_rate - min_rate) * (1 + cos(PI * f(curr_gen)))
//
// f(curr_gen)
// | (curr_gen % period) / period

type cosine struct {
	rate   float64
	min    float64
	max    float64
	period int
}

func Cosine(min, max float64, period int) *cosine {
	return &cosine{
		rate:   max,
		min:    min,
		max:    max,
		period: period,
	}
}

func (s *cosine) Rate() float64 {
	return s.rate
}

func (s *cosine) Adapt(m Metrics) {
	k := float64(m.CurGen%s.period) / float64(s.period)
	s.rate = s.min + 0.5*(s.max-s.min)*(1+math.Cos(math.Pi*k))
}
