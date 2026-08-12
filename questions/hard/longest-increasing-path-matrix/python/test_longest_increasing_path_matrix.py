import unittest

from longest_increasing_path_matrix import longest_increasing_path


class TestLongestIncreasingPath(unittest.TestCase):
    def test_examples(self):
        tests = [
            ([[9, 9, 4], [6, 6, 8], [2, 1, 1]], 4),
            ([[3, 4, 5], [3, 2, 6], [2, 2, 1]], 4),
            ([[1]], 1),
        ]
        for matrix, want in tests:
            with self.subTest(matrix=matrix):
                self.assertEqual(longest_increasing_path(matrix), want)


if __name__ == "__main__":
    unittest.main()
