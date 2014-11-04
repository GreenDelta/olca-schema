"""

"""

import rdflib
from rdflib.namespace import RDF
from rdflib.namespace import RDFS

OLCA = "http://openlca.org/"


def main():
    g = rdflib.Graph()
    g.load('../schema.ttl', format='n3')

    for c in get_classes(g):
        print("%s" % c)
        for p in get_class_properties(g, c):
            print("\t%s" % p)

    for t in g:
        print(t)

def get_classes(graph):
    classes = []
    for s, p, o in graph.triples((None, RDF.type, RDFS.Class)):
        classes.append(s)
    return classes


def get_properties(graph):
    props = []
    for s, p, o in graph.triples((None, RDF.type, RDF.Property)):
        props.append(s)
    return props


def get_class_properties(graph, clazz):
    props = []
    for s, p, o in graph.triples((None, RDFS.domain, clazz)):
        props.append(s)
    return props


if __name__ == '__main__':
    main()
