# The `osch` tool

The `osch` tool takes the schema definition and converts it to a defined target.
It is a command line tool written in Go:

```bash
cd olca-schema/osch # in the following, all commands are executed in this folder
go build   # comile the tool
osch help  # prints the help of the tool
osch check # validates the schema definition
```

### Generating the schema documentation

The following command generates the
[mdBook](https://github.com/rust-lang/mdBook) sources from the schema
definition in the `build/docs` folder:

```
$ osch doc
```

With the `mdbook` command, the documentation can be then generated in the
`docs` folder (relative to the book sources) via the following command:

```
$ mdbook build -d ../../docs ../build/docs
```
