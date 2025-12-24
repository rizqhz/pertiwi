package random

import (
	std "image/color"
	"iter"
	"math/rand/v2"
)

type color struct {
	r    *rand.Rand
	gray bool
}

func Color(r *rand.Rand, rgba bool) *color {
	return &color{r, !rgba}
}

func (c *color) Next() std.Color {
	k := c.r.Uint32()
	if c.gray {
		return std.Gray{uint8(k)}
	}
	return std.RGBA{
		R: uint8(k >> 24),
		G: uint8(k >> 16),
		B: uint8(k >> 8),
		A: 255,
	}
}

func (c *color) Fill(dst []std.Color) {
	for i := range dst {
		dst[i] = c.Next()
	}
}

func (c *color) Iter() iter.Seq[std.Color] {
	return func(yield func(std.Color) bool) {
		for {
			if !yield(c.Next()) {
				return
			}
		}
	}
}

func (c *color) IterN(n int) iter.Seq[std.Color] {
	return func(yield func(std.Color) bool) {
		for range n {
			if !yield(c.Next()) {
				return
			}
		}
	}
}
