import olca_schema as lca

units = lca.unit_group_of('Units of mass', 'kg')
kg = units.units[0]
print(units.to_json())

mass = lca.flow_property_of('Mass', units)
print(mass.to_json())

steel = lca.product_flow_of('Steel', mass)
print(steel.to_json())

process = lca.process_of('Steel production')
lca.output_of(process, steel, 1, kg)
print(process)
