# SPDX-License-Identifier: MIT
# Copyright Bernardo Soares

bench-matrix:
	mkdir -p results/matrix
	go test -race -v -bench="Matrix" ./... -benchtime=1x -benchmem -json | tee results/matrix/$(shell date +%Y-%m-%d-%T).json

bench-utils:
	go test -race -v -bench="Utils" ./... -benchtime=100000x -benchmem

clean:
	rm -rf cgo/*.test
	rm -rf purego/*.test
	rm -rf purego/asm.txt
	rm -rf purego/profile
	rm -rf cgo/asm.txt
	rm -rf cgo/profile

test:test-purego test-cgo

.PHONY:test-purego
test-purego:
	GOAMD64=v3 GOEXPERIMENT="simd" CGO_ENABLED=0 go test -tags=purego -v  ./...

.PHONY:test-cgo
test-cgo:
	GOEXPERIMENT="simd" go test -tags=cgo -v  ./...

build:build-purego build-cgo

build-purego:
	GOAMD64=v3 GOEXPERIMENT="simd" CGO_ENABLED=0 go test -tags=purego -v ./... -c -o purego/

build-cgo:
	GOEXPERIMENT="simd" go test -tags=cgo -v -bench="Cgo" -benchtime=50x ./... -c -o cgo/

bench: build bench-cgo bench-purego

bench-cgo: build-cgo
	cgo/supertrapezoid.test -test.bench="Cgo" -test.v -test.benchmem -test.benchtime=50x -test.cpuprofile cgo/profile

bench-purego: build-purego
	GOAMD64=v3 purego/supertrapezoid.test -test.bench="PureGo" -test.v -test.benchmem -test.benchtime=50x -test.cpuprofile purego/profile

asm-cgo: build-cgo
	go tool objdump -s trapezoid cgo/supertrapezoid.test > cgo/asm.txt

asm-purego: build-purego
	go tool objdump -s trapezoid purego/supertrapezoid.test > purego/asm.txt

asm: asm-cgo asm-purego

profile-cgo: bench-cgo
	go tool pprof -http=localhost:8081 cgo/supertrapezoid.test cgo/profile

profile-purego: bench-purego
	go tool pprof -http=localhost:8082 purego/supertrapezoid.test purego/profile
