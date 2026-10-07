// SPDX-License-Identifier: MIT
// Copyright Bernardo Soares

package purego

import (
	"simd"
)

// When using platform agnostic SIMD (`simd`) instead of `archsimd`,
// even though we support AVX2, our vectors are capped at 128bit size
// (2 lanes for float64). This is a deliberate decision in the Go
// implementation: the priority is portability.
// On the other hand, we can access the intrinsics directly via
// `archsimd` package with the drawback of lack of portability (each arch
// needs different code)

//go:noinline
func DoNothingPureGo() error { return nil }

func integrableFunc(x float64) float64 {
	// x * (x - 1)
	// expanded to reduce data dependency
	return x*x - x
}

func TrapezoidPureGo(lower, upper float64, resolution int) float64 {
	// avoid conversions in the hot path
	res := float64(resolution)

	// interval width: horizontal width of every
	// slice in the area
	h := (upper - lower) / res

	// the heights of the curve at the edges dont share a boundary
	// with an adjacent slice, so we give it only half weight.
	// inner loop should skip both of these
	// thats why it starts at 1 and ends with i < res.
	sum := .5 * (integrableFunc(lower) + integrableFunc(upper))

	for i := 1.0; i < res; i++ {
		sum += integrableFunc(lower + (i * h))
	}

	return sum * h
}

// same as the one above
func integrableFuncSIMD(x simd.Float64s, minusOne simd.Float64s) simd.Float64s {
	return x.MulAdd(x, minusOne.Mul(x))
}

func TrapezoidSIMDPureGo(lower, upper float64, resolution int) float64 {
	res := float64(resolution)

	h := (upper - lower) / res

	sum := .5 * (integrableFunc(lower) + integrableFunc(upper))

	var v simd.Float64s
	width := v.Len()

	// prepare step index grids directly inside CPU registers
	steps := make([]float64, width)
	for i := 0; i < width; i++ {
		steps[i] = float64(i)
	}

	// step indexes, like a "lane id".
	// in scalar code, this is the id that walks
	// forward in the iterations. since we are dealing
	// with multiple lanes, we store the offset of each of them
	// (ex: 0,1,2,3 for a 4-lane SIMD)
	vecSteps := simd.LoadFloat64s(steps)

	// declare our vector variables
	// avoid broadcasting in the hot path for constants
	// since it may cause pipeline stalls (its slow)
	vH := simd.BroadcastFloat64s(h)
	vLower := simd.BroadcastFloat64s(lower)
	// how much we stepping at each loop
	// h * width
	vBlockStep := simd.BroadcastFloat64s(h * float64(width))
	// vX (x vector) starts at the first block coordinates
	// subtract the block step to allow for proper pipelining,
	// reducing data dependency that cause pipeline stalls
	vX := vLower.Add(vecSteps.Mul(vH)).Sub(vBlockStep)
	minusOne := simd.BroadcastFloat64s(-1)

	// accumulator
	vSum := simd.BroadcastFloat64s(0)

	// number of iterations, so we can
	// do a for loop over blocks as opposed to
	// comparing with a larger resolution number.
	blocks := resolution / width

	for _ = range blocks {
		vX = vX.Add(vBlockStep)
		vSum = vSum.Add(integrableFuncSIMD(vX, minusOne))
	}

	i := 1 + (blocks * width)

	for ; i < resolution; i++ {
		sum += integrableFunc(lower + (float64(i) * h))
	}

	vSum.Store(steps)
	for _, i := range steps {
		sum += i
	}

	return sum * h
}
