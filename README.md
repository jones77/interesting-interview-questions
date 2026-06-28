# Interesting Interview Questions

A curated collection of interview programming questions organized by difficulty.

## Difficulty Scale

### Easy — Fizz Buzz

Print numbers 1 to n, replacing multiples of 3 with "Fizz", multiples of 5 with "Buzz",
and multiples of both with "FizzBuzz".

One loop, conditionals, modulo. No data structures.
One trap: check divisibility by 15 before 3 or 5.

### Medium — Product of Array Except Self

Given an integer array `nums`, return an array `output` where
`output[i]` is the product of all elements of `nums` except `nums[i]`.

The constraint: O(n) time, no division.
Division makes this trivial (total product / self).
Without it, you need two passes building running products
from left and right. Intuitive once seen, but the constraint
blocks the obvious path.

### Hard — LRU Cache

Design a data structure supporting `get` and `put` in O(1) average time,
evicting the least recently used entry when full.

A hash map can't track recency.
Pair it with a doubly-linked list; update pointers on every get and put.
The idea is simple. The pointer bookkeeping is not.

### Expert — Serialize and Deserialize Binary Tree

Convert a binary tree to a string representation and back.

Preorder traversal with null markers preserves structure.
Encoding and parsing must agree on format.
Edge cases: skewed trees, empty trees, deep recursion.
Easy to describe. Easy to get wrong.
