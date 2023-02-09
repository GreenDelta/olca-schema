import uuid

import olca_schema as o
import unittest


class RootTypeTest(unittest.TestCase):

    def test_conv(self):
        types = [
            o.Actor,
            o.Currency,
            o.DQSystem,
            o.Epd,
            o.Flow,
            o.FlowProperty,
            o.ImpactCategory,
            o.ImpactMethod,
            o.Location,
            o.Parameter,
            o.Process,
            o.ProductSystem,
            o.Project,
            o.Result,
            o.SocialIndicator,
            o.Source,
            o.UnitGroup,
        ]

        for t in types:
            entity = t()
            entity.id = str(uuid.uuid4())
            entity.name = t.__name__

            json_copy = t.from_json(entity.to_json())
            self.assertEqual(entity.id, json_copy.id)
            self.assertEqual(entity.name, json_copy.name)

            dict_copy = t.from_dict(entity.to_dict())
            self.assertEqual(entity.id, dict_copy.id)
            self.assertEqual(entity.name, dict_copy.name)

            ref = entity.to_ref()
            self.assertEqual(entity.id, ref.id)
            self.assertEqual(entity.name, ref.name)
