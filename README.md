# Interesting Interview Questions

A collection of interview programming questions curated through
agentic engineering with [opencode](https://opencode.ai)
and the `big-pickle` model. Every question, difficulty label, and
convention in this repo was shaped through iterative agent-assisted
development — not batch-generated or scrape-harvested.

## Difficulty Scale

### Easy — Fizz Buzz

Print numbers 1 to n, replacing multiples of 3 with "Fizz", multiples of 5 with "Buzz",
and multiples of both with "FizzBuzz".

One loop, conditionals, modulo. No data structures.
The only subtlety: order of conditionals matters.

### Medium — Product of Array Except Self

Given an integer array `nums`, return an array `output` where
`output[i]` is the product of all elements of `nums` except `nums[i]`.

The constraint: O(n) time, no division.
Division makes this trivial. Without it, the obvious approach
(nested loops) is O(n²), and the O(n) path is not immediate.

### Hard — LRU Cache

Design a data structure supporting `get` and `put` in O(1) average time,
evicting the least recently used entry when full.

O(1) get suggests a hash map. O(1) put with eviction means
something else must track recency. Combining both correctly
is where candidates get tangled.

### Expert — Serialize and Deserialize Binary Tree

Convert a binary tree to a string representation and back.

Straightforward at first glance, but the format must encode
structure unambiguously. Edge cases — skewed trees, empty trees,
deep recursion — separate solid solutions from shaky ones.
