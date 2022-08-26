import zipfile

import olca_schema as schema

from typing import Optional, Type


class ZipWriter:

    def __init__(self, file_name: str):
        self.__zip = zipfile.ZipFile(
            file_name, mode='a', compression=zipfile.ZIP_DEFLATED)
        if 'olca-schema.json' not in self.__zip.namelist():
            self.__zip.writestr('olca-schema.json', '{"version": 2}')

    def __enter__(self):
        return self

    def __exit__(self, type, value, traceback):
        self.close()

    def close(self):
        self.__zip.close()

    def write(self, entity: schema.RootEntity):
        if entity.id is None or entity.id == '':
            raise ValueError('entity must have an ID')
        folder = _folder_of_entity(entity)
        path = f'{folder}/{entity.id}.json'
        self.__zip.writestr(path, entity.to_json())


class ZipReader:

    def __init__(self, file_name: str):
        self.__zip = zipfile.ZipFile(file_name, mode='r')

    def __enter__(self):
        return self

    def __exit__(self, type, value, traceback):
        self.close()

    def close(self):
        self.__zip.close()

    def read(self, class_type: Type[schema.RootEntity],
             uid: str) -> Optional[schema.RootEntity]:
        folder = _folder_of_class(class_type)
        path = f'{folder}/{uid}.json'
        if path not in self.__zip.namelist():
            return None
        data = self.__zip.read(path)
        return class_type.from_json(data)

    def read_actor(self, uid: str) -> Optional[schema.Actor]:
        return self.read(schema.Actor, uid)


def _folder_of_entity(entity: schema.RootEntity) -> str:
    if entity is None:
        raise ValueError("unknown root entity type")
    return _folder_of_class(type(entity))


def _folder_of_class(t: type) -> str:
    if t == schema.Actor:
        return 'actors'
    if t == schema.Currency:
        return 'currencies'
    if t == schema.DQSystem:
        return 'dq_systems'
    if t == schema.Epd:
        return 'epds'
    if t == schema.Flow:
        return 'flows'
    if t == schema.FlowProperty:
        return 'flow_properties'
    if t == schema.ImpactCategory:
        return 'lcia_categories'
    if t == schema.ImpactMethod:
        return 'lcia_methods'
    if t == schema.Location:
        return 'locations'
    if t == schema.Parameter:
        return 'parameters'
    if t == schema.Process:
        return 'processes'
    if t == schema.ProductSystem:
        return 'product_systems'
    if t == schema.Project:
        return 'projects'
    if t == schema.Result:
        return 'results'
    if t == schema.SocialIndicator:
        return 'social_indicators'
    if t == schema.Source:
        return 'sources'
    if t == schema.UnitGroup:
        return 'unit_groups'
    raise ValueError(f'not a known root entity type: {t}')
