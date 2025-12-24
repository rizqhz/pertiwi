package random

import (
	"iter"
	"math/rand/v2"
)

type matrix struct {
	r   *rand.Rand
	row int
	col int
}

func Matrix(r *rand.Rand, row, col int) *matrix {
	return &matrix{r, row, col}
}

func (m *matrix) Next() [][]int {
	M := make([][]int, m.row)
	for r := range M {
		M[r] = make([]int, m.col)
		for c := range M[r] {
			M[r][c] = m.r.Int()
		}
	}
	return M
}

func (m *matrix) Fill(dst [][]int) {
	for r := range dst {
		for c := range dst[r] {
			dst[r][c] = m.r.Int()
		}
	}
}

func (m *matrix) Iter() iter.Seq[[][]int] {
	return func(yield func([][]int) bool) {
		for {
			if !yield(m.Next()) {
				return
			}
		}
	}
}

func (m *matrix) IterN(n int) iter.Seq[[][]int] {
	return func(yield func([][]int) bool) {
		for range n {
			if !yield(m.Next()) {
				return
			}
		}
	}
}
