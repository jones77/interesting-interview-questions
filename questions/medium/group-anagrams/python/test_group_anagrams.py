import unittest

from group_anagrams import group_anagrams


def _normalize(groups):
    return sorted(sorted(group) for group in groups)


class TestGroupAnagrams(unittest.TestCase):
    def test_examples(self):
        tests = [
            (
                ["eat", "tea", "tan", "ate", "nat", "bat"],
                [["bat"], ["nat", "tan"], ["ate", "eat", "tea"]],
            ),
            ([""], [[""]]),
            (["a"], [["a"]]),
        ]
        for strs, want in tests:
            with self.subTest(strs=strs):
                got = group_anagrams(strs)
                self.assertEqual(_normalize(got), _normalize(want))


if __name__ == "__main__":
    unittest.main()
