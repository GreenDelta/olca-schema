# DO NOT CHANGE THIS CODE AS THIS IS GENERATED AUTOMATICALLY

# This module contains a Python API for reading and writing data sets in
# the JSON based openLCA data exchange format. For more information see
# http://greendelta.github.io/olca-schema

from enum import Enum
from dataclasses import dataclass
from typing import Any, Dict, List, Optional, Union

from .schema import *


@dataclass
class CostValue:

    amount: Optional[float] = None
    currency: Optional[Ref] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.amount is not None:
            d['amount'] = self.amount
        if self.currency is not None:
            d['currency'] = self.currency.to_dict()
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'CostValue':
        cost_value = CostValue()
        if (v := d.get('amount')) or v is not None:
            cost_value.amount = v
        if (v := d.get('currency')) or v is not None:
            cost_value.currency = Ref.from_dict(v)
        return cost_value


@dataclass
class EnviFlow:

    flow: Optional[Ref] = None
    is_input: Optional[bool] = None
    location: Optional[Ref] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.flow is not None:
            d['flow'] = self.flow.to_dict()
        if self.is_input is not None:
            d['isInput'] = self.is_input
        if self.location is not None:
            d['location'] = self.location.to_dict()
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'EnviFlow':
        envi_flow = EnviFlow()
        if (v := d.get('flow')) or v is not None:
            envi_flow.flow = Ref.from_dict(v)
        if (v := d.get('isInput')) or v is not None:
            envi_flow.is_input = v
        if (v := d.get('location')) or v is not None:
            envi_flow.location = Ref.from_dict(v)
        return envi_flow


@dataclass
class EnviFlowValue:

    amount: Optional[float] = None
    envi_flow: Optional[EnviFlow] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.amount is not None:
            d['amount'] = self.amount
        if self.envi_flow is not None:
            d['enviFlow'] = self.envi_flow.to_dict()
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'EnviFlowValue':
        envi_flow_value = EnviFlowValue()
        if (v := d.get('amount')) or v is not None:
            envi_flow_value.amount = v
        if (v := d.get('enviFlow')) or v is not None:
            envi_flow_value.envi_flow = EnviFlow.from_dict(v)
        return envi_flow_value


@dataclass
class ImpactValue:

    amount: Optional[float] = None
    impact_category: Optional[Ref] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.amount is not None:
            d['amount'] = self.amount
        if self.impact_category is not None:
            d['impactCategory'] = self.impact_category.to_dict()
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'ImpactValue':
        impact_value = ImpactValue()
        if (v := d.get('amount')) or v is not None:
            impact_value.amount = v
        if (v := d.get('impactCategory')) or v is not None:
            impact_value.impact_category = Ref.from_dict(v)
        return impact_value


@dataclass
class ResultState:

    id: Optional[str] = None
    error: Optional[str] = None
    is_ready: Optional[bool] = None
    is_scheduled: Optional[bool] = None
    time: Optional[int] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.id is not None:
            d['@id'] = self.id
        if self.error is not None:
            d['error'] = self.error
        if self.is_ready is not None:
            d['isReady'] = self.is_ready
        if self.is_scheduled is not None:
            d['isScheduled'] = self.is_scheduled
        if self.time is not None:
            d['time'] = self.time
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'ResultState':
        result_state = ResultState()
        if (v := d.get('@id')) or v is not None:
            result_state.id = v
        if (v := d.get('error')) or v is not None:
            result_state.error = v
        if (v := d.get('isReady')) or v is not None:
            result_state.is_ready = v
        if (v := d.get('isScheduled')) or v is not None:
            result_state.is_scheduled = v
        if (v := d.get('time')) or v is not None:
            result_state.time = v
        return result_state


@dataclass
class TechFlow:

    flow: Optional[Ref] = None
    provider: Optional[Ref] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.flow is not None:
            d['flow'] = self.flow.to_dict()
        if self.provider is not None:
            d['provider'] = self.provider.to_dict()
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'TechFlow':
        tech_flow = TechFlow()
        if (v := d.get('flow')) or v is not None:
            tech_flow.flow = Ref.from_dict(v)
        if (v := d.get('provider')) or v is not None:
            tech_flow.provider = Ref.from_dict(v)
        return tech_flow


@dataclass
class TechFlowValue:

    amount: Optional[float] = None
    tech_flow: Optional[TechFlow] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.amount is not None:
            d['amount'] = self.amount
        if self.tech_flow is not None:
            d['techFlow'] = self.tech_flow.to_dict()
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'TechFlowValue':
        tech_flow_value = TechFlowValue()
        if (v := d.get('amount')) or v is not None:
            tech_flow_value.amount = v
        if (v := d.get('techFlow')) or v is not None:
            tech_flow_value.tech_flow = TechFlow.from_dict(v)
        return tech_flow_value


@dataclass
class CalculationSetup:

    allocation: Optional[AllocationType] = None
    amount: Optional[float] = None
    flow_property: Optional[Ref] = None
    impact_method: Optional[Ref] = None
    nw_set: Optional[Ref] = None
    parameters: Optional[List[ParameterRedef]] = None
    target: Optional[Ref] = None
    unit: Optional[Ref] = None
    with_costs: Optional[bool] = None
    with_regionalization: Optional[bool] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.allocation is not None:
            d['allocation'] = self.allocation.value
        if self.amount is not None:
            d['amount'] = self.amount
        if self.flow_property is not None:
            d['flowProperty'] = self.flow_property.to_dict()
        if self.impact_method is not None:
            d['impactMethod'] = self.impact_method.to_dict()
        if self.nw_set is not None:
            d['nwSet'] = self.nw_set.to_dict()
        if self.parameters is not None:
            d['parameters'] = [e.to_dict() for e in self.parameters]
        if self.target is not None:
            d['target'] = self.target.to_dict()
        if self.unit is not None:
            d['unit'] = self.unit.to_dict()
        if self.with_costs is not None:
            d['withCosts'] = self.with_costs
        if self.with_regionalization is not None:
            d['withRegionalization'] = self.with_regionalization
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'CalculationSetup':
        calculation_setup = CalculationSetup()
        if (v := d.get('allocation')) or v is not None:
            calculation_setup.allocation = AllocationType.get(v)
        if (v := d.get('amount')) or v is not None:
            calculation_setup.amount = v
        if (v := d.get('flowProperty')) or v is not None:
            calculation_setup.flow_property = Ref.from_dict(v)
        if (v := d.get('impactMethod')) or v is not None:
            calculation_setup.impact_method = Ref.from_dict(v)
        if (v := d.get('nwSet')) or v is not None:
            calculation_setup.nw_set = Ref.from_dict(v)
        if (v := d.get('parameters')) or v is not None:
            calculation_setup.parameters = [ParameterRedef.from_dict(e) for e in v]
        if (v := d.get('target')) or v is not None:
            calculation_setup.target = Ref.from_dict(v)
        if (v := d.get('unit')) or v is not None:
            calculation_setup.unit = Ref.from_dict(v)
        if (v := d.get('withCosts')) or v is not None:
            calculation_setup.with_costs = v
        if (v := d.get('withRegionalization')) or v is not None:
            calculation_setup.with_regionalization = v
        return calculation_setup



RefEntity = Union[RootEntity, Unit, NwSet]
