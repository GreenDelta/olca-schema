import glob
import yaml


class Model:
    def __init__(self):
        self.types = []

    @staticmethod
    def load_yaml(folder):
        """
        Loads a model from a directory which contains YAML files that describe
        the model.
        :param dir: The directory that contains the YAML files (extension *.yaml)
        """
        m = Model()
        d = folder if folder.endswith('/') or folder.endswith('\\') else folder + '/'
        files = glob.glob(d + "*.yaml")
        for file_path in files:
            with open(file_path, 'r') as f:
                yaml_model = yaml.load(f)
                if 'class' in yaml_model:
                    m.types.append(ClassType.load_yaml(yaml_model['class']))
        return m


class ClassType:
    def __init__(self, name=None, super_class=None, doc=None):
        self.name = name
        self.super_class = super_class
        self.doc = doc
        self.properties = []

    @staticmethod
    def load_yaml(yaml_model):
        c = ClassType()
        c.name = yaml_model['name']
        if 'doc' in yaml_model.keys():
            c.doc = yaml_model['doc']
        else:
            c.doc = ''
        if 'superClass' in yaml_model.keys():
            c.super_class = yaml_model['superClass']
        if 'properties' in yaml_model.keys():
            for prop in yaml_model['properties']:
                c.properties.append(Property.load_yaml(prop))
        return c


class Property:
    def __init__(self, name=None, field_type=None, doc=None):
        self.name = name
        self.field_type = field_type
        self.doc = doc

    @staticmethod
    def load_yaml(yaml_model):
        p = Property()
        p.name = yaml_model['name']
        p.field_type = yaml_model['type']
        if 'doc' in yaml_model:
            p.doc = yaml_model['doc']
        return p


class EnumType:
    def __init__(self, name, doc):
        self.name = name
        self.doc = doc
        self.items = []


class EnumItem:
    def __init__(self, name):
        self.name = name