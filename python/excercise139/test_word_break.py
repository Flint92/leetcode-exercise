import unittest

from excercise139.word_break import word_break


class TestWorkBreak(unittest.TestCase):
    def test_add(self):
        self.assertEqual(word_break("leetcode", ["leet", "code"]), True)
        self.assertEqual(word_break("applepenapple", ["apple", "pen"]), True)
        self.assertEqual(word_break("catsandog", ["cats", "dog", "sand", "and", "cat"]), False)

if __name__ == '__main__':
    unittest.main()
