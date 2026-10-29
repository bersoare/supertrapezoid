// SPDX-License-Identifier: MIT
// Copyright Bernardo Soares

//go:build cgo && !purego && !assembly

package supertrapezoid

import (
	"testing"

	"supertrapezoid/pkg/trapezoid"
	"supertrapezoid/pkg/trapezoid/cgo"

	"github.com/stretchr/testify/require"
)

func TestTrapezoidCgo(t *testing.T) {
	t.Run("test simple version", func(t *testing.T) {
		result := cgo.TrapezoidCgo(trapezoid.Lower, trapezoid.Upper, trapezoid.Resolution)
		require.NotZero(t, result)
		require.Equal(t, trapezoid.TestResult, result)
	})

	t.Run("test 128SIMD version", func(t *testing.T) {
		result := cgo.TrapezoidSIMDAVX128Cgo(trapezoid.Lower, trapezoid.Upper, trapezoid.Resolution)
		require.NotZero(t, result)
		require.InDelta(t, trapezoid.TestResult, result, 1e-9)
	})

	t.Run("test 246SIMD version", func(t *testing.T) {
		result := cgo.TrapezoidSIMDAVX256Cgo(trapezoid.Lower, trapezoid.Upper, trapezoid.Resolution)
		require.NotZero(t, result)
		require.InDelta(t, trapezoid.TestResult, result, 1e-9)
	})
}

func BenchmarkEmptyFunctionCallCgo(b *testing.B) {
	var err error

	for range b.N {
		err = cgo.DoNothingCgo()
		require.NoError(b, err)
	}
}

func BenchmarkTrapezoidCgo(b *testing.B) {
	b.Run("1000", func(b *testing.B) {
		var result float64
		for range b.N {
			result = cgo.TrapezoidCgo(trapezoid.Lower, trapezoid.Upper, 1000)
		}

		require.NotZero(b, result)
	})

	b.Run("1000000", func(b *testing.B) {
		var result float64
		for range b.N {
			result = cgo.TrapezoidCgo(trapezoid.Lower, trapezoid.Upper, 1000000)
		}

		require.NotZero(b, result)
	})

	b.Run("500000000", func(b *testing.B) {
		var result float64
		for range b.N {
			result = cgo.TrapezoidCgo(trapezoid.Lower, trapezoid.Upper, 500000000)
		}

		require.NotZero(b, result)
	})
}

func BenchmarkTrapezoidSIMD128Cgo(b *testing.B) {
	b.Run("1000", func(b *testing.B) {
		var result float64
		for range b.N {
			result = cgo.TrapezoidSIMDAVX128Cgo(trapezoid.Lower, trapezoid.Upper, 1000)
		}

		require.NotZero(b, result)
	})

	b.Run("1000000", func(b *testing.B) {
		var result float64
		for range b.N {
			result = cgo.TrapezoidSIMDAVX128Cgo(trapezoid.Lower, trapezoid.Upper, 1000000)
		}

		require.NotZero(b, result)
	})

	b.Run("500000000", func(b *testing.B) {
		var result float64
		for range b.N {
			result = cgo.TrapezoidSIMDAVX128Cgo(trapezoid.Lower, trapezoid.Upper, 500000000)
		}

		require.NotZero(b, result)
	})
}

func BenchmarkTrapezoidSIMD256Cgo(b *testing.B) {
	b.Run("1000", func(b *testing.B) {
		var result float64
		for range b.N {
			result = cgo.TrapezoidSIMDAVX256Cgo(trapezoid.Lower, trapezoid.Upper, 1000)
		}

		require.NotZero(b, result)
	})

	b.Run("1000000", func(b *testing.B) {
		var result float64
		for range b.N {
			result = cgo.TrapezoidSIMDAVX256Cgo(trapezoid.Lower, trapezoid.Upper, 1000000)
		}

		require.NotZero(b, result)
	})

	b.Run("500000000", func(b *testing.B) {
		var result float64
		for range b.N {
			result = cgo.TrapezoidSIMDAVX256Cgo(trapezoid.Lower, trapezoid.Upper, 500000000)
		}

		require.NotZero(b, result)
	})
}
