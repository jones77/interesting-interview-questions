import unittest

from word_break_ii import word_break


class TestWordBreak(unittest.TestCase):
    def test_word_break(self):
        tests = [
            (
                "catsanddog",
                ["cat", "cats", "and", "sand", "dog"],
                ["cats and dog", "cat sand dog"],
            ),
            (
                "pineapplepenapple",
                ["apple", "pen", "applepen", "pine", "pineapple"],
                ["pine apple pen apple", "pineapple pen apple", "pine applepen apple"],
            ),
            (
                "catsandog",
                ["cats", "dog", "sand", "and", "cat"],
                None,
            ),
        ]
        for s, word_dict, want in tests:
            with self.subTest(s=s, word_dict=word_dict):
                got = word_break(s, word_dict)
                self.assertEqual(got, want)


if __name__ == "__main__":
    unittest.main()
