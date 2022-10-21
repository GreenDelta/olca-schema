import unittest
import olca_schema.results as res


class ResultTypesTest(unittest.TestCase):

    def test_envi_flow(self):
        envi_flow = res.EnviFlow.from_dict({
            'flow': {
                '@id': 'flow-id',
                'name': 'flow-name',
                'refUnit': 'flow-unit'
            },
            'location': {
                '@id': 'location-id',
                'name': 'location-name',
            },
            'isInput': True,
        })
        value = res.EnviFlowValue.from_dict({
            'enviFlow': envi_flow.to_dict(),
            'amount': 42,
        })

        self.assertEqual(envi_flow.flow.id, 'flow-id')
        self.assertEqual(envi_flow.flow.name, 'flow-name')
        self.assertEqual(envi_flow.flow.ref_unit, 'flow-unit')
        self.assertEqual(envi_flow.location.id, 'location-id')
        self.assertEqual(envi_flow.location.name, 'location-name')
        self.assertTrue(envi_flow.is_input)
        self.assertEqual(value.amount, 42)
        self.assertEqual(value.envi_flow.flow.id, 'flow-id')
