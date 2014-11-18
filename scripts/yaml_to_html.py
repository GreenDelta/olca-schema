"""
Generates the HTML documentation from the YAML model.
"""

import model
from jinja2 import Environment, FileSystemLoader


def main():
    env = Environment(loader=FileSystemLoader('./templates'))
    index_template = env.get_template('index_template.html')
    class_template = env.get_template('class_template.html')
    enum_template = env.get_template('enum_template.html')
    m = model.Model.load_yaml('../yaml')
    write_index(m, index_template)
    for t in m.types:
        if type(t) == model.ClassType:
            write_class(class_template, m, t)
        if type(t) == model.EnumType:
            write_enum(enum_template, t)


def write_enum(template, t):
    text = template.render(model=t)
    file_path = '../html/%s.html' % t.name
    with open(file_path, 'w') as f:
        f.write(text)


def write_class(template, m, t):
    super_classes = m.get_super_classes(t)
    example = get_example(t)
    text = template.render(model=t, super_classes=super_classes,
                           example=example)
    file_path = '../html/%s.html' % t.name
    with open(file_path, 'w') as f:
        f.write(text)


def get_example(t):
    if t.example is None:
        return None
    path = '../examples/' + t.example
    with open(path, 'r') as f:
        return f.read()


def write_index(m, template):
    concepts = []
    for t in m.types:
        concepts.append(t.name)
    text = template.render(concepts=concepts)
    with open('../index.html', 'w') as f:
        f.write(text)


if __name__ == '__main__':
    main()