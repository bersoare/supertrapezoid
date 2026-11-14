# Supertrapezoid

**An experimental study of SIMD acceleration performance in Go in comparison with CGO**

Supertrapezoid is an academic research project developed as part of a Master's degree.

The project investigates the performance characteristics of SIMD vectorization using Pure Go (using GOEXPERIMENTAL=simd) against CGO.

It explores scalar and SIMD execution, Pure Go and CGO implementations, and the overhead associated with interoperability between Go and native code.

The primary objective is to provide an experimental environment for studying low-level optimization techniques, compiler behavior and performance measurement.

## Research Objectives

The project investigates the following questions:

- How does SIMD acceleration affect the performance of numerical integration compared to scalar execution?
- How do Pure Go implementations compare with native implementations accessed through CGO?
- What is the performance overhead associated with Go/C interoperability?
- How effectively does the Go toolchain generate SIMD instructions?
- How do different vector widths influence computational performance?

And probably the most important one:
- What trade-offs exist between portability, implementation complexity, and execution speed?

## Experimental Scope

The experiments cover several implementation strategies:

- Scalar numerical integration in Pure Go.
- SIMD-oriented numerical integration in Pure Go.
- Native scalar numerical integration accessed through CGO.
- Native 128-bit and 256-bit SIMD implementations.
- Benchmarking, CPU profiling, and machine-code inspection.

The results are intended to support comparative analysis and discussion within the context of the academic project.

## Implementations

The project explores several implementations of the composite trapezoidal rule.

| Implementation | Description |
|---|---|
| Scalar | Conventional sequential implementation (Go and CGO) |
| Portable SIMD | SIMD-oriented implementation using the platform agnostic API (Go only) |
| SIMD-128 | 128-bit SIMD implementation with intrinsics / archsimd |
| SIMD-256 | 256-bit SIMD implementation with intrinsics / archsimd |

These implementations provide a common basis for comparing execution strategies and investigating their respective performance characteristics.

## Requirements

- Go 1.27.2
- GNU Make.
- A C compiler and CGO-compatible toolchain for native implementations.
- An x86-64 processor supporting the required instruction sets.

The Pure Go build configuration uses `GOAMD64=v3`, establishing an x86-64-v3 instruction-set baseline.

The project also uses Go's experimental SIMD functionality through `GOEXPERIMENT=simd`.

## Getting Started

Clone the repository:

```bash
git clone https://github.com/bersoare/supertrapezoid.git
cd supertrapezoid/code
```

Download dependencies:

```bash
go mod download
```

## Building and Testing

The Makefile provides targets for building and testing both implementation configurations.

### Run all tests

```bash
make test
```

This runs both Pure Go and CGO test suites.

### Run Pure Go tests

```bash
make test-purego
```

This configuration disables CGO and enables the experimental SIMD implementation.

### Run CGO tests

```bash
make test-cgo
```

This configuration enables native CGO implementations.

### Build test binaries

```bash
make build
```

Individual configurations can also be built separately:

```bash
make build-purego
make build-cgo
```

Build-related arguments can be passed through `BUILD_ARGS`.

## Benchmarking

Benchmarking is a central part of the experimental methodology.

The project provides benchmarks for comparing implementation strategies across different integration resolutions.

Run the complete benchmark workflow:

```bash
make bench
```

This builds both configurations and executes their respective benchmark suites.

By default, benchmarks run for 50 iterations per benchmark case.

### Benchmark output

The Makefile supports JSON-formatted benchmark output. Ideal for programatically parsing results:

```bash
mkdir -p results
make bench-json
```

Results are written to timestamped files under `results/`.

Additional benchmark arguments can be supplied through `BENCH_ARGS`, while output configuration can be controlled through `OUTPUT_ARGS`.

#### Core pinning

To pin the execution to a certain cpu thread, wrap the command with `GOMAXPROCS=1 taskset -c 14`

### Bench cases

write the bench cases here

## Performance Analysis

The project includes additional tooling for investigating performance beyond execution time.

### CPU profiling

Generate and inspect CPU profiles:

```bash
make profile-cgo
```

```bash
make profile-purego
```

These targets generate CPU profiles and launch the Go `pprof` interface.

The default local endpoints are:

- CGO: http://localhost:8081
- Pure Go: http://localhost:8082

### Assembly inspection

Generate disassembly output for both configurations:

```bash
make asm
```

Or inspect them independently:

```bash
make asm-cgo
make asm-purego
```

Disassembly is written to:

- `cgo/asm.txt`
- `purego/asm.txt`

This allows inspection of the machine instructions generated by the compiler and investigation of whether expected vectorization and optimization strategies are being applied.

## Numerical Method

The project uses the composite trapezoidal rule to approximate definite integrals (https://en.wikipedia.org/wiki/Trapezoidal_rule).

To keep it simple, we use the function `f(x) = x*x-x` (equivalent to `f(x) = x(x-1)`)

The focus of this study is not the numerical correctness: the sole goal of this algorihm is to provide arithmetical intensity for the CPU to run.

## Research Methodology

The project uses comparative benchmarking to investigate the performance characteristics of different implementation strategies.

The main dimensions of comparison include:

- **Execution strategy:** Scalar versus SIMD.
- **Implementation language:** Pure Go versus Cgo.
- **Vector width:** 128-bit versus 256-bit SIMD.
- **Workload size:** Different integration resolutions.

Benchmarking is complemented by CPU profiling and assembly inspection to provide additional insight into observed performance differences.

For meaningful comparisons, measurements should be performed under consistent hardware and software conditions, documenting relevant environmental factors such as:

- CPU model and supported instruction sets
- Go compiler and toolchain version
- Operating system and architecture
- Build configuration and compiler flags
- Benchmark parameters
- Environment isolation (ex: `isolcpus`)

## Makefile Reference

The Makefile provides a collection of targets for managing the experimental workflow.

| Target | Description |
|---|---|
| `clean` | Remove generated binaries, assembly and profiles |
| `test` | Run Pure Go and CGO tests |
| `test-purego` | Run Pure Go tests |
| `test-cgo` | Run CGO tests |
| `build` | Build both configurations |
| `build-purego` | Build Pure Go test binaries |
| `build-cgo` | Build CGO test binaries |
| `bench` | Build and run both benchmark suites |
| `bench-json` | Run benchmarks with JSON output |
| `bench-cgo-profile` | Generate a CGO CPU profile |
| `bench-purego-profile` | Generate a Pure Go CPU profile |
| `asm` | Generate assembly output for both configurations |
| `asm-cgo` | Generate CGO disassembly |
| `asm-purego` | Generate Pure Go disassembly |
| `profile-cgo` | Open the CGO CPU profile |
| `profile-purego` | Open the Pure Go CPU profile |

## Project Status and Limitations

This repository contains experimental implementations and benchmarking infrastructure developed for academic research.

It is not intended to be a production-ready numerical integration library.

The implementations, dependencies, and build configurations may evolve throughout the research process.

Performance results should be interpreted in the context of the hardware, compiler, runtime, and experimental methodology used. Results obtained on different systems may not be directly comparable.

The use of experimental Go functionality also means that compatibility and behavior may depend on the specific toolchain version.

## License

This project is licensed under the MIT License. See [LICENSE](../LICENSE) for details.

---

**Author:** Bernardo Soares

**Repository:** https://github.com/bersoare/supertrapezoid