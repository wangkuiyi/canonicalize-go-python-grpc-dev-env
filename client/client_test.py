import unittest
import client


class SQLFlowGRPCTest(unittest.TestCase):
    def testRun(self):
        self.assertEqual(client.SQLFlowRun(), 100)


if __name__ == '__main__':
    unittest.main()
