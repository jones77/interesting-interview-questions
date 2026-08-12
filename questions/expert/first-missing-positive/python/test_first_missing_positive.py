import unittest

from first_missing_positive import first_missing_positive


class TestFirstMissingPositive(unittest.TestCase):
    def test_first_missing_positive(self):
        tests = [
            ([1, 2, 0], 3),
            ([3, 4, -1, 1], 2),
            ([7, 8, 9, 11, 12], 1),
            ([1], 2),
            ([], 1),
        ]
        for nums, want in tests:
            with self.subTest(nums=nums):
                got = first_missing_positive(nums)
                self.assertEqual(got, want)


if __name__ == "__main__":
    unittest.main()
