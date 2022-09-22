
import olca_schema as lca
import unittest


class EnumConvTest(unittest.TestCase):

    def test_flow_type(self):
        val: str = lca.FlowType.ELEMENTARY_FLOW.value
        self.assertEqual('ELEMENTARY_FLOW', val)
        self.assertEqual(lca.FlowType.ELEMENTARY_FLOW, lca.FlowType[val])
        self.assertEqual(lca.FlowType.ELEMENTARY_FLOW, lca.FlowType.get(val))
        flow = lca.Flow(flow_type=lca.FlowType.ELEMENTARY_FLOW)
        flow_dict = flow.to_dict()
        self.assertEqual(val, flow_dict['flowType'])
        flow = lca.Flow.from_dict(flow_dict)
        self.assertEqual(lca.FlowType.ELEMENTARY_FLOW, flow.flow_type)

    def test_process_type(self):
        val: str = lca.ProcessType.UNIT_PROCESS.value
        self.assertEqual('UNIT_PROCESS', val)
        self.assertEqual(lca.ProcessType.UNIT_PROCESS, lca.ProcessType[val])
        self.assertEqual(lca.ProcessType.UNIT_PROCESS,
                         lca.ProcessType.get(val))
        process = lca.Process(process_type=lca.ProcessType.UNIT_PROCESS)
        process_dict = process.to_dict()
        self.assertEqual(val, process_dict['processType'])
        process = lca.Process.from_dict(process_dict)
        self.assertEqual(lca.ProcessType.UNIT_PROCESS, process.process_type)
