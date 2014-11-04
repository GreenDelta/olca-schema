"""
Generates the HTML documentation from the YAML model.
"""

import model
from jinja2 import Environment, FileSystemLoader


def main():
    env = Environment(loader=FileSystemLoader('./templates'))
    class_template = env.get_template('class_template.html')
    enum_template = env.get_template('enum_template.html')
    m = model.Model.load_yaml('../yaml')
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
    text = template.render(model=t, super_classes=super_classes)
    file_path = '../html/%s.html' % t.name
    with open(file_path, 'w') as f:
        f.write(text)

if __name__ == '__main__':
    main()