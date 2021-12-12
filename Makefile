.PHONY: test
test:
	@go test ./aoc2021/...



.PHONY: bench
bench:
	@go test -run=None -bench=. ./aoc2021/... | grep "Part"