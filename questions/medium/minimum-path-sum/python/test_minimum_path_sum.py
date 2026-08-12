import unittest

from minimum_path_sum import min_path_sum


class TestMinPathSum(unittest.TestCase):
    def test_examples(self):
        tests = [
            ([[1, 3, 1], [1, 5, 1], [4, 2, 1]], 7),
            ([[1, 2, 3], [4, 5, 6]], 12),
            ([[5]], 5),
        ]
        for grid, want in tests:
            with self.subTest(grid=grid):
                got = min_path_sum(grid)
                self.assertEqual(got, want)


if __name__ == "__main__":
    unittest.main()
