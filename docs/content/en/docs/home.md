---
title: Quick Start
type: docs
---
Get started with 🚀🇱🅰🆖 quickly with these examples:

```js
let input = open("examples/aoc/2021/day-1/input").lines()


let a = []
foreach i, number in input {
  a.yoink(number.strip().plz_i())
}
input = a

let increase = 0
foreach i, number in input {
  if (number > input[i-1]) {
    increase = increase + 1
  }
}
puts(increase + 1)

increase = 0
foreach i, number in input {
  let sum = number + input[i+1] + input[i+2]
  let sum_two = input[i+1] + input[i+2] + input[i+3]
  
  if (sum_two > sum) {
    increase = increase + 1
  }
}
puts(increase + 1)
```