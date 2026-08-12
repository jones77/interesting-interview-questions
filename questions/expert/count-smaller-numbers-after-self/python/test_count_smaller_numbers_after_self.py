import unittest

from count_smaller_numbers_after_self import count_smaller


class TestCountSmaller(unittest.TestCase):
    def test_count_smaller(self):
        tests = [
            ([5, 2, 6, 1], [2, 1, 1, 0]),
            ([-1, -1], [0, 0]),
            ([], []),
        ]
        for nums, want in tests:
            with self.subTest(nums=nums):
                got = count_smaller(nums)
                self.assertEqual(got, want)


if __name__ == "__main__":
    unittest.main()
