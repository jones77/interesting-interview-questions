import unittest

from frog_jump import can_cross


class TestCanCross(unittest.TestCase):
    def test_can_cross(self):
        tests = [
            ([0, 1, 3, 5, 6, 8, 12, 17], True),
            ([0, 1, 2, 3, 4, 8, 9, 11], False),
            ([0, 1], True),
        ]
        for stones, want in tests:
            with self.subTest(stones=stones):
                got = can_cross(stones)
                self.assertEqual(got, want)


if __name__ == "__main__":
    unittest.main()
