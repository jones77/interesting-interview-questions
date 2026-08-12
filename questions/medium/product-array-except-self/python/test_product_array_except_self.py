import unittest

from product_array_except_self import product_except_self


class TestProductExceptSelf(unittest.TestCase):
    def test_examples(self):
        tests = [
            ([1, 2, 3, 4], [24, 12, 8, 6]),
            ([-1, 1, 0, -3, 3], [0, 0, 9, 0, 0]),
        ]
        for nums, want in tests:
            with self.subTest(nums=nums):
                got = product_except_self(nums)
                self.assertEqual(got, want)


if __name__ == "__main__":
    unittest.main()
