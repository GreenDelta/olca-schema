import unittest

import olca_schema as lca


class TestBools(unittest.TestCase):

    def test_false(self):
        p = lca.new_parameter('param', '1 + 1')
        self.assertIsNotNone(p.is_input_parameter)
        self.assertFalse(p.is_input_parameter)
        p = lca.Parameter.from_json(p.to_json())
        self.assertIsNotNone(p.is_input_parameter)
        self.assertFalse(p.is_input_parameter)

    def test_true(self):
        p = lca.new_parameter('param', 42)
        self.assertTrue(p.is_input_parameter)
        p = lca.Parameter.from_json(p.to_json())
        self.assertTrue(p.is_input_parameter)


if __name__ == '__main__':
    unittest.main()
