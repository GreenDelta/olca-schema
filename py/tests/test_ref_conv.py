import olca_schema as lca
import unittest


class RefConvTest(unittest.TestCase):

    def test_ref(self):
        prop = lca.FlowProperty(unit_group=lca.Ref(id='abc'))
        prop = lca.FlowProperty.from_json(prop.to_json())
        self.assertEqual('abc', prop.unit_group.id)
