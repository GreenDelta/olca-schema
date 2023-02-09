import olca_schema as o
import olca_schema.zipio as zipio


def main():
    units = o.new_unit_group('Units of mass', 'kg')
    kg = units.units[0]
    print(units.to_json())

    mass = o.new_flow_property('Mass', units)
    print(mass.to_json())

    steel = o.new_product('Steel', mass)
    print(steel.to_json())

    process = o.new_process('Steel production')
    output = o.new_output(process, steel, 1, kg)
    output.is_quantitative_reference = True
    print(process.to_json())

    with zipio.ZipWriter('../build/example.zip') as w:
        for entity in [units, mass, steel, process]:
            w.write(entity)


if __name__ == '__main__':
    main()
