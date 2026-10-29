// SPDX-License-Identifier: MIT
// Copyright Bernardo Soares

package utils

import (
	"fmt"
	"simd"
	"simd/archsimd"
)

func PrintAVXSupport() {
	var f64Type simd.Float64s
	laneCount := f64Type.Len()

	widthBits := simd.VectorBitSize()
	isEmulated := simd.Emulated()

	// Platform independent SIMD
	fmt.Printf("--- Runtime SIMD Profile ---\n")
	fmt.Printf("Vector Register Size : %d-bit\n", widthBits)
	fmt.Printf("Parallel Lane Width  : %d elements\n", laneCount)
	fmt.Printf("Is Software Emulated : %t\n", isEmulated)

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

	fmt.Printf("--- ARCHSIMD ---\n")

	fmt.Println("AVX512 supported?", archsimd.X86.AVX512())
	fmt.Println("AVX2 supported?", archsimd.X86.AVX2())
	fmt.Println("AVX supported?", archsimd.X86.AVX())
}
