import unittest

from excercise413.arithmetic_slices import number_of_arithmetic_slices

class TestArithmeticSlices(unittest.TestCase):
    def test_arithmetic_slices(self):
        self.assertEqual(number_of_arithmetic_slices([1,2,3,4]), 3)
        self.assertEqual(number_of_arithmetic_slices([1]), 0)

if __name__ == '__main__':
    unittest.main()
