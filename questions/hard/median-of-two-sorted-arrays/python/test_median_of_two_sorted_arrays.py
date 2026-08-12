import unittest

from median_of_two_sorted_arrays import find_median_sorted_arrays


class TestFindMedianSortedArrays(unittest.TestCase):
    def test_examples(self):
        tests = [
            ([1, 3], [2], 2.0),
            ([1, 2], [3, 4], 2.5),
            ([0, 0], [0, 0], 0.0),
            ([], [1], 1.0),
            ([2], [], 2.0),
        ]
        for nums1, nums2, want in tests:
            with self.subTest(nums1=nums1, nums2=nums2):
                self.assertEqual(find_median_sorted_arrays(nums1, nums2), want)


if __name__ == "__main__":
    unittest.main()
