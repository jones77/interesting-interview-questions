import unittest

from reverse_string import reverse_string


class TestReverseString(unittest.TestCase):
    def test_reverse_string(self):
        tests = [
            (["h", "e", "l", "l", "o"], ["o", "l", "l", "e", "h"]),
            (["H", "a", "n", "n", "a", "h"], ["h", "a", "n", "n", "a", "H"]),
        ]
        for s, want in tests:
            with self.subTest(s=s):
                reverse_string(s)
                self.assertEqual(s, want)


if __name__ == "__main__":
    unittest.main()
