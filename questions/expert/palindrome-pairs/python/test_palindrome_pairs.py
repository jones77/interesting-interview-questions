import unittest

from palindrome_pairs import palindrome_pairs


class TestPalindromePairs(unittest.TestCase):
    def test_palindrome_pairs(self):
        tests = [
            (["abcd", "dcba", "lls", "s", "sssll"], [[0, 1], [1, 0], [3, 2], [2, 4]]),
            (["bat", "tab", "cat"], [[0, 1], [1, 0]]),
            (["a", ""], [[0, 1], [1, 0]]),
        ]
        for words, want in tests:
            with self.subTest(words=words):
                got = palindrome_pairs(words)
                self.assertEqual(got, want)


if __name__ == "__main__":
    unittest.main()
