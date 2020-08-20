import glob
import re

import yaml


class Model:
    def __init__(self):
        self.types = []

    @staticmethod
    def load_yaml(folder):
        """
        Loads a model from a directory which contains YAML files that describe
        the model.
        :param folder: The directory that contains the YAML files (extension
                       *.yaml)
        """
        m = Model()
        d = folder if folder.endswith(
            '/') or folder.endswith('\\') else folder + '/'
        files = glob.glob(d + "*.yaml")
        for file_path in files:
            with open(file_path, 'r') as f:
                yaml_model = yaml.load(f)
                if 'class' in yaml_model:
                    m.types.append(ClassType.load_yaml(yaml_model['class']))
                if 'enum' in yaml_model:
                    m.types.append(EnumType.load_yaml(yaml_model['enum']))
        return m

    def find_type(self, name: str):
        if name is None:
            return None
        if name.startswith('Ref['):
            return self.find_type('Ref')
        for t in self.types:
            if t.name == name:
                return t
        return None

    def get_super_classes(self, clazz):
        classes = []
        c = self.find_type(clazz.super_class)
        while c is not None:
            classes.append(c)
            c = self.find_type(c.super_class)
        return classes


class ClassType:
    def __init__(self, name=None, super_class=None, doc=None):
        self.name = name
        self.super_class = super_class
        self.doc = format_doc(doc)
        self.example = None
        self.properties = []

    @staticmethod
    def load_yaml(yaml_model):
        c = ClassType()
        c.name = yaml_model['name']
        if 'doc' in yaml_model:
            c.doc = format_doc(yaml_model['doc'])
        else:
            c.doc = ''
        if 'superClass' in yaml_model:
            c.super_class = yaml_model['superClass']
        if 'properties' in yaml_model:
            for prop in yaml_model['properties']:
                c.properties.append(Property.load_yaml(prop))
        if 'example' in yaml_model:
            c.example = yaml_model['example']
        return c


class Property:
    def __init__(self, name=None, field_type=None, doc=None):
        self.name = name
        self.field_type = field_type
        self.doc = format_doc(doc)

    @staticmethod
    def load_yaml(yaml_model):
        p = Property()
        p.name = yaml_model['name']
        p.field_type = yaml_model['type']
        if 'doc' in yaml_model:
            p.doc = format_doc(yaml_model['doc'])
        else:
            p.doc = ''
        return p

    @property
    def html_type_link(self):
        t = self.field_type  # type: str
        if t == 'GeoJSON':
            return 'https://tools.ietf.org/html/rfc7946'
        if t.startswith('List[Ref['):
            t = 'Ref'
        elif t.startswith('List['):
            end = len(t) - 1
            t = t[5:end]
        elif t.startswith('Ref['):
            t = 'Ref'
        if t[0].isupper():
            return "./%s.html" % t
        else:
            return "http://www.w3.org/TR/xmlschema-2/#%s" % t


class EnumType:
    def __init__(self, name=None, doc=None):
        self.name = name
        self.doc = format_doc(doc)
        self.items = []

    @staticmethod
    def load_yaml(yaml_model):
        e = EnumType()
        e.name = yaml_model['name']
        if 'doc' in yaml_model:
            e.doc = format_doc(yaml_model['doc'])
        else:
            e.doc = ''
        if 'items' in yaml_model:
            for item in yaml_model['items']:
                elem = EnumItem(item['name'])
                if 'doc' in item:
                    elem.doc = format_doc(item['doc'])
                e.items.append(elem)
        return e


class EnumItem:
    def __init__(self, name=None):
        self.name = name
        self.doc = ''


def format_doc(doc: str) -> str:
    if doc is None:
        return ''
    for match in re.findall('\\[[^\\]]*\\]', doc):
        link_type = match[1:(len(match)-1)]
        link = '<a href="./%s.html">%s</a>' % (link_type, link_type)
        doc = doc.replace(match, link)
    for match in re.findall('`[^`]+`', doc):
        code = match[1:(len(match)-1)]
        span = '<span style="color: #7c4dff">%s</span>' % (code)
        doc = doc.replace(match, span)
    return doc
