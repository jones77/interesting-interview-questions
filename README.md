# Interesting Interview Questions

Clone this repo, implement your own solutions, and make the tests pass.
Reference solutions are in `solutions.yaml` for when you want to check
your work — solve it first, then peek.

The questions and reference solutions were assembled through agentic engineering
with [opencode](https://opencode.ai) and the `big-pickle` model.
Difficulty labels follow well-known problem ratings.

## Difficulty Scale

The examples below illustrate the kind of problem you will find at each level.

### Easy — Fizz Buzz

[`questions/easy/fizzbuzz/fizzbuzz.md`](questions/easy/fizzbuzz/fizzbuzz.md)

Print numbers 1 to n, replacing multiples of 3 with "Fizz",
multiples of 5 with "Buzz", and multiples of both with "FizzBuzz".

One loop, conditionals, and modulo. No data structures needed.
The order of conditionals is the only trap.

### Medium — Product of Array Except Self

[`questions/medium/product-array-except-self/product-array-except-self.md`](questions/medium/product-array-except-self/product-array-except-self.md)

Given an integer array `nums`, return an array `output` where
`output[i]` is the product of all elements of `nums` except `nums[i]`.

The constraint requires O(n) time and bans division.
Division would make the problem simple.
Nested loops give O(n²), and the faster path needs a different angle.

### Hard — LRU Cache

[`questions/hard/lru-cache/lru-cache.md`](questions/hard/lru-cache/lru-cache.md)

Design a data structure supporting `get` and `put` in O(1) average time,
evicting the least recently used entry when full.

O(1) get suggests a hash map.
O(1) put with eviction means something else must track recency.
Combining both correctly is where candidates get tangled.

### Expert — Serialize and Deserialize Binary Tree

[`questions/expert/serialize-deserialize-binary-tree/serialize-deserialize-binary-tree.md`](questions/expert/serialize-deserialize-binary-tree/serialize-deserialize-binary-tree.md)

Convert a binary tree to a string representation and back.

The problem looks simple.
The format needs to encode tree structure clearly.
Edge cases like skewed trees, empty trees, and deep recursion
separate solid solutions from shaky ones.
