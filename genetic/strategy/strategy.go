package strategy

type Metrics struct {
	CurGen    int
	MaxGen    int
	BestScore float64
	Rate      float64
}

type Strategy interface {
	Rate() float64
	Adapt(Metrics)
}
