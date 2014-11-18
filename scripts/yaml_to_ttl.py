
from io import StringIO
import model


def main():
    buffer = StringIO()
    buffer.write(get_header())
    m = model.Model.load_yaml('../yaml')
    classes = []
    properties = {}
    for t in m.types:
        if type(t) == model.EnumType:
            write_enum(t, buffer)
        if type(t) == model.ClassType:
            classes.append(t)
            collect_properties(t, properties)
    buffer.write('### Class types ###\n\n')
    write_classes(classes, buffer)
    write_properties(properties, buffer)
    with open('../generated/schema.ttl', 'w') as f:
        f.write(buffer.getvalue())


def write_enum(t, buffer):
    buffer.write(':%s a rdfs:Class;\n' % t.name)
    buffer.write('  rdfs:subClassOf :Enumeration;\n')
    buffer.write('  rdfs:comment "%s"\n' % t.doc)
    buffer.write('  .\n\n')
    for item in t.items:
        buffer.write(':%s a rdfs:Class;\n' % item.name)
        buffer.write('  rdfs:subClassOf :%s;\n' % t.name)
        buffer.write('  .\n\n')


def collect_properties(t, properties):
    for prop in t.properties:
        if not prop.name in properties:
            properties[prop.name] = (prop, [t])
        else:
            entry = properties[prop.name]
            entry[1].append(t)


def write_classes(classes, buffer):
    for t in classes:
        buffer.write(':%s a rdfs:Class;\n' % t.name)
        if not t.super_class is None:
            buffer.write('  rdfs:subClassOf :%s;\n' % t.super_class)
        buffer.write('  rdfs:comment "%s"\n' % t.doc)
        buffer.write('  .\n\n')


def write_properties(properties, buffer):
    for name in properties:
        entry = properties[name]
        domain = get_domain(entry[1])
        prop = entry[0]
        rdf_range = get_range(prop)
        buffer.write(':%s a rdf:Property;\n' % name)
        buffer.write('  rdfs:domain %s;\n' % domain)
        buffer.write('  rdfs:range %s;\n' % rdf_range)
        if len(entry[1]) == 1 and not prop.doc is None:
            buffer.write('  rdfs:comment "%s"\n' % prop.doc)
        buffer.write('  .\n\n')


def get_domain(class_list):
    if len(class_list) == 1:
        return ':' + class_list[0].name
    text = '[ a owl:Class; owl:unionOf ('
    i = 0
    for c in class_list:
        i += 1
        text += ':' + c.name
        if i < len(class_list):
            text += ' '
    return text + ')]'


def get_range(prop):
    t = prop.field_type
    if t.startswith('List['):
        return 'rdf:List'
    if t[0].isupper():
        return ':' + t
    else:
        return 'xsd:' + t


def get_header():
    return """
@prefix : <http://openlca.org/> .
@prefix owl: <http://www.w3.org/2002/07/owl#> .
@prefix rdfs: <http://www.w3.org/2000/01/rdf-schema#> .
@prefix rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#> .
@prefix xsd: <http://www.w3.org/2001/XMLSchema#> .

### Enumeration types ###

:Enumeration a rdfs:Class;
  rdfs:comment "The super-class of all enumeration types."
  .

"""


if __name__ == '__main__':
    main()