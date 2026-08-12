import unittest

from burst_balloons import max_coins


class TestMaxCoins(unittest.TestCase):
    def test_examples(self):
        tests = [
            ([3, 1, 5, 8], 167),
            ([1, 5], 10),
            ([], 0),
        ]
        for nums, want in tests:
            with self.subTest(nums=nums):
                self.assertEqual(max_coins(nums), want)


if __name__ == "__main__":
    unittest.main()
