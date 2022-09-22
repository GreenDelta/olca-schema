import unittest

import olca_schema as lca


class TestJson(unittest.TestCase):

    def test_ref_dict(self):
        ref = lca.Ref(model_type='Flow', id='co2', name='CO2')
        ref_dict = ref.to_dict()
        self.assertEqual('Flow', ref_dict['@type'])
        self.assertEqual('co2', ref_dict['@id'])
        self.assertEqual('CO2', ref_dict['name'])

    def test_type_tags(self):
        instances = [
            lca.Actor(),
            lca.Source(),
            lca.UnitGroup(),
            lca.FlowProperty(),
            lca.SocialIndicator(),
            lca.Flow(),
            lca.Process(),
            lca.ImpactCategory(),
            lca.ImpactMethod(),
            lca.ProductSystem(),
            lca.Project(),
            lca.Result(),
            lca.Epd()
        ]
        for i in instances:
            d = i.to_dict()
            self.assertEqual(type(i).__name__, d['@type'])

    def test_process(self):
        p1 = lca.Process(
            name='a process',
            process_type=lca.ProcessType.UNIT_PROCESS)
        p2 = lca.Process.from_json(p1.to_json())
        self.assertEqual('a process', p2.name)
        self.assertEqual(lca.ProcessType.UNIT_PROCESS, p2.process_type)

    def test_type_tag(self):
        flow = lca.Flow()
        self.assertEqual('Flow', flow.to_ref().model_type)
        flow_dict = flow.to_dict()
        self.assertEqual('Flow', flow_dict['@type'])


if __name__ == '__main__':
    unittest.main()
