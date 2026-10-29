# SPDX-License-Identifier: MIT
# Copyright Bernardo Soares

clean:
	rm -rf cgo/*.test
	rm -rf purego/*.test
	rm -rf purego/asm.txt
	rm -rf purego/profile
	rm -rf cgo/asm.txt
	rm -rf cgo/profile

test: test-purego test-cgo

.PHONY:test-purego
test-purego:
	GOAMD64=v3 GOEXPERIMENT="simd" CGO_ENABLED=0 go test -tags=purego -v  ./...

.PHONY:test-cgo
test-cgo:
	GOEXPERIMENT="simd" go test -tags=cgo -v  ./...

build:build-purego build-cgo

build-purego:
	GOAMD64=v3 GOEXPERIMENT="simd" CGO_ENABLED=0 go test -tags=purego -v $(BUILD_ARGS) ./... -c -o purego/

build-cgo:
	GOEXPERIMENT="simd" go test -tags=cgo -v $(BUILD_ARGS) ./... -c -o cgo/

bench: build bench-cgo-base bench-purego-base

bench-cgo-base: build-cgo
	cgo/supertrapezoid.test -test.bench="Cgo" -test.v -test.benchmem -test.benchtime=50x $(BENCH_ARGS) $(OUTPUT_ARGS)

bench-purego-base: build-purego
	GOAMD64=v3 purego/supertrapezoid.test -test.bench="PureGo" -test.v -test.benchmem -test.benchtime=50x $(BENCH_ARGS) $(OUTPUT_ARGS)

OUTPUT_FILE := results/$(shell date +%Y-%m-%d-%T).json
bench-json: OUTPUT_ARGS += | go tool test2json | tee -a $(OUTPUT_FILE)
bench-json: build bench-cgo-base bench-purego-base

bench-cgo-profile: BENCH_ARGS += -test.cpuprofile cgo/profile
bench-cgo-profile: build-cgo bench-cgo-base

bench-purego-profile: BENCH_ARGS += -test.cpuprofile purego/profile
bench-purego-profile: build-purego bench-purego-base

asm-cgo: build-cgo
	go tool objdump -s trapezoid cgo/supertrapezoid.test > cgo/asm.txt

asm-purego: build-purego
	go tool objdump -s trapezoid purego/supertrapezoid.test > purego/asm.txt

asm: asm-cgo asm-purego

profile-cgo: bench-cgo-profile
	go tool pprof -http=localhost:8081 cgo/supertrapezoid.test cgo/profile

profile-purego: bench-purego-profile
	go tool pprof -http=localhost:8082 purego/supertrapezoid.test purego/profile

