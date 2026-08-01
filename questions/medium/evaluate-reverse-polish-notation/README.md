# Evaluate Reverse Polish Notation

Source: https://leetcode.com/problems/evaluate-reverse-polish-notation/

Evaluate the value of an arithmetic expression in Reverse Polish
Notation. Valid operators are `+`, `-`, `*`, and `/`. Each operand may
be an integer or another expression. Division between two integers
always truncates toward zero.

**Example 1:**

```
Input: tokens = ["2","1","+","3","*"]
Output: 9
Explanation: ((2 + 1) * 3) = 9
```

**Example 2:**

```
Input: tokens = ["4","13","5","/","+"]
Output: 6
Explanation: (4 + (13 / 5)) = 6
```

**Constraints:**

- `1 <= tokens.length <= 10^4`
- `tokens[i]` is an operator (`+`, `-`, `*`, `/`) or an integer in range `[-200, 200]`
