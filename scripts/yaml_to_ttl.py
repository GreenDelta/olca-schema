
from io import StringIO
import model


def main():
    buffer = StringIO()
    buffer.write(get_header())
    m = model.Model.load_yaml('../yaml')
    classes = []
    for t in m.types:
        if type(t) == model.EnumType:
            write_enum(t, buffer)
        if type(t) == model.ClassType:
            classes.append(t)
    buffer.write('### Class types ###\n\n')
    write_classes(classes, buffer)
    print(buffer.getvalue())


def write_enum(t, buffer):
    buffer.write(':%s a rdfs:Class;\n' % t.name)
    buffer.write('  rdfs:subClassOf :Enumeration;\n')
    buffer.write('  rdfs:comment "%s"\n' % t.doc)
    buffer.write('  .\n\n')
    for item in t.items:
        buffer.write(':%s a rdfs:Class;\n' % item.name)
        buffer.write('  rdfs:subClassOf :%s;\n' % t.name)
        buffer.write('  .\n\n')


def write_classes(classes, buffer):
    for t in classes:
        buffer.write(':%s a rdfs:Class;\n' % t.name)
        if not t.super_class is None:
            buffer.write('  rdfs:subClassOf :%s;\n' % t.super_class)
        buffer.write('  rdfs:comment "%s"\n' % t.doc)
        buffer.write('  .\n\n')


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