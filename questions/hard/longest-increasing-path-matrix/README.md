# Longest Increasing Path in a Matrix

Source: https://leetcode.com/problems/longest-increasing-path-in-a-matrix/

Given an `m x n` integers matrix, return the length of the longest
increasing path in the matrix. From each cell, you can move in four
directions: left, right, up, or down. You may not move diagonally
or outside the boundary.

**Example 1:**

```
Input: matrix = [[9,9,4],[6,6,8],[2,1,1]]
Output: 4
Explanation: 1 → 2 → 6 → 9 (or 1 → 2 → 6 → 9)
```

**Example 2:**

```
Input: matrix = [[3,4,5],[3,2,6],[2,2,1]]
Output: 4
Explanation: 3 → 4 → 5 → 6
```

**Constraints:**

- `m == matrix.length`
- `n == matrix[i].length`
- `1 <= m, n <= 200`
- `0 <= matrix[i][j] <= 2³¹ - 1`
