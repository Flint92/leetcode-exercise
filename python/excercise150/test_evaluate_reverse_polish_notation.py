import unittest

from excercise150.evaluate_reverse_polish_notation import evaluate_reverse_polish_notation


class TestEvaluateReversePolishNotation(unittest.TestCase):
    def test_evaluate_reverse_polish_notation(self):
        self.assertEqual(evaluate_reverse_polish_notation(["2", "1", "+", "3", "*"]), 9)
        self.assertEqual(evaluate_reverse_polish_notation(["4", "13", "5", "/", "+"]), 6)

if __name__ == '__main__':
    unittest.main()
