import os
import tempfile
import unittest
import uuid
from pathlib import Path
from typing import Type


import olca_schema as lca
import olca_schema.zipio as zio


class ZipTest(unittest.TestCase):
    def test_read(self):
        (specs, path) = _setup()
        with zio.ZipReader(path) as reader:
            for spec in specs:
                (type_, uid) = spec
                entity = reader.read(type_, uid)
                if entity is None:
                    self.fail(f"{type_}::{uid} is None")
                self.assertEqual(type_, type(entity))
                self.assertEqual(uid, entity.id)
                self.assertEqual(_name_of(type_), entity.name)
                self.assertEqual("some/test/data", entity.category)
                self.assertEqual(["test", "data"], entity.tags)
        os.unlink(path)

    def test_read_ids(self):
        (specs, path) = _setup()
        with zio.ZipReader(path) as reader:
            for (type_, uid) in specs:
                ids = reader.ids_of(type_)
                self.assertTrue(uid in ids)
        os.unlink(path)

    def test_read_each(self):
        (specs, path) = _setup()
        with zio.ZipReader(path) as reader:
            for (type_, uid) in specs:
                instance = None
                for e in reader.read_each(type_):
                    if e.id == uid:
                        instance = e
                self.assertIsNotNone(instance)
        os.unlink(path)


def _name_of(root_type: Type[lca.RootEntity]) -> str:
    return f"an instance of {root_type.__name__}"


Spec = tuple[Type[lca.RootEntity], str]


def _setup() -> tuple[list[Spec], Path]:
    uid = lambda: str(uuid.uuid4())
    specs = [
        (lca.Actor, uid()),
        (lca.Currency, uid()),
        (lca.DQSystem, uid()),
        (lca.Epd, uid()),
        (lca.Flow, uid()),
        (lca.FlowProperty, uid()),
        (lca.ImpactCategory, uid()),
        (lca.ImpactMethod, uid()),
        (lca.Location, uid()),
        (lca.Parameter, uid()),
        (lca.Process, uid()),
        (lca.ProductSystem, uid()),
        (lca.Project, uid()),
        (lca.Result, uid()),
        (lca.SocialIndicator, uid()),
        (lca.Source, uid()),
        (lca.UnitGroup, uid()),
    ]

    path = Path(tempfile.mktemp(".zip", "olca"))
    with zio.ZipWriter(path) as writer:
        for spec in specs:
            _write(spec, writer)
    return (specs, path)


def _write(spec: Spec, writer: zio.ZipWriter):
    (root_type, uid) = spec
    entity = root_type()
    entity.id = uid
    entity.name = _name_of(root_type)
    entity.category = "some/test/data"
    entity.tags = ["test", "data"]
    writer.write(entity)
