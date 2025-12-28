package evaluator

import (
	. "math"

	. "github.com/rizqhz/pertiwi/genetic/chromosome"
)

type euclidean struct {
	points [][2]float64
}

func Euclidean(coord ...[2]float64) *euclidean {
	return &euclidean{coord}
}

func (e *euclidean) Evaluate(c Repr) {
	distance := func(x1, y1 float64, x2, y2 float64) float64 {
		return Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
	}
	genes := c.Genes().([]int)
	var x1, x2, y1, y2 float64
	var d float64
	for i := range c.Len() - 1 {
		x1, y1 = e.points[genes[i+0]][0], e.points[genes[i+0]][1]
		x2, y2 = e.points[genes[i+1]][0], e.points[genes[i+1]][1]
		d += distance(x1, y1, x2, y2)
	}
	x1, y1 = e.points[genes[c.Len()-1]][0], e.points[genes[c.Len()-1]][1]
	x2, y2 = e.points[genes[0]][0], e.points[genes[0]][1]
	d += distance(x1, y1, x2, y2)
	c.Update(1.0 / (d + 1e-6))
}
