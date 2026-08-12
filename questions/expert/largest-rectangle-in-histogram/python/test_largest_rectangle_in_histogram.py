import unittest

from largest_rectangle_in_histogram import largest_rectangle_area


class TestLargestRectangleArea(unittest.TestCase):
    def test_largest_rectangle_area(self):
        tests = [
            ([2, 1, 5, 6, 2, 3], 10),
            ([2, 4], 4),
            ([1], 1),
            ([], 0),
        ]
        for heights, want in tests:
            with self.subTest(heights=heights):
                got = largest_rectangle_area(heights)
                self.assertEqual(got, want)


if __name__ == "__main__":
    unittest.main()
