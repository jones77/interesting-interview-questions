# Interesting Interview Questions

A curated collection of interview programming questions organized by difficulty.

## Difficulty Scale

### Easy — Fizz Buzz

Print numbers 1 to n, replacing multiples of 3 with "Fizz", multiples of 5 with "Buzz", and multiples of both with "FizzBuzz".

One loop, conditionals, modulo arithmetic — that's it. No data structures, no algorithms. The only trap: check divisibility by 15 before 3 or 5. If you can do that, you can solve it.

### Medium — Two Sum

Given an array of integers, find the two numbers that add up to a target and return their indices.

A brute-force nested loop works in O(n²). The O(n) solution requires the insight "for each number, look up whether `target - num` was already seen in a hash map." That conceptual leap — trading space for time — is what separates medium from easy: the problem is trivial once you see the trick, but it's not obvious on first glance.

### Hard — LRU Cache

Design a data structure that supports `get` and `put` in O(1) average time, evicting the least recently used entry when the cache reaches capacity.

A hash map alone can't track access order. The canonical solution pairs it with a doubly-linked list, updating pointers on every get and put. The difficulty isn't the idea — it's executing the pointer bookkeeping correctly under interview pressure without introducing cycles or losing nodes.

### Expert — Serialize and Deserialize Binary Tree

Convert a binary tree to a string representation and back.

Preorder traversal with null markers is straightforward to describe but surprisingly easy to botch. Null nodes must be encoded to preserve structure, the serialization format must be parseable unambiguously, and recursion depth edge cases (skewed trees, empty trees) all matter. The problem exposes whether you truly understand tree traversal or just memorized the textbook version.
