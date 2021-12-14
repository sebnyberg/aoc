# :santa: :christmas_tree: Advent of Code in Go :santa: :christmas_tree:

Advent of Code solutions in Go.

Helper functionality can be found in `ax`. Really looking forward to generics
in three months!

I typically make comments in the code only if the exercise is medium-hard. These
comments usually discuss whether exhaustive search is plausible from a time
complexity point of view.

## 2021 Benchmarks

| day    | part   | time/op     | alloc/op    | allocs/op  |
| ------ | ------ | ----------- | ----------- | ---------- |
| Day 1  | Part 1 | 128µs ± 6%  | 148kB ± 0%  | 2.02k ± 0% |
|        | Part 2 | 127µs ± 9%  | 148kB ± 0%  | 2.02k ± 0% |
| Day 2  | Part 1 | 417µs ± 4%  | 179kB ± 0%  | 3.04k ± 0% |
|        | Part 2 | 421µs ± 7%  | 179kB ± 0%  | 3.04k ± 0% |
| Day 3  | Part 1 | 110µs ± 2%  | 70.8kB ± 0% | 1.02k ± 0% |
|        | Part 2 | 486µs ± 4%  | 174kB ± 0%  | 1.15k ± 0% |
| Day 4  | Part 1 | 372µs ± 5%  | 320kB ± 0%  | 2.54k ± 0% |
|        | Part 2 | 445µs ± 2%  | 320kB ± 0%  | 2.55k ± 0% |
| Day 5  | Part 1 | 569µs ± 8%  | 118kB ± 0%  | 1.55k ± 0% |
|        | Part 2 | 887µs ±31%  | 118kB ± 0%  | 1.55k ± 0% |
| Day 6  | Part 1 | 24.3µs ±31% | 9.78kB ± 0% | 10.0 ± 0%  |
|        | Part 2 | 24.8µs ± 4% | 9.79kB ± 0% | 10.0 ± 0%  |
| Day 7  | Part 1 | 48.9µs ± 4% | 24.8kB ± 0% | 10.0 ± 0%  |
|        | Part 2 | 3.04ms ± 7% | 33.0kB ± 0% | 11.0 ± 0%  |
| Day 8  | Part 1 | 72.6µs ± 2% | 50.3kB ± 0% | 616 ± 0%   |
|        | Part 2 | 166µs ± 5%  | 82.4kB ± 0% | 816 ± 0%   |
| Day 9  | Part 1 | 268µs ± 5%  | 112kB ± 0%  | 216 ± 0%   |
|        | Part 2 | 398µs ± 2%  | 286kB ± 0%  | 220 ± 0%   |
| Day 10 | Part 1 | 110µs ± 2%  | 21.6kB ± 0% | 321 ± 0%   |
|        | Part 2 | 115µs ± 2%  | 22.5kB ± 0% | 323 ± 0%   |
| Day 11 | Part 1 | 74.1µs ±16% | 4.90kB ± 0% | 22.0 ± 0%  |
|        | Part 2 | 179µs ± 4%  | 4.90kB ± 0% | 22.0 ± 0%  |
| Day 12 | Part 1 | 51.4µs ± 2% | 67.6kB ± 0% | 79.0 ± 0%  |
|        | Part 2 | 116µs ± 9%  | 92.5kB ± 0% | 80.0 ± 0%  |
| Day 13 | Part 1 | 208µs ± 0%  | 183kB ± 0%  | 2.05k ± 0% |
|        | Part 2 | 212µs ± 2%  | 111kB ± 0%  | 2.06k ± 0% |
| Day 14 | Part 1 | 1.12ms ± 1%  | 338kB ± 0% | 19.8k ± 0% |
|        | Part 2 | 75.3µs ± 5% | 25.9kB ± 0% | 384 ± 0% |
