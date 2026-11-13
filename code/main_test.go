// SPDX-License-Identifier: MIT
// Copyright Bernardo Soares

package supertrapezoid

import (
	"testing"

	perf "github.com/elastic/go-perf"
)

var (
	testCases = []int{
		1000,
		1000000,
		500000000,
	}
)

func withPerfMetrics(b *testing.B, f func()) {
	var g perf.Group

	// warmup
	for range 3 {
		f()
	}

	g.Add(perf.CPUCycles)
	g.Add(perf.Instructions)
	g.Add(perf.CacheMisses)

	hw, err := g.Open(perf.CallingThread, perf.AnyCPU)
	if err != nil {
		b.Fatalf("Failed to initialize hardware counters: %v (Verify sudo permissions or paranoid flags)", err)
	}
	defer hw.Close()

	b.ResetTimer()

	stats, err := hw.MeasureGroup(func() {
		for range b.N {
			f()
		}
	})

	if err != nil {
		b.Fatalf("Performance counter sampling error: %v", err)
	}

	cyclesVal := float64(stats.Values[0].Value)
	instVal := float64(stats.Values[1].Value)
	cacheMissesVal := float64(stats.Values[2].Value)

	if cyclesVal > 0 {
		ipc := instVal / cyclesVal
		b.ReportMetric(ipc, "IPC")
	}
	b.ReportMetric(cacheMissesVal/float64(b.N), "cache-misses/op")

}
