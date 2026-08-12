import unittest

from happy_number import is_happy


class TestIsHappy(unittest.TestCase):
    def test_is_happy(self):
        tests = [
            (19, True),
            (2, False),
            (1, True),
        ]
        for n, want in tests:
            with self.subTest(n=n):
                self.assertEqual(is_happy(n), want)


if __name__ == "__main__":
    unittest.main()
