// SPDX-License-Identifier: MIT
// Copyright Bernardo Soares

package cgo

/*
#cgo CFLAGS: -O3 -msse4.2 -mavx

#include <stddef.h>
#include <stdio.h>

void do_nothing() {}

double integrableFunc(double x) {
	return x * x -x;
}

double trapezoid(double lower, double upper, int resolution) {
	double h = (upper - lower) / (double) resolution;

	double sum = .5 * (integrableFunc(lower) + integrableFunc(upper));
	for (int i = 1; i < resolution; i++) {
		sum += integrableFunc(lower + ((double) i) * h);
	}

	return sum*h;
}

// SIMD implementation below.
// We implement for both 2 and 4 lanes
// this way we can compare it vs platform independent SIMD
// in pure go.

#if defined(__AVX__) || defined(__AVX2__)
    #define LANE_COUNT 4
    #define BYTE_SIZE  32  // 4 doubles * 8 bytes = 32 bytes (256-bit AVX)
#else
    #define LANE_COUNT 2
    #define BYTE_SIZE  16  // 2 doubles * 8 bytes = 16 bytes (128-bit SSE2/Neon)
#endif

typedef double vdf __attribute__ ((vector_size (BYTE_SIZE)));

double simd_trapezoid_128bit(double lower, double upper, int resolution) {
	double h = (upper - lower) / (double) resolution;
	double sum = .5 * (integrableFunc(lower) + integrableFunc(upper));

	int blocks = resolution / LANE_COUNT;

	// todo: declare constants and vars

	for (int block = 0; block < blocks; block++) {
	// todo: put code in here
		continue;
	}

	for (int i = 1 + (blocks * width); i < resolution; i++) {
		sum += integrableFunc(lower + ((double) i) * h);
	}

	return sum*h;
}

double simd_trapezoid_256bit(double lower, double upper, int resolution) {
	double h = (upper - lower) / (double) resolution;
	double sum = .5 * (integrableFunc(lower) + integrableFunc(upper));

	int blocks = resolution / LANE_COUNT;

	// todo: declare constants and vars

	for (int block = 0; block < blocks; block++) {
	// todo: put code in here
		continue;
	}

	for (int i = 1 + (blocks * width); i < resolution; i++) {
		sum += integrableFunc(lower + ((double) i) * h);
	}

	return sum*h;
}

*/
import "C"

func DoNothingCgo() error {
	C.do_nothing()
	return nil
}

func TrapezoidCgo(lower, upper float64, resolution int) float64 {
	result := C.trapezoid(C.double(lower), C.double(upper), C.int(resolution))
	return float64(result)
}

func TrapezoidSIMDCgo(lower, upper float64, resolution int) float64 {
	result := C.simd_trapezoid(C.double(lower), C.double(upper), C.int(resolution))
	return float64(result)
}
