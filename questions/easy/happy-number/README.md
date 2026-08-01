# Happy Number

Source: https://leetcode.com/problems/happy-number/

A happy number is a number defined by the following process:
Starting with any positive integer, replace the number by the sum of
the squares of its digits, and repeat until the number equals 1, or
it loops endlessly in a cycle that does not include 1. Return true
if the number is happy.

**Example 1:**

```
Input: n = 19
Output: true
Explanation: 1² + 9² = 82, 8² + 2² = 68, 6² + 8² = 100, 1² + 0² + 0² = 1
```

**Example 2:**

```
Input: n = 2
Output: false
```

**Constraints:**

- `1 <= n <= 2³¹ - 1`
