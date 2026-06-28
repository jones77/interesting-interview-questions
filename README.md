# Interesting Interview Questions

A curated collection of interview programming questions organized by difficulty.

## Difficulty Scale

### Easy — Fizz Buzz

Print numbers 1 to n, replacing multiples of 3 with "Fizz", multiples of 5 with "Buzz",
and multiples of both with "FizzBuzz".

One loop, conditionals, modulo. No data structures.
One trap: check divisibility by 15 before 3 or 5.

### Medium — Two Sum

Given an array of integers, find the two numbers that add up to a target
and return their indices.

Brute force is O(n²). O(n) needs a hash map:
store each number, check if `target - num` was already seen.
The trick is obvious once shown, unobvious before.

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
