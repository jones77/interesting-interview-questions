# Regular Expression Matching

Source: https://leetcode.com/problems/regular-expression-matching/

Implement regular expression matching with support for `'.'` and `'*'`.

- `'.'` Matches any single character.
- `'*'` Matches zero or more of the preceding element.

The matching should cover the **entire** input string (not partial).

**Example 1:**

```
Input: s = "aa", p = "a"
Output: false
```

**Example 2:**

```
Input: s = "aa", p = "a*"
Output: true
```

**Example 3:**

```
Input: s = "ab", p = ".*"
Output: true
```

**Example 4:**

```
Input: s = "aab", p = "c*a*b"
Output: true
```

**Constraints:**

- `1 <= s.length <= 20`
- `1 <= p.length <= 30`
- `s` contains only lowercase English letters.
- `p` contains only lowercase English letters, `'.'`, and `'*'`.
