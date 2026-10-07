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
	result := cgo.TrapezoidCgo(trapezoid.Lower, trapezoid.Upper, trapezoid.Resolution)
	require.NotZero(t, result)
	require.Equal(t, trapezoid.TestResult, result)
}

func BenchmarkEmptyFunctionCallCgo(b *testing.B) {
	var err error

	for range b.N {
		err = cgo.DoNothingCgo()
		require.NoError(b, err)
	}
}

func BenchmarkTrapezoidCgo(b *testing.B) {
	b.Run("trapezoidCgo-1000", func(b *testing.B) {
		var result float64
		for range b.N {
			result = cgo.TrapezoidCgo(trapezoid.Lower, trapezoid.Upper, 1000)
		}

		require.NotZero(b, result)
	})

	b.Run("trapezoidCgo-1000000", func(b *testing.B) {
		var result float64
		for range b.N {
			result = cgo.TrapezoidCgo(trapezoid.Lower, trapezoid.Upper, 1000000)
		}

		require.NotZero(b, result)
	})
}
