import os
import tempfile
import unittest
import uuid

import olca_schema as lca
import olca_schema.zipio as zio

from typing import Type


class ZipTest(unittest.TestCase):

    def test_zip_io(self):
        specs: list[tuple[Type[lca.RootEntity], str]] = [
            (lca.Actor, str(uuid.uuid4())),
            (lca.Currency, str(uuid.uuid4())),
            (lca.DQSystem, str(uuid.uuid4())),
            (lca.Epd, str(uuid.uuid4())),
            (lca.Flow, str(uuid.uuid4())),
            (lca.FlowProperty, str(uuid.uuid4())),
            (lca.ImpactCategory, str(uuid.uuid4())),
            (lca.ImpactMethod, str(uuid.uuid4())),
            (lca.Location, str(uuid.uuid4())),
            (lca.Parameter, str(uuid.uuid4())),
            (lca.Process, str(uuid.uuid4())),
            (lca.ProductSystem, str(uuid.uuid4())),
            (lca.Project, str(uuid.uuid4())),
            (lca.Result, str(uuid.uuid4())),
            (lca.SocialIndicator, str(uuid.uuid4())),
            (lca.Source, str(uuid.uuid4())),
            (lca.UnitGroup, str(uuid.uuid4())),
        ]

        f = tempfile.mktemp('.zip', 'olca')
        with zio.ZipWriter(f) as writer:
            for spec in specs:
                _write(spec, writer)

        with zio.ZipReader(f) as reader:
            for spec in specs:
                (root_type, uid) = spec
                entity = reader.read(root_type, uid)
                self.assertEqual(root_type, type(entity))
                self.assertEqual(uid, entity.id)
                self.assertEqual(_name_of(root_type), entity.name)
                self.assertEqual('some/test/data', entity.category)
                self.assertEqual(['test', 'data'], entity.tags)

        os.unlink(f)


def _write(spec: tuple[Type[lca.RootEntity], str], writer: zio.ZipWriter):
    (root_type, uid) = spec
    entity = root_type()
    entity.id = uid
    entity.name = _name_of(root_type)
    entity.category = 'some/test/data'
    entity.tags = ['test', 'data']
    writer.write(entity)


def _name_of(root_type: Type[lca.RootEntity]) -> str:
    return f'an instance of {root_type.__name__}'
