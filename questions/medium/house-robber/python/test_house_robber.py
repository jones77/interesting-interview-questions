import unittest

from house_robber import rob


class TestRob(unittest.TestCase):
    def test_examples(self):
        tests = [
            ([1, 2, 3, 1], 4),
            ([2, 7, 9, 3, 1], 12),
            ([2, 1], 2),
            ([1], 1),
            ([], 0),
        ]
        for nums, want in tests:
            with self.subTest(nums=nums):
                got = rob(nums)
                self.assertEqual(got, want)


if __name__ == "__main__":
    unittest.main()
