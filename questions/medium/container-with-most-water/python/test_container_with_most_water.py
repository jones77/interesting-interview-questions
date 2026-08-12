import unittest

from container_with_most_water import max_area


class TestMaxArea(unittest.TestCase):
    def test_examples(self):
        tests = [
            ([1, 8, 6, 2, 5, 4, 8, 3, 7], 49),
            ([1, 1], 1),
            ([4, 3, 2, 1, 4], 16),
        ]
        for height, want in tests:
            with self.subTest(height=height):
                got = max_area(height)
                self.assertEqual(got, want)


if __name__ == "__main__":
    unittest.main()
