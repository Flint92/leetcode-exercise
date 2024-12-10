import unittest

from excercise62.unique_paths import unique_paths

class TestUniquePaths(unittest.TestCase):
    def test_unique_paths(self):
        self.assertEqual(unique_paths(3, 7), 28)
        self.assertEqual(unique_paths(3, 2), 3)

if __name__ == '__main__':
    unittest.main()
