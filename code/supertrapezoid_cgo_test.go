// SPDX-License-Identifier: MIT
// Copyright Bernardo Soares

//go:build cgo && !purego

package supertrapezoid

import (
	"fmt"
	"testing"

	"supertrapezoid/pkg/trapezoid"
	"supertrapezoid/pkg/trapezoid/cgo"

	"github.com/stretchr/testify/require"
)

func TestTrapezoidCgo(t *testing.T) {
	t.Run("scalar", func(t *testing.T) {
		result := cgo.TrapezoidCgo(trapezoid.Lower, trapezoid.Upper, trapezoid.Resolution)
		require.NotZero(t, result)
		require.Equal(t, trapezoid.TestResult, result)
	})

	t.Run("SIMD-128", func(t *testing.T) {
		result := cgo.TrapezoidSIMDAVX128Cgo(trapezoid.Lower, trapezoid.Upper, trapezoid.Resolution)
		require.NotZero(t, result)
		require.InDelta(t, trapezoid.TestResult, result, 1e-9)
	})

	t.Run("SIMD-256", func(t *testing.T) {
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

func BenchmarkCgo(b *testing.B) {
	b.Run("scalar", func(b *testing.B) {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("%d", tc), func(b *testing.B) {
				var result float64
				for range b.N {
					result = cgo.TrapezoidCgo(trapezoid.Lower, trapezoid.Upper, tc)
				}

				require.NotZero(b, result)
			})
		}
	})

	b.Run("SIMD-128", func(b *testing.B) {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("%d", tc), func(b *testing.B) {
				var result float64
				for range b.N {
					result = cgo.TrapezoidSIMDAVX128Cgo(trapezoid.Lower, trapezoid.Upper, tc)
				}

				require.NotZero(b, result)
			})
		}
	})

	b.Run("SIMD-256", func(b *testing.B) {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("%d", tc), func(b *testing.B) {
				var result float64
				for range b.N {
					result = cgo.TrapezoidSIMDAVX256Cgo(trapezoid.Lower, trapezoid.Upper, tc)
				}

				require.NotZero(b, result)
			})
		}
	})
}
