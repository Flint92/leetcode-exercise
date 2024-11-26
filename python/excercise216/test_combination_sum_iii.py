import unittest

from excercise216.combination_sum_iii import combination_sum3


class TestCombinationSum3(unittest.TestCase):
    def test_combination_sum3(self):
        self.assertEqual(combination_sum3(3, 7), [[1, 2, 4]])
        self.assertEqual(combination_sum3(3, 9), [[1, 2, 6], [1, 3, 5], [2, 3, 4]])
        self.assertEqual(combination_sum3(4, 1), [])


if __name__ == '__main__':
    unittest.main()
