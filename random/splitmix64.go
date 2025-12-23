package random

import "sync/atomic"

// goldenRatio is the 64-bit constant increment used by the SplitMix64
// sequence generator. It corresponds to the fractional part of the
// golden ratio scaled by 2^64.
const goldenRatio uint64 = 0x9e3779b97f4a7c15

// SplitMix64 applies the standard SplitMix64 mixing function. It takes
// a 64-bit input and performs a series of reversible bitwise and
// multiplicative transformations to produce a statistically uniform
// 64-bit output.
//
// See http://xorshift.di.unimi.it/splitmix64.c
func SplitMix64(state *atomic.Uint64) uint64 {
	x := state.Add(goldenRatio)
	x = (x ^ (x >> 30)) * 0xbf58476d1ce4e5b9
	x = (x ^ (x >> 27)) * 0x94d049bb133111eb
	return x ^ (x >> 31)
}
