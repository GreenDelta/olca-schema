
import olca_schema as o
import unittest


class EnumConvTest(unittest.TestCase):

    def test_flow_type(self):
        val: str = o.FlowType.ELEMENTARY_FLOW.value
        self.assertEqual('ELEMENTARY_FLOW', val)
        self.assertEqual(o.FlowType.ELEMENTARY_FLOW, o.FlowType[val])
        self.assertEqual(o.FlowType.ELEMENTARY_FLOW, o.FlowType.get(val))
        flow = o.Flow(flow_type=o.FlowType.ELEMENTARY_FLOW)
        flow_dict = flow.to_dict()
        self.assertEqual(val, flow_dict['flowType'])
        flow = o.Flow.from_dict(flow_dict)
        self.assertEqual(o.FlowType.ELEMENTARY_FLOW, flow.flow_type)

    def test_process_type(self):
        val: str = o.ProcessType.UNIT_PROCESS.value
        self.assertEqual('UNIT_PROCESS', val)
        self.assertEqual(o.ProcessType.UNIT_PROCESS, o.ProcessType[val])
        self.assertEqual(o.ProcessType.UNIT_PROCESS,
                         o.ProcessType.get(val))
        process = o.Process(process_type=o.ProcessType.UNIT_PROCESS)
        process_dict = process.to_dict()
        self.assertEqual(val, process_dict['processType'])
        process = o.Process.from_dict(process_dict)
        self.assertEqual(o.ProcessType.UNIT_PROCESS, process.process_type)
