package burstballoons

/*
maxCoins solves Burst Balloons.

You are given `n` balloons, indexed from `0` to `n - 1`. Each balloon is
painted with a number on it represented by an array `nums`. You are
asked to burst all the balloons. If you burst the i-th balloon, you
get `nums[i - 1] * nums[i] * nums[i + 1]` coins. After bursting `i`,
`i - 1` and `i + 1` become adjacent. Return the maximum coins you can
collect by bursting the balloons wisely.

You may imagine `nums[-1] = nums[n] = 1`. They are not real, you
cannot burst them.

**Example 1:**

```
Input: nums = [3,1,5,8]
Output: 167
Explanation:
nums = [3,1,5,8] → [3,5,8] → [3,8] → [8] → []
coins = 3*1*5 + 3*5*8 + 1*3*8 + 1*8*1 = 167
```

**Example 2:**

```
Input: nums = [1,5]
Output: 10
```

**Constraints:**

- `n == nums.length`
- `1 <= n <= 300`
- `0 <= nums[i] <= 100`
*/
func MaxCoins(nums []int) int {
	return 0
}
