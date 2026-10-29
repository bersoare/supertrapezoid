// SPDX-License-Identifier: MIT
// Copyright Bernardo Soares

package cgo

/*
#cgo CFLAGS: -O3 -mavx -march=native

#include <stddef.h>
#include <stdio.h>
#include <immintrin.h>

void do_nothing() {}

static inline double __attribute__((always_inline)) integrableFunc(double x) {
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
// in pure go

static inline __m128d __attribute__((always_inline)) integrableFuncSIMD128(__m128d x) {
	return _mm_fmsub_pd(x, x, x);
}

// for AVX-128: we can have up to 2 doubles (2x64b) on a vector.
double simd_trapezoid_128bit(double lower, double upper, int resolution) {
	double h = (upper - lower) / (double) resolution;
	double sum = .5 * (integrableFunc(lower) + integrableFunc(upper));

	int blocks = resolution / 2;

	double steps[2] = {0,1};

	__m128d vec_h = _mm_set1_pd(h);
	__m128d vec_steps = _mm_loadu_pd(steps);
	__m128d vec_lower = _mm_set1_pd(lower);
	__m128d vec_block_step = _mm_set1_pd(h * 2);
	__m128d vec_sum = {};

	__m128d vec_x = vec_lower + ((vec_steps * vec_h) - vec_block_step);

	// todo: declare constants and vars

	for (int block = 0; block < blocks; block++) {
		vec_x += vec_block_step;
		vec_sum += integrableFuncSIMD128(vec_x);
	}

	for (int i = 1 + (blocks * 2); i < resolution; i++) {
		sum += integrableFunc(lower + ((double) i) * h);
	}

	__m128d high_lane = _mm_unpackhi_pd(vec_sum, vec_sum); // Copies lane 1 into lane 0
	__m128d combined  = _mm_add_pd(vec_sum, high_lane);   // Adds lane 0 + lane 1
	sum += _mm_cvtsd_f64(combined);                       // Extract scalar

	// sum += _mm_cvtsd_f64(_mm_hadd_pd(vec_sum, vec_sum));

	return sum*h;
}

static inline __m256d __attribute__((always_inline)) integrableFuncSIMD256(__m256d x) {
	return _mm256_fmsub_pd(x, x, x);
}

// for AVX-256: we can have up to 4 doubles (4x64b) on a vector.
double simd_trapezoid_256bit(double lower, double upper, int resolution) {
	double h = (upper - lower) / (double) resolution;
	double sum = .5 * (integrableFunc(lower) + integrableFunc(upper));

	int blocks = resolution / 4;

	double steps[4] = {0,1,2,3};

	__m256d vec_h = _mm256_set1_pd(h);
	__m256d vec_steps = _mm256_loadu_pd(steps);
	__m256d vec_lower = _mm256_set1_pd(lower);
	__m256d vec_block_step = _mm256_set1_pd(h * 4);
	__m256d vec_sum = {};

	__m256d vec_x = vec_lower + ((vec_steps * vec_h) - vec_block_step);

	// todo: declare constants and vars

	for (int block = 0; block < blocks; block++) {
		vec_x += vec_block_step;
		vec_sum += integrableFuncSIMD256(vec_x);
	}

	for (int i = 1 + (blocks * 4); i < resolution; i++) {
		sum += integrableFunc(lower + ((double) i) * h);
	}

	// sum += _mm256_cvtsd_f64(_mm256_hadd_pd(vec_sum, vec_sum));
	__m256d perm = _mm256_permute2f128_pd(vec_sum, vec_sum, 1);
    __m256d sum_halves = _mm256_add_pd(vec_sum, perm);
    __m256d shifted = _mm256_permute_pd(sum_halves, 5);
    __m256d final_vec = _mm256_add_pd(sum_halves, shifted);

    sum += _mm256_cvtsd_f64(final_vec);

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

func TrapezoidSIMDAVX128Cgo(lower, upper float64, resolution int) float64 {
	result := C.simd_trapezoid_128bit(C.double(lower), C.double(upper), C.int(resolution))
	return float64(result)
}

func TrapezoidSIMDAVX256Cgo(lower, upper float64, resolution int) float64 {
	result := C.simd_trapezoid_256bit(C.double(lower), C.double(upper), C.int(resolution))
	return float64(result)
}
