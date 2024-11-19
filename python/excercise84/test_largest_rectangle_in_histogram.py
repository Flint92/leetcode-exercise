import unittest

from excercise84.largest_rectangle_in_histogram import largest_rectangle_in_histogram


class TestLargestRectangleInHistogram(unittest.TestCase):
    def test_add(self):
        self.assertEqual(largest_rectangle_in_histogram([2,1,5,6,2,3]), 10)

if __name__ == '__main__':
    unittest.main()
