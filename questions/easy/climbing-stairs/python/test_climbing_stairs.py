import unittest

from climbing_stairs import climb_stairs


class TestClimbStairs(unittest.TestCase):
    def test_climb_stairs(self):
        tests = [
            (1, 1),
            (2, 2),
            (3, 3),
            (4, 5),
            (5, 8),
        ]
        for n, want in tests:
            with self.subTest(n=n):
                self.assertEqual(climb_stairs(n), want)


if __name__ == "__main__":
    unittest.main()
