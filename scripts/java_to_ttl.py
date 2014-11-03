"""
Converts Java classes to a RDF vocabulary. This script is only used for
initially generating the openLCA vocabulary from the openLCA model.
"""
import glob


def main():
    java_model = "C:/Users/Besitzer/Projects/git_repos/olca-modules/olca-core/src/main/java/org/openlca/core/model"
    gen_model(java_model)


def gen_model(java_dir):
    files = glob.glob(java_dir + "/*.java")
    for f in files:
        print("## %s" % f)
        with open(f, 'r') as stream:
            handle_class(stream)


def handle_class(stream):
    clazz = None
    for line in stream:
        feed = line.strip()
        if feed.startswith("public class "):
            clazz = handle_class_def(feed)
        if feed.startswith("private ") and not "(" in feed:
            handle_field_def(feed, clazz)



def handle_class_def(line):
    clazz = line.lstrip("public class ").rstrip("{").strip()
    super_clazz = None
    if " extends " in clazz:
        (clazz, super_clazz) = clazz.partition(" extends ")[::2]
    write_class_def(clazz, super_clazz)
    return clazz


def write_class_def(clazz, super_clazz):
    print("\n:%s a rdfs:Class;" % clazz)
    if super_clazz is not None:
        print("\trdfs:subClassOf :%s;" % super_clazz)
    print('\trdfs:comment ""')
    print('\t.\n')


def handle_field_def(line, clazz):
    type, field = line.lstrip("private ").partition(" ")[::2]
    type = get_type(type)
    field = field.rstrip(";")
    print(":%s a rdf:Property;" % field)
    print("\trdfs:domain :%s;" % clazz)
    print("\trdfs:range :%s;" % type)
    print('\trdfs:comment :""')
    print('\t.\n')


def get_type(field_type):
    if field_type == "String":
        return "xsd:string"
    return ":%s" % field_type

if __name__ == '__main__':
    main()