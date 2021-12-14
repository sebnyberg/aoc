#!/usr/bin/env bash

set -eu

rm -rf tmp
mkdir tmp
touch tmp/res.txt

for i in {1..5} ; do
    go test -run=None -bench=Day13 -benchmem ./aoc2021/... | grep Part >> tmp/res.txt
done

benchstat tmp/res.txt
rm -rf tmp