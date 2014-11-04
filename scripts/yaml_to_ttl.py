
import model


def get_header():
    return """
@prefix : <http://openlca.org/> .
@prefix owl: <http://www.w3.org/2002/07/owl#> .
@prefix rdfs: <http://www.w3.org/2000/01/rdf-schema#> .
@prefix rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#> .
@prefix xsd: <http://www.w3.org/2001/XMLSchema#> .

:Enumeration a rdfs:Class;
  rdfs:comment "The super class for all enumeration types."
  .
"""

m = model.Model()
m.load_yaml('../yaml')

text_model = get_header()



print(text_model)


