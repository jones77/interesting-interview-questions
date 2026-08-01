# Count of Smaller Numbers After Self

Source: https://leetcode.com/problems/count-of-smaller-numbers-after-self/

Given an integer array `nums`, return an integer array `counts` where
`counts[i]` is the number of smaller elements to the right of `nums[i]`.

**Example 1:**

```
Input: nums = [5,2,6,1]
Output: [2,1,1,0]
Explanation:
- 5 has 2 smaller: 2 and 1
- 2 has 1 smaller: 1
- 6 has 1 smaller: 1
- 1 has 0 smaller
```

**Example 2:**

```
Input: nums = [-1,-1]
Output: [0,0]
```

**Constraints:**

- `1 <= nums.length <= 10^5`
- `-10^4 <= nums[i] <= 10^4`
