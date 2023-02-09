import unittest

import olca_schema as o


class TestFactory(unittest.TestCase):
    def test_unit(self):
        kg = o.new_unit("kg")
        self.assertEqual(kg.name, "kg")
        self.assertEqual(kg.conversion_factor, 1.0)

    def test_unit_group(self):
        units = o.new_unit_group("Units of mass", "kg")
        self.assertEqual(units.units[0].name, "kg")

    def test_flow_property(self):
        units = o.new_unit_group("Units of mass", "kg")
        mass = o.new_flow_property("Mass", units)
        self.assertEqual(mass.name, "Mass")
        self.assertEqual(mass.unit_group.name, units.name)

    def test_flow(self):
        units = o.new_unit_group("Units of mass", "kg")
        mass = o.new_flow_property("Mass", units)
        steel = o.new_flow("Steel", o.FlowType.PRODUCT_FLOW, mass)
        self.assertEqual(steel.name, "Steel")
        self.assertEqual(steel.flow_properties[0].flow_property.name, "Mass")

    def test_product(self):
        units = o.new_unit_group("Units of mass", "kg")
        mass = o.new_flow_property("Mass", units)
        steel = o.new_product("Steel", mass)
        self.assertEqual(o.FlowType.PRODUCT_FLOW, steel.flow_type)
        self.assertEqual(steel.name, "Steel")
        self.assertEqual(steel.flow_properties[0].flow_property.name, "Mass")

    def test_waste(self):
        units = o.new_unit_group("Units of mass", "kg")
        mass = o.new_flow_property("Mass", units)
        scrap = o.new_waste("Scrap", mass)
        self.assertEqual(o.FlowType.WASTE_FLOW, scrap.flow_type)
        self.assertEqual(scrap.name, "Scrap")
        self.assertEqual(scrap.flow_properties[0].flow_property.name, "Mass")

    def test_elementary_flow(self):
        units = o.new_unit_group("Units of mass", "kg")
        mass = o.new_flow_property("Mass", units)
        co2 = o.new_elementary_flow("CO2", mass)
        self.assertEqual(o.FlowType.ELEMENTARY_FLOW, co2.flow_type)
        self.assertEqual(co2.name, "CO2")
        self.assertEqual(co2.flow_properties[0].flow_property.name, "Mass")

    def test_process(self):
        process = o.new_process("Steel production")
        self.assertEqual(process.name, "Steel production")

    def test_exchange(self):
        units = o.new_unit_group("Units of mass", "kg")
        mass = o.new_flow_property("Mass", units)
        steel = o.new_product("Steel", mass)
        process = o.new_process("Steel production")
        output = o.new_output(process, steel)
        output.quantitative_reference = True
        self.assertEqual(1, len(process.exchanges))
        self.assertEqual(output.flow.name, "Steel")
        self.assertEqual(output.amount, 1.0)


if __name__ == "__main__":
    unittest.main()
