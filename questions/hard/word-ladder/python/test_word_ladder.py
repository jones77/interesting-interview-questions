import unittest

from word_ladder import ladder_length


class TestLadderLength(unittest.TestCase):
    def test_examples(self):
        tests = [
            ("hit", "cog", ["hot", "dot", "dog", "lot", "log", "cog"], 5),
            ("hit", "cog", ["hot", "dot", "dog", "lot", "log"], 0),
            ("a", "c", ["a", "b", "c"], 2),
        ]
        for begin_word, end_word, word_list, want in tests:
            with self.subTest(begin_word=begin_word, end_word=end_word):
                self.assertEqual(
                    ladder_length(begin_word, end_word, word_list), want
                )


if __name__ == "__main__":
    unittest.main()
