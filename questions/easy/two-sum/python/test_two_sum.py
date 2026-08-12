import unittest

from two_sum import two_sum


class TestTwoSum(unittest.TestCase):
    def test_two_sum(self):
        tests = [
            ([2, 7, 11, 15], 9, [0, 1]),
            ([3, 2, 4], 6, [1, 2]),
            ([3, 3], 6, [0, 1]),
        ]
        for nums, target, want in tests:
            with self.subTest(nums=nums, target=target):
                got = two_sum(nums, target)
                self.assertEqual(sorted(got), sorted(want))


if __name__ == "__main__":
    unittest.main()
