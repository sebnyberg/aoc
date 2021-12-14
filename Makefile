.PHONY: test
test:
	@go test ./aoc2021/...



.PHONY: bench
bench:
	@bash bench.sh