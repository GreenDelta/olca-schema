import unittest

import olca_schema as lca


class TestFactory(unittest.TestCase):
    def test_unit(self):
        kg = lca.new_unit("kg")
        self.assertEqual(kg.name, "kg")
        self.assertEqual(kg.conversion_factor, 1.0)

    def test_unit_group(self):
        units = lca.new_unit_group("Units of mass", "kg")
        self.assertEqual(units.units[0].name, "kg")

    def test_flow_property(self):
        units = lca.new_unit_group("Units of mass", "kg")
        mass = lca.new_flow_property("Mass", units)
        self.assertEqual(mass.name, "Mass")
        self.assertEqual(mass.unit_group.name, units.name)

    def test_flow(self):
        units = lca.new_unit_group("Units of mass", "kg")
        mass = lca.new_flow_property("Mass", units)
        steel = lca.new_flow("Steel", lca.FlowType.PRODUCT_FLOW, mass)
        self.assertEqual(steel.name, "Steel")
        self.assertEqual(steel.flow_properties[0].flow_property.name, "Mass")

    def test_product(self):
        units = lca.new_unit_group("Units of mass", "kg")
        mass = lca.new_flow_property("Mass", units)
        steel = lca.new_product("Steel", mass)
        self.assertEqual(lca.FlowType.PRODUCT_FLOW, steel.flow_type)
        self.assertEqual(steel.name, "Steel")
        self.assertEqual(steel.flow_properties[0].flow_property.name, "Mass")

    def test_waste(self):
        units = lca.new_unit_group("Units of mass", "kg")
        mass = lca.new_flow_property("Mass", units)
        scrap = lca.new_waste("Scrap", mass)
        self.assertEqual(lca.FlowType.WASTE_FLOW, scrap.flow_type)
        self.assertEqual(scrap.name, "Scrap")
        self.assertEqual(scrap.flow_properties[0].flow_property.name, "Mass")

    def test_elementary_flow(self):
        units = lca.new_unit_group("Units of mass", "kg")
        mass = lca.new_flow_property("Mass", units)
        co2 = lca.new_elementary_flow("CO2", mass)
        self.assertEqual(lca.FlowType.ELEMENTARY_FLOW, co2.flow_type)
        self.assertEqual(co2.name, "CO2")
        self.assertEqual(co2.flow_properties[0].flow_property.name, "Mass")

    def test_process(self):
        process = lca.new_process("Steel production")
        self.assertEqual(process.name, "Steel production")

    def test_exchange(self):
        units = lca.new_unit_group("Units of mass", "kg")
        mass = lca.new_flow_property("Mass", units)
        steel = lca.new_product("Steel", mass)
        process = lca.new_process("Steel production")
        output = lca.new_output(process, steel)
        output.quantitative_reference = True
        self.assertEqual(1, len(process.exchanges))
        self.assertEqual(output.flow.name, "Steel")
        self.assertEqual(output.amount, 1.0)


if __name__ == "__main__":
    unittest.main()
