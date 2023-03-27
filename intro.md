# openLCA schema

The openLCA schema is the data exchange format of openLCA. It is the only
format that supports all modeling features of openLCA and thus, a lossless data
exchange.

## Why another LCA data exchange format?

There are other LCA data formats, but the concepts of these formats are
sometimes quite different or they are missing features that are important for
openLCA. However, whenever possible we try to be compatible with these formats:

* __EcoSpold 1__: It was one of the first standardized LCA data formats but it
  is not really used anymore. The format is nice and simple but quite some
  essential features are missing, like parameters and formulas. The openLCA
  schema adopted the general structure of processes, with inputs and outputs as
  exchanges, flexible schemes of allocation factors etc., from this format.
* __ISO/TS 14048__: ISO 14048 defines the general requirements for inventory
  data in form of process data sets. The openLCA schema implements these
  requirements for process data sets.
* __SimaPro CSV__: It is similar to EcoSpold 1 regarding the structural
  simplicity (which is good) and the fact that things are identified by names
  (which can be an issue in collaborative environments where name clashes could
  happen). It has some really nice concepts, like parameters on different
  levels, which was adopted in the openLCA schema. One thing that is really
  missing for openLCA in this format, is the concept of product systems: there
  is no way of linking the same processes differently in multiple product
  systems; everything is linked hard by name in this format. Also, there is no
  formal specification of this format; it is a proprietary format with some very
  specific properties.
* __ILCD__: The nice concept of different data set types and references between
  data sets of this format was adopted for the openLCA schema. For data exchange
  (e.g. over web APIs), it has quite some advantages when not everything is
  stored in a single file. The openLCA schema has the same general concepts as
  the ILCD format. However, there are still quite some things missing in the
  ILCD format, like stand-alone global parameters, and other things are a bit
  strange (e.g. it has unit groups with unit conversions but there is no
  possibility for using these units for input and output amounts; LCIA method
  data sets in this format are in fact LCIA categories; etc.). With the extended
  ILCD format eILCD, it is now possible to also define linked product systems in
  this format, but eILCD has a concept of process instances (means copies of
  processes) in product systems which is not so practical for matrix based LCA
  software. Also, the linking is limited in eILCD: the same flow in two
  different exchanges of a process can be only linked to different processes
  when the location is different (a direct linking of exchanges is missing).
* __EcoSpold II__: This format solves some problems of the EcoSpold I format
  (e.g. supports parameters and formulas, provides unique IDs for entities,
  etc.). However, it reflects more the structure and needs of the ecoinvent
  database (e.g. has a concept of parent and child activities) and thus, is
  rarely used in other contexts.

Finally, there are things that are not covered by these formats but supported in
openLCA and the openLCA schema:

* social indicators that can be linked as social aspects in processes
* costs can be attached to every input and output of a process, also as
  formulas, with referenced currencies
* allocation factors for multiple allocation methods can be stored in the same
  process, also using formulas (the calculation picks then the respective
  factors depending on the selected allocation method)
* standalone LCIA categories that can be referenced in different LCIA methods,
  also supporting parameters and formulas
* result data sets that can contain LCI and LCIA results
* EPD data sets that can reference different result data sets for there
  respective life cycle stages / modules
* product system data sets in which processes can be flexibly linked, also to
  sub-systems and results (yes, also results of EPDs)
* ... and more


## Format concepts

The openLCA schema is a typed data format with the type `Entity` as starting
point. An `Entity` is basically a set of key-value pairs, also called fields.
Every field has its specific type which can be:

* a number (integer or floating point number),
* a Boolean value (`true` or `false`),
* a string,
* again an `Entity`,
* or a list of such values.

An entity type can inherit the fields from another entity type where the root of
this inheritance tree is always `Entity`. The type `RefEntity` describes
entities that can be referenced by a unique ID, stored in the field `@id`.
Another data set can point to such a `RefEntity` via a `Ref` that contains that
`@id`. With this, we do not need to repeat the information when the same data
set is referenced multiple times (e.g. when the same flow is used in different
processes).

A `RootEntity` describes a stand-alone data set, like a `Flow` or `Process`.
These data set types form the root of an entity tree. All other entity types are
always part of such an entity tree. For example, a `Unit` is not a `RootEntity`,
it always lives within a `UnitGroup`. But it is a `RefEntity` because it can be
referenced from other entities, e.g. in an `Exchange` of a process.

It should be quite easy to implement these concepts in common programming
languages an serialization formats. In fact, such implementations can be
directly generated from the schema definition.


# JSON-LD / RDF

In most cases, the openLCA schema is just used as a structured data format.
However, the JSON serialization uses the two standard JSON-LD annotations
`@type` and `@id` for the type and identifier of a data set or [data set
reference](./Ref). We also provide a [JSON-LD context](./context.jsonld) and a
[RDF ontology](./schema.ttl) for the schema. The vocabulary base of the schema
is `http://greendelta.github.io/olca-schema#` but the data set identifiers are
relative to the specific context. Instances of the openLCA schema classes or
properties are not part of its ontology. For example, in the following document,
the type is expanded to `http://greendelta.github.io/olca-schema#Flow` but the
ID is relative to the respective context:

```json
{
  "@type": "Flow",
  "@id": "123"
}
```


## Zip packages

For the data exchange with openLCA, openLCA schema data sets are typically
packed as JSON serialized files in zip files. Data sets of the different root
entity types are then stored in the following folders within such a zip file:

| type              | folder              |
|-------------------|---------------------|
| `Actor`           | `actors`            |
| `Currency`        | `currencies`        |
| `DQSystem`        | `dq_systems`        |
| `Epd`             | `epds`              |
| `Flow`            | `flows`             |
| `FlowProperty`    | `flow_properties`   |
| `ImpactCategory`  | `lcia_categories`   |
| `ImpactMethod`    | `lcia_methods`      |
| `Location`        | `locations`         |
| `Parameter`       | `parameters`        |
| `Process`         | `processes`         |
| `ProductSystem`   | `product_systems`   |
| `Project`         | `projects`          |
| `Result`          | `results`           |
| `SocialIndicator` | `social_indicators` |
| `Source`          | `sources`           |
| `UnitGroup`       | `unit_groups`       |

The name of the file is then the ID of the data set followed by the `.json`
extension. At the root level, such a zip-file contains a `olca-schema.json` file
that contains the version of the package format and possibly some other
meta-data:

```
+ actors
  - 23af...e4.json
  - 1e32...f1.json
  - ...
+ ...
- olca-schema.json
```
