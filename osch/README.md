# The `osch` tool

The `osch` tool takes the schema definition and converts it to a defined target.
It is a command line tool written in Go:

```bash
cd olca-schema/osch # in the following, all commands are executed in this folder
go build            # comile the tool
./osch help         # prints the help of the tool
./osch check        # validates the schema definition
```


### Generating the schema documentation

The following command generates the
[mdBook](https://github.com/rust-lang/mdBook) sources from the schema
definition in the `build/docs` folder:

```
$ ./osch doc
```

With the `mdbook` command, the documentation can be then generated in the
`docs` folder (relative to the book sources) via the following command:

```
$ mdbook build -d ../../docs ../build/docs
```

This folder is then served as our online documentation via GitHub pages.


### Protocol Buffers

The `proto` command will generate a `olca.proto` file in the `build` folder:

```
$ osch proto
```

This contains then the Protocol Buffers schema. This schema is used in the
[olca-proto project](https://github.com/GreenDelta/olca-proto) and can be
used for data exchange with openLCA via Protocol Buffers and gRPC.


### Python

The class defintions of the Python package can be generated via the `py`
command:

```bash
$ osch py
```
