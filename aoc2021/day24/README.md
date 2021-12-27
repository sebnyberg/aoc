# Day 24

This day took a really long time to figure out, so I thought it would be best to include a README.

The first order of business is to read the input and try to parse instructions into some kind of simplified set of rules.

A good thing to note is that only three lines are possibly different between instruction sets:

* Line 4: divisor (henceforth `d`)
* Line 5: constant added to x (henceforth `a`)
* Line 15: constant added to y (henceforth `b`)

```go
if z % 26 + a == w {
  x := 0
  z := z / div
  z *= (25*x + 1)
  z += x*(w+b)
} else {
  x := 1
  z := z / div
  z *= (25*x + 1)
  z += x*(w+b)
}
```

This can be simplified as:

```go
cond := z%26 + a == w
z /= div
if !cond {
  z = z*26 + w + b
}
```

When the divisor is `1`, the constant `a` happens to be positive and greater than 10 (see input). This means that the condition will always fail, resulting in the `z*26` operation.

There are 7 such `div z 1` operations in the input. This means that `z*26` will happen at least 7 times. Since the goal is to have z be zero (0), the other 7 operations must undo these actions somehow.

This tells us that whenever there is `div z 26`, the condition must be true, i.e. `z%26+a == w` must be true.

This is how the input `w` is formed: each digit must be such that it is possible to match relevent criteria.

For my input this gives:

1. [d1+7]
2. [d1+7,d2+8]
3. [d1+7,d2+8,d3+10]
4. [d1+7,d2+8], [d3+10-2==d4]
5. [d1+7], [d3+10-2==d4,d2+8-10==d5]
6. [d1+7,d6+6], [d3+10-2==d4,d2+8-10==d5]
7. [d1+7], [d3+10-2==d4,d2+8-10==d5,d6+6-14==d7]
8. [], [d3+10-2==d4,d2+8-10==d5,d6+6-14==d7,d1+7-5==d8]
9. [d9+1], [d3+10-2==d4,d2+8-10==d5,d6+6-14==d7,d1+7-5==d8]
10. [d9+1,d10+8], [d3+10-2==d4,d2+8-10==d5,d6+6-14==d7,d1+7-5==d8]
11. [d9+1], [d3+10-2==d4,d2+8-10==d5,d6+6-14==d7,d1+7-5==d8,d10+8-14=d11]
12. [d9+1,d12+13], [d3+10-2==d4,d2+8-10==d5,d6+6-14==d7,d1+7-5==d8,d10+8-14=d11]
13. [d9+1], [d3+10-2==d4,d2+8-10==d5,d6+6-14==d7,d1+7-5==d8,d10+8-14=d11,d12+13-14=d13]
14. [], [d3+10-2==d4,d2+8-10==d5,d6+6-14==d7,d1+7-5==d8,d10+8-14=d11,d12+4-14=d13,d9+1-5=d14]

Constraints:

* d1 + 2 = d8
* d2 - 2 = d5
* d3 + 8 = d4
* d6 - 8 = d7
* d9 - 4 = d14
* d10 - 6 = d11
* d12 - 1 = d13

Maximizing constraints yield the max answer:

7......9......
