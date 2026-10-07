// SPDX-License-Identifier: MIT
// Copyright Bernardo Soares

package utils

import (
	"fmt"
	"simd"
)

func PrintAVXSupport() {
	var f64Type simd.Float64s
	laneCount := f64Type.Len()

	widthBits := simd.VectorBitSize()
	isEmulated := simd.Emulated()

	fmt.Printf("--- Runtime SIMD Profile ---\n")
	fmt.Printf("Vector Register Size : %d-bit\n", widthBits)
	fmt.Printf("Parallel Lane Width  : %d elements\n", laneCount)
	fmt.Printf("Is Software Emulated : %t\n", isEmulated)

	// Correlate metrics back to the exact compiler targets seen in your dumps
	switch {
	case isEmulated:
		fmt.Println("Active Variant       : @simd0 (Pure Go Software Baseline)")
	case laneCount == 2:
		fmt.Println("Active Variant       : @simd128 (SSE / AVX-128 / NEON)")
	case laneCount == 4:
		fmt.Println("Active Variant       : @simd256 (AVX2)")
	case laneCount == 8:
		fmt.Println("Active Variant       : @simd512 (AVX-512)")
	}
}
