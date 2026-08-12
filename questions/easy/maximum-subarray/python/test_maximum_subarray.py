import unittest

from maximum_subarray import max_sub_array


class TestMaxSubArray(unittest.TestCase):
    def test_max_sub_array(self):
        tests = [
            ([-2, 1, -3, 4, -1, 2, 1, -5, 4], 6),
            ([1], 1),
            ([5, 4, -1, 7, 8], 23),
            ([-1], -1),
        ]
        for nums, want in tests:
            with self.subTest(nums=nums):
                self.assertEqual(max_sub_array(nums), want)


if __name__ == "__main__":
    unittest.main()
