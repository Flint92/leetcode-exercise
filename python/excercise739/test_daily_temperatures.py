import unittest

from excercise739.daily_temperatures import daily_temperatures


class TestDailyTemperatures(unittest.TestCase):
    def test_daily_temperatures(self):
        self.assertEqual(daily_temperatures([73, 74, 75, 71, 69, 72, 76, 73]), [1, 1, 4, 2, 1, 1, 0, 0])
        self.assertEqual(daily_temperatures([30, 40, 50, 60]), [1, 1, 1, 0])
        self.assertEqual(daily_temperatures([30, 60, 90]), [1, 1, 0])

if __name__ == '__main__':
    unittest.main()
