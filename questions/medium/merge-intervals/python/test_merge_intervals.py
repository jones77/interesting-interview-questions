import unittest

from merge_intervals import merge


class TestMerge(unittest.TestCase):
    def test_examples(self):
        tests = [
            ([[1, 3], [2, 6], [8, 10], [15, 18]], [[1, 6], [8, 10], [15, 18]]),
            ([[1, 4], [4, 5]], [[1, 5]]),
            ([[1, 4], [0, 4]], [[0, 4]]),
            ([[1, 4]], [[1, 4]]),
        ]
        for intervals, want in tests:
            with self.subTest(intervals=intervals):
                got = merge(intervals)
                self.assertEqual(got, want)


if __name__ == "__main__":
    unittest.main()
