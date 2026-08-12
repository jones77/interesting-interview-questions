import unittest

from regular_expression_matching import is_match


class TestIsMatch(unittest.TestCase):
    def test_is_match(self):
        tests = [
            ("aa", "a", False),
            ("aa", "a*", True),
            ("ab", ".*", True),
            ("aab", "c*a*b", True),
            ("mississippi", "mis*is*p*.", False),
            ("", "c*", True),
        ]
        for s, p, want in tests:
            with self.subTest(s=s, p=p):
                got = is_match(s, p)
                self.assertEqual(got, want)


if __name__ == "__main__":
    unittest.main()
