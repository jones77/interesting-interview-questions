import unittest

from minimum_window_substring import min_window


class TestMinWindow(unittest.TestCase):
    def test_min_window(self):
        tests = [
            ("ADOBECODEBANC", "ABC", "BANC"),
            ("a", "a", "a"),
            ("a", "aa", ""),
        ]
        for s, t, want in tests:
            with self.subTest(s=s, t=t):
                got = min_window(s, t)
                self.assertEqual(got, want)


if __name__ == "__main__":
    unittest.main()
