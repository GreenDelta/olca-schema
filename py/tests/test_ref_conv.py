import olca_schema as o
import unittest


class RefConvTest(unittest.TestCase):

    def test_ref(self):
        prop = o.FlowProperty(unit_group=o.Ref(id='abc'))
        prop = o.FlowProperty.from_json(prop.to_json())
        self.assertEqual('abc', prop.unit_group.id)

    def test_refs(self):
        refs = [o.Ref(id=str(i)) for i in range(0, 3)]
        method = o.ImpactMethod(impact_categories=refs)
        method = o.ImpactMethod.from_json(method.to_json())
        for i in range(0, 3):
            ref = method.impact_categories[i]
            self.assertEqual(str(i), ref.id)
