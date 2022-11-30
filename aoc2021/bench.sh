#!/usr/bin/env bash

set -eu

rm -rf tmp
mkdir tmp
touch tmp/res.txt

for i in {1..5} ; do
    go test -run=None -bench=Day19 -benchmem ./aoc2021/... | grep Part >> tmp/res.txt
done

benchstat tmp/res.txt
rm -rf tmp

# go test -count 1 -run=None -bench=Day19Part2 -benchmem -cpuprofile=cpu.pprof -memprofile=mem.pprof -trace=trace.out ./aoc2021/day19/...