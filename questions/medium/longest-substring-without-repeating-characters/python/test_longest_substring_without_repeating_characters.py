import unittest

from longest_substring_without_repeating_characters import length_of_longest_substring


class TestLengthOfLongestSubstring(unittest.TestCase):
    def test_examples(self):
        tests = [
            ("abcabcbb", 3),
            ("bbbbb", 1),
            ("pwwkew", 3),
            ("", 0),
            (" ", 1),
        ]
        for s, want in tests:
            with self.subTest(s=s):
                got = length_of_longest_substring(s)
                self.assertEqual(got, want)


if __name__ == "__main__":
    unittest.main()
