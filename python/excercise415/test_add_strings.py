import unittest

from excercise415.add_strings import add_strings

class AddStrings(unittest.TestCase):
    def test_add_strings(self):
        self.assertEqual(add_strings('11', '123'), '134')
        self.assertEqual(add_strings('456', '77'), '533')
        self.assertEqual(add_strings('0', '0'), '0')

if __name__ == '__main__':
    unittest.main()
