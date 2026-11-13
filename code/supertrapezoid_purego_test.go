// SPDX-License-Identifier: MIT
// Copyright Bernardo Soares

//go:build purego && !cgo

package supertrapezoid

import (
	"fmt"
	"runtime"
	"testing"

	"supertrapezoid/pkg/trapezoid"
	"supertrapezoid/pkg/trapezoid/purego"
	"supertrapezoid/pkg/trapezoid/utils"

	"github.com/stretchr/testify/require"
)

func TestTrapezoidPurego(t *testing.T) {
	t.Run("scalar", func(t *testing.T) {
		result := purego.TrapezoidPureGo(trapezoid.Lower, trapezoid.Upper, trapezoid.Resolution)
		require.NotZero(t, result)
		require.InDelta(t, trapezoid.TestResult, result, 1e-9)
	})

	utils.PrintAVXSupport()
	t.Run("SIMD-portable", func(t *testing.T) {
		result := purego.TrapezoidSIMDPureGo(trapezoid.Lower, trapezoid.Upper, trapezoid.Resolution)
		require.NotZero(t, result)
		require.InDelta(t, trapezoid.TestResult, result, 1e-9)
	})

	t.Run("SIMD-128", func(t *testing.T) {
		result := purego.TrapezoidArchSIMD128PureGo(trapezoid.Lower, trapezoid.Upper, trapezoid.Resolution)
		require.NotZero(t, result)
		require.InDelta(t, trapezoid.TestResult, result, 1e-9)
	})

	t.Run("SIMD-256", func(t *testing.T) {
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
func BenchmarkPureGo(b *testing.B) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	b.Run("scalar", func(b *testing.B) {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("%d", tc), func(b *testing.B) {
				var result float64
				withPerfMetrics(b, func() {
					result = purego.TrapezoidPureGo(trapezoid.Lower, trapezoid.Upper, tc)
				})

				require.NotZero(b, result)
			})
		}
	})

	b.Run("SIMD-portable", func(b *testing.B) {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("%d", tc), func(b *testing.B) {
				var result float64
				withPerfMetrics(b, func() {
					result = purego.TrapezoidSIMDPureGo(trapezoid.Lower, trapezoid.Upper, tc)
				})

				require.NotZero(b, result)
			})
		}
	})

	b.Run("SIMD-128", func(b *testing.B) {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("%d", tc), func(b *testing.B) {
				var result float64
				withPerfMetrics(b, func() {
					result = purego.TrapezoidArchSIMD128PureGo(trapezoid.Lower, trapezoid.Upper, tc)
				})

				require.NotZero(b, result)
			})
		}
	})

	b.Run("SIMD-256", func(b *testing.B) {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("%d", tc), func(b *testing.B) {
				var result float64
				withPerfMetrics(b, func() {
					result = purego.TrapezoidArchSIMD256PureGo(trapezoid.Lower, trapezoid.Upper, tc)
				})

				require.NotZero(b, result)
			})
		}
	})
}
