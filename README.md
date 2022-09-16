# openLCA schema

The openLCA schema is the data exchange format of
[openLCA](http://www.openlca.org/). This repository contains the specification
of this format. The best way to read this specification is to:

[__Read the documentation online__](http://greendelta.github.io/olca-schema)


## Contributing

The openLCA schema reflects the openLCA model. If we change the schema, we need
to change the openLCA software. This can be quite some effort. However, format
extensions in form of additional fields are easy and could be added if they
are useful (as they could be just ignored if not supported by the software yet).
What needs to be improved always, is the documentation of the format. The data
set types and fields are specified and documented in their respective files in
the [yaml](./yaml) folder. Pull requests are welcome.


## Building the artifacts



Currently, openLCA supports
[JSON](https://www.json.org) and [Protocol
Buffers](https://developers.google.com/protocol-buffers) as serialization
formats over files, REST/IPC services, or [gRPC](https://grpc.io/). In addition,
a [JSON-LD](https://json-ld.org/) context is provided so that it can be used as
a Linked Data format. openLCA schema is based on a few primitive building blocks
like strings, numbers, lists and structured data types that can be easily
implemented in different programming languages or serialization formats. For
custom structured data types, we use the term `class`. These start with an
uppercase letter and map to structures in the respective implementation (e.g.
objects in JSON, messages in Protocol Buffers, classes and instances in Python etc.).
Primitive data types, which start with a lower case letter, as well as lists
(type `List`) are typically directly provided by the respective format
implementation. Classes have properties (or fields) which can hold values of a
specific type. In addition, there are a few enumeration types.

The different data types are defined in simple [YAML](http://yaml.org/) files
with a file for each type in the [yaml folder](./yaml)... You can also browse
the [HTML documentation online](http://greendelta.github.io/olca-schema).

## License
This openLCA Schema is in the worldwide public domain, released under the [CC0
1.0 Universal Public Domain
Dedication](https://creativecommons.org/publicdomain/zero/1.0/).

![Public Domain Dedication](https://licensebuttons.net/p/zero/1.0/88x31.png)
