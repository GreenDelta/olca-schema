"""
Generates the HTML documentation from the YAML model.
"""

import model
from jinja2 import Environment, FileSystemLoader

env = Environment(loader=FileSystemLoader('./templates'))

template = env.get_template('class_template.html')
m = model.Model.load_yaml('../yaml')

for clazz in m.types:
    super_classes = m.get_super_classes(clazz)
    text = template.render(model=clazz, super_classes=super_classes)
    file_path = '../html/%s.html' % clazz.name
    with open(file_path, 'w') as f:
        f.write(text)
