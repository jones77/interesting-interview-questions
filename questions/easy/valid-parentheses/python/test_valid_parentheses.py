import unittest

from valid_parentheses import is_valid


class TestIsValid(unittest.TestCase):
    def test_is_valid(self):
        tests = [
            ("()", True),
            ("()[]{}", True),
            ("([])", True),
            ("(]", False),
            ("([)]", False),
            ("", False),
            ("(", False),
        ]
        for s, want in tests:
            with self.subTest(s=s):
                self.assertEqual(is_valid(s), want)


if __name__ == "__main__":
    unittest.main()
