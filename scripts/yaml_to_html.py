"""
Generates the HTML documentation from the YAML model.
"""

import model
from jinja2 import Environment, FileSystemLoader

env = Environment(loader=FileSystemLoader('./templates'))

template = env.get_template('class_template.html')
m = model.Model.load_yaml('../yaml')

for clazz in m.types:
    print(template.render(model=clazz))
