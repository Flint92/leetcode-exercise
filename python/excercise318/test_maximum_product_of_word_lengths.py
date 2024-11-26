import unittest
from excercise318.maximum_product_of_word_lengths import maximum_product_of_word_lengths

class TestMaximumProductOfWordLengths(unittest.TestCase):
    def test_maximum_product_of_word_lengths(self):
        self.assertEqual(maximum_product_of_word_lengths(["abcw","baz","foo","bar","xtfn","abcdef"]), 16)
        self.assertEqual(maximum_product_of_word_lengths(["a","ab","abc","d","cd","bcd","abcd"]), 4)
        self.assertEqual(maximum_product_of_word_lengths(["a","aa","aaa","aaaa"]), 0)

if __name__ == '__main__':
    unittest.main()
