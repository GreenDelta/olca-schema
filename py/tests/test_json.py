import unittest

import olca_schema as o


class TestJson(unittest.TestCase):

    def test_ref_dict(self):
        ref = o.Ref(model_type='Flow', id='co2', name='CO2')
        ref_dict = ref.to_dict()
        self.assertEqual('Flow', ref_dict['@type'])
        self.assertEqual('co2', ref_dict['@id'])
        self.assertEqual('CO2', ref_dict['name'])

    def test_type_tags(self):
        instances = [
            o.Actor(),
            o.Source(),
            o.UnitGroup(),
            o.FlowProperty(),
            o.SocialIndicator(),
            o.Flow(),
            o.Process(),
            o.ImpactCategory(),
            o.ImpactMethod(),
            o.ProductSystem(),
            o.Project(),
            o.Result(),
            o.Epd()
        ]
        for i in instances:
            d = i.to_dict()
            self.assertEqual(type(i).__name__, d['@type'])

    def test_process(self):
        p1 = o.Process(
            name='a process',
            process_type=o.ProcessType.UNIT_PROCESS)
        p2 = o.Process.from_json(p1.to_json())
        self.assertEqual('a process', p2.name)
        self.assertEqual(o.ProcessType.UNIT_PROCESS, p2.process_type)

    def test_type_tag(self):
        flow = o.Flow()
        self.assertEqual('Flow', flow.to_ref().model_type)
        flow_dict = flow.to_dict()
        self.assertEqual('Flow', flow_dict['@type'])


if __name__ == '__main__':
    unittest.main()
