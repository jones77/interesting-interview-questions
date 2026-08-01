# Minimum Path Sum

Source: https://leetcode.com/problems/minimum-path-sum/

Given a `m x n` grid filled with non-negative numbers, find a path from
top-left to bottom-right which minimises the sum of all numbers along
the path. You can only move down or right at any point in time.

**Example 1:**

```
Input: grid = [[1,3,1],[1,5,1],[4,2,1]]
Output: 7
Explanation: 1 → 3 → 1 → 1 → 1
```

**Example 2:**

```
Input: grid = [[1,2,3],[4,5,6]]
Output: 12
```

**Constraints:**

- `m == grid.length`
- `n == grid[i].length`
- `1 <= m, n <= 200`
- `0 <= grid[i][j] <= 200`
