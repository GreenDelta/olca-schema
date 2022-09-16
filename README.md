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

The data types are defined in simple YAML format. The following things can be
generated from the files in the [yaml](./yaml) folder with the [osch](./osch)
tool:

* the [schema documentation](http://greendelta.github.io/olca-schema)
* a [Python package](https://pypi.org/project/olca-schema/) for reading and
  writing openLCA schema data sets in Json
* a [Protocol Buffers schema](https://github.com/GreenDelta/olca-proto/blob/master/proto/olca.proto)


## License
The openLCA schema is in the worldwide public domain, released under the
[CC0 1.0 Universal Public Domain Dedication](https://creativecommons.org/publicdomain/zero/1.0/).

![Public Domain Dedication](https://licensebuttons.net/p/zero/1.0/88x31.png)
