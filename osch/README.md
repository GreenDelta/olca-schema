# The `osch` tool

The `osch` tool takes the schema definition and converts it to a defined target.
Currently, the following targets are implemented:

* ...

### Building

```bash
cd oschgo ; go build -o ../osch ; cd ..
```

or on Windows:

```batch
cd oschgo && go build -o ..\osch.exe && cd ..
```

### Usage

```
usage:

$ osch [command] [options]

commands:

  check  - checks the schema
  doc    - generates the schema documentation
  help   - prints this help
  proto  - converts the schema to ProtocolBuffers
  python - generates a Python class model for the schema

```

