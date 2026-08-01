# Word Ladder

Source: https://leetcode.com/problems/word-ladder/

A transformation sequence from word `beginWord` to word `endWord` using a dictionary `wordList` is a sequence of words `beginWord = s0, s1, ..., sk = endWord` such that:

- Each adjacent pair of words differs by a single letter.
- Every `si` for `1 <= i <= k` is in `wordList`.
- `beginWord` does not need to be in `wordList`.

Return the number of words in the shortest transformation sequence, or `0` if no such sequence exists.

**Example 1:**

```
Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
Output: 5
```

**Example 2:**

```
Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
Output: 0
```

**Constraints:**

- `1 <= beginWord.length <= 10`
- `endWord.length == beginWord.length`
- `1 <= wordList.length <= 5000`
- All words consist of lowercase English letters.
