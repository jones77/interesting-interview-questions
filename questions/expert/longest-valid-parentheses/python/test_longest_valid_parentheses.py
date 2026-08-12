import unittest

from longest_valid_parentheses import longest_valid_parentheses


class TestLongestValidParentheses(unittest.TestCase):
    def test_longest_valid_parentheses(self):
        tests = [
            ("(()", 2),
            (")()())", 4),
            ("", 0),
            ("()(()", 2),
            ("()(())", 6),
        ]
        for s, want in tests:
            with self.subTest(s=s):
                got = longest_valid_parentheses(s)
                self.assertEqual(got, want)


if __name__ == "__main__":
    unittest.main()
