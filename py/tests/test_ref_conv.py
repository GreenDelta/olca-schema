import olca_schema as lca
import unittest


class RefConvTest(unittest.TestCase):

    def test_ref(self):
        prop = lca.FlowProperty(unit_group=lca.Ref(id='abc'))
        prop = lca.FlowProperty.from_json(prop.to_json())
        self.assertEqual('abc', prop.unit_group.id)

    def test_refs(self):
        refs = [lca.Ref(id=str(i)) for i in range(0, 3)]
        method = lca.ImpactMethod(impact_categories=refs)
        method = lca.ImpactMethod.from_json(method.to_json())
        for i in range(0, 3):
            ref = method.impact_categories[i]
            self.assertEqual(str(i), ref.id)
