# Interesting Interview Questions

A collection of interview programming questions assembled through agentic engineering
with [opencode](https://opencode.ai) and the `big-pickle` model.
Difficulty labels follow well-known problem ratings.
Browse all questions in the [`questions/`](questions/) directory.

## Difficulty Scale

The examples below illustrate the kind of problem you will find at each level.

### Easy — Fizz Buzz

Print numbers 1 to n, replacing multiples of 3 with "Fizz",
multiples of 5 with "Buzz", and multiples of both with "FizzBuzz".

One loop, conditionals, and modulo. No data structures needed.
The order of conditionals is the only trap.

### Medium — Product of Array Except Self

Given an integer array `nums`, return an array `output` where
`output[i]` is the product of all elements of `nums` except `nums[i]`.

The constraint requires O(n) time and bans division.
Division would make the problem simple.
Nested loops give O(n²), and the faster path needs a different angle.

### Hard — LRU Cache

Design a data structure supporting `get` and `put` in O(1) average time,
evicting the least recently used entry when full.

O(1) get suggests a hash map.
O(1) put with eviction means something else must track recency.
Combining both correctly is where candidates get tangled.

### Expert — Serialize and Deserialize Binary Tree

Convert a binary tree to a string representation and back.

The problem looks simple.
The format needs to encode tree structure clearly.
Edge cases like skewed trees, empty trees, and deep recursion
separate solid solutions from shaky ones.
