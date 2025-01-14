import unittest

from excercise8.string_to_integer_atoi import string_to_integer_atoi

class TestStringToIntegerAtoi(unittest.TestCase):
    def test_string_to_integer_atoi(self):
        self.assertEqual(string_to_integer_atoi('42'), 42)
        self.assertEqual(string_to_integer_atoi('  -42 '), -42)
        self.assertEqual(string_to_integer_atoi('1337c0d3'), 1337)
        self.assertEqual(string_to_integer_atoi('0-1'), 0)
        self.assertEqual(string_to_integer_atoi('words and 987'), 0)

if __name__ == '__main__':
    unittest.main()
