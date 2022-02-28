# Changes in version 2

## New types

The following new types were added to the schema:

* `Epd`
* `ParameterRedefSet`
* `Result`
  * `ImpactResult`
  * `FlowResult`


## New fields

The following fields were added to existing classes:

| class              | field                   | type                      |
|--------------------|-------------------------|---------------------------|
| `AllocationFactor` | `formula`               | `string`                  |
| `Exchange`         | `location`              | `Ref[Location]`           |
| `ImpactCategory`   | `code`                  | `string`                  |
| `ImpactCategory`   | `parameters`            | `List[Parameter`          |
| `ImpactCategory`   | `source`                | `Ref[Source]`             |
| `ImpactFactor`     | `location`              | `Ref[Location]`           |
| `ImpactMethod`     | `code`                  | `string`                  |
| `ImpactMethod`     | `source`                | `Ref[Source]`             |
| `ParameterRedef`   | `isProtected`           | `boolean`                 |
| `ProductSystem`    | `parameterSets`         | `List[ParameterRedefSet]` |
| `Project`          | `isWithCosts`           | `boolean`                 |
| `Project`          | `isWithRegionalization` | `boolean`                 |
| `ProjectVariant`   | `description`           | `string`                  |
| `ProjectVariant`   | `isDisabled`            | `boolean`                 |


## Renamed fields

| class                  | old name                | new name
|------------------------|-------------------------|---------------------------
| `Exchange`             | `avoidedProduct`        | `isAvoidedProduct`
| `Exchange`             | `input`                 | `isInput`
| `Exchange`             | `quantitativeReference` | `isQuantitativeReference`
| `Flow`                 | `infrastructureFlow`    | `isInfrastructureFlow`
| `FlowPropertyFactor`   | `referenceFlowProperty` | `isReferenceFlowProperty`
| `ImpactCategory`       | `referenceUnitName`     | `refUnit`
| `Parameter`            | `inputParameter`        | `isIputParameter`
| `Process`              | `infrastructureProcess` | `isInfrastructureProcess`
| `ProcessDocumentation` | `copyright`             | `hasCopyright`
| `Unit`                 | `referenceUnit`         | `isReferenceUnit`


## Type changes

* the field `category` in `RootEntity` is now a string of the full category path
  instead of a reference to a category
* the abstract class `RootEntity` was renamed to `RefEntity`; the abstract
  class `CategorizedEntity` to `RootEntity`


## Removed fields

The following fields were removed:

| class              | field
|--------------------|----------------
| `ImpactMethod`     | `parameterMean`
| `ImpactMethod`     | `parameters`
| `ImpactMethod`     | `parameters`
| `Parameter`        | `externalSource`
| `Parameter`        | `sourceType`
| `ProcessLink`      | `isSystemLink`
| `ProductSystem`    | `parameterRedefs`
| `ProductSystem`    | `parameterRedefs`
| `Project`          | `creationDate`
| `Project`          | `functionalUnit`
| `Project`          | `goal`
| `Project`          | `lastModificationDate`
| `Project`          | `author`
