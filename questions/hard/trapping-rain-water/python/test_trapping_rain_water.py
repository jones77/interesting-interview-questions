import unittest

from trapping_rain_water import trap


class TestTrap(unittest.TestCase):
    def test_examples(self):
        tests = [
            ([0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1], 6),
            ([4, 2, 0, 3, 2, 5], 9),
            ([], 0),
            ([1], 0),
        ]
        for height, want in tests:
            with self.subTest(height=height):
                self.assertEqual(trap(height), want)


if __name__ == "__main__":
    unittest.main()
