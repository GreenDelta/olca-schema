# olca-schema

This is a package for reading and writing data sets in the [openLCA
schema](https://github.com/GreenDelta/olca-schema) format version 2. It provides
a typed class model of the schema elements, methods for reading and writing them
in the JSON format, reading and writing data sets in zip packages, and some
utility methods.

## Usage

The package is published on [PyPI](https://pypi.org/project/olca-schema/) and
can be installed with `pip`:

```bash
pip install olca-schema
```

Here is a small example that creates a package that can be imported into
openLCA 2:

```python
import olca_schema as lca
import olca_schema.zipio as zipio

# create a unit group and flow property
units = lca.new_unit_group('Units of mass', 'kg')
kg = units.units[0]
mass = lca.new_flow_property('Mass', units)

# create a product flow and a process with the product as output
steel = lca.new_product('Steel', mass)
process = lca.new_process('Steel production')
output = lca.new_output(process, steel, 1, kg)
output.is_quantitative_reference = True

# prints the Json string of the process
print(process.to_json())

# write a zip package with the data sets
with zipio.ZipWriter('path/to/example.zip') as w:
    for entity in [units, mass, steel, process]:
        w.write(entity)
```
