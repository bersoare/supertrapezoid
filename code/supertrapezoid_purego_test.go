// SPDX-License-Identifier: MIT
// Copyright Bernardo Soares

//go:build purego && !cgo && !assembly

package supertrapezoid

import (
	"testing"

	"supertrapezoid/pkg/trapezoid"
	"supertrapezoid/pkg/trapezoid/purego"
	"supertrapezoid/pkg/trapezoid/utils"

	"github.com/stretchr/testify/require"
)

func TestTrapezoidPurego(t *testing.T) {
	t.Run("test simple version", func(t *testing.T) {
		result := purego.TrapezoidPureGo(trapezoid.Lower, trapezoid.Upper, trapezoid.Resolution)
		require.NotZero(t, result)
		require.InDelta(t, trapezoid.TestResult, result, 1e-9)
	})

	utils.PrintAVXSupport()
	t.Run("test SIMD version", func(t *testing.T) {
		result := purego.TrapezoidSIMDPureGo(trapezoid.Lower, trapezoid.Upper, trapezoid.Resolution)
		require.NotZero(t, result)
		require.InDelta(t, trapezoid.TestResult, result, 1e-9)
	})

	t.Run("test ArchSIMD 128 version", func(t *testing.T) {
		result := purego.TrapezoidArchSIMD128PureGo(trapezoid.Lower, trapezoid.Upper, trapezoid.Resolution)
		require.NotZero(t, result)
		require.InDelta(t, trapezoid.TestResult, result, 1e-9)
	})

	t.Run("test ArchSIMD 256 version", func(t *testing.T) {
		result := purego.TrapezoidArchSIMD256PureGo(trapezoid.Lower, trapezoid.Upper, trapezoid.Resolution)
		require.NotZero(t, result)
		require.InDelta(t, trapezoid.TestResult, result, 1e-9)
	})
}

func BenchmarkEmptyFunctionCallPureGo(b *testing.B) {
	var err error

	for range b.N {
		err = purego.DoNothingPureGo()
		require.NoError(b, err)
	}
}

func BenchmarkTrapezoidPureGo(b *testing.B) {
	b.Run("1000", func(b *testing.B) {
		var result float64
		for range b.N {
			result = purego.TrapezoidPureGo(trapezoid.Lower, trapezoid.Upper, 1000)
		}

		require.NotZero(b, result)
	})

	b.Run("1000000", func(b *testing.B) {
		var result float64
		for range b.N {
			result = purego.TrapezoidPureGo(trapezoid.Lower, trapezoid.Upper, 1000000)
		}

		require.NotZero(b, result)
	})

	b.Run("500000000", func(b *testing.B) {
		var result float64
		for range b.N {
			result = purego.TrapezoidPureGo(trapezoid.Lower, trapezoid.Upper, 500000000)
		}

		require.NotZero(b, result)
	})
}

func BenchmarkTrapezoidSIMDPureGo(b *testing.B) {
	b.Run("1000", func(b *testing.B) {
		var result float64
		for range b.N {
			result = purego.TrapezoidSIMDPureGo(trapezoid.Lower, trapezoid.Upper, 1000)
		}

		require.NotZero(b, result)
	})

	b.Run("1000000", func(b *testing.B) {
		var result float64
		for range b.N {
			result = purego.TrapezoidSIMDPureGo(trapezoid.Lower, trapezoid.Upper, 1000000)
		}

		require.NotZero(b, result)
	})

	b.Run("500000000", func(b *testing.B) {
		var result float64
		for range b.N {
			result = purego.TrapezoidSIMDPureGo(trapezoid.Lower, trapezoid.Upper, 500000000)
		}

		require.NotZero(b, result)
	})
}

func BenchmarkTrapezoidArchSIMD256PureGo(b *testing.B) {
	b.Run("1000", func(b *testing.B) {
		var result float64
		for range b.N {
			result = purego.TrapezoidArchSIMD256PureGo(trapezoid.Lower, trapezoid.Upper, 1000)
		}

		require.NotZero(b, result)
	})

	b.Run("1000000", func(b *testing.B) {
		var result float64
		for range b.N {
			result = purego.TrapezoidArchSIMD256PureGo(trapezoid.Lower, trapezoid.Upper, 1000000)
		}

		require.NotZero(b, result)
	})

	b.Run("500000000", func(b *testing.B) {
		var result float64
		for range b.N {
			result = purego.TrapezoidArchSIMD256PureGo(trapezoid.Lower, trapezoid.Upper, 500000000)
		}

		require.NotZero(b, result)
	})
}

func BenchmarkTrapezoidArchSIMD128PureGo(b *testing.B) {
	b.Run("1000", func(b *testing.B) {
		var result float64
		for range b.N {
			result = purego.TrapezoidArchSIMD128PureGo(trapezoid.Lower, trapezoid.Upper, 1000)
		}

		require.NotZero(b, result)
	})

	b.Run("1000000", func(b *testing.B) {
		var result float64
		for range b.N {
			result = purego.TrapezoidArchSIMD128PureGo(trapezoid.Lower, trapezoid.Upper, 1000000)
		}

		require.NotZero(b, result)
	})

	b.Run("500000000", func(b *testing.B) {
		var result float64
		for range b.N {
			result = purego.TrapezoidArchSIMD128PureGo(trapezoid.Lower, trapezoid.Upper, 500000000)
		}

		require.NotZero(b, result)
	})
}
