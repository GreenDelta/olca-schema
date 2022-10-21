# DO NOT CHANGE THIS CODE AS THIS IS GENERATED AUTOMATICALLY

# This module contains a Python API for reading and writing data sets in
# the JSON based openLCA data exchange format. For more information see
# http://greendelta.github.io/olca-schema

from enum import Enum
from dataclasses import dataclass
from typing import Any, Dict, List, Optional, Union

from .schema import *


@dataclass
class EnviFlow:

    flow: Optional[Ref] = None
    is_input: Optional[bool] = None
    location: Optional[Ref] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.flow:
            d['flow'] = self.flow.to_dict()
        if self.is_input:
            d['isInput'] = self.is_input
        if self.location:
            d['location'] = self.location.to_dict()
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'EnviFlow':
        envi_flow = EnviFlow()
        if v := d.get('flow'):
            envi_flow.flow = Ref.from_dict(v)
        if v := d.get('isInput'):
            envi_flow.is_input = v
        if v := d.get('location'):
            envi_flow.location = Ref.from_dict(v)
        return envi_flow


@dataclass
class EnviFlowValue:

    amount: Optional[float] = None
    envi_flow: Optional[EnviFlow] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.amount:
            d['amount'] = self.amount
        if self.envi_flow:
            d['enviFlow'] = self.envi_flow.to_dict()
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'EnviFlowValue':
        envi_flow_value = EnviFlowValue()
        if v := d.get('amount'):
            envi_flow_value.amount = v
        if v := d.get('enviFlow'):
            envi_flow_value.envi_flow = EnviFlow.from_dict(v)
        return envi_flow_value


@dataclass
class ImpactValue:

    amount: Optional[float] = None
    impact_category: Optional[Ref] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.amount:
            d['amount'] = self.amount
        if self.impact_category:
            d['impactCategory'] = self.impact_category.to_dict()
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'ImpactValue':
        impact_value = ImpactValue()
        if v := d.get('amount'):
            impact_value.amount = v
        if v := d.get('impactCategory'):
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
        if self.id:
            d['@id'] = self.id
        if self.error:
            d['error'] = self.error
        if self.is_ready:
            d['isReady'] = self.is_ready
        if self.is_scheduled:
            d['isScheduled'] = self.is_scheduled
        if self.time:
            d['time'] = self.time
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'ResultState':
        result_state = ResultState()
        if v := d.get('@id'):
            result_state.id = v
        if v := d.get('error'):
            result_state.error = v
        if v := d.get('isReady'):
            result_state.is_ready = v
        if v := d.get('isScheduled'):
            result_state.is_scheduled = v
        if v := d.get('time'):
            result_state.time = v
        return result_state


@dataclass
class TechFlow:

    flow: Optional[Ref] = None
    provider: Optional[Ref] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.flow:
            d['flow'] = self.flow.to_dict()
        if self.provider:
            d['provider'] = self.provider.to_dict()
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'TechFlow':
        tech_flow = TechFlow()
        if v := d.get('flow'):
            tech_flow.flow = Ref.from_dict(v)
        if v := d.get('provider'):
            tech_flow.provider = Ref.from_dict(v)
        return tech_flow


@dataclass
class TechFlowValue:

    amount: Optional[float] = None
    tech_flow: Optional[TechFlow] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.amount:
            d['amount'] = self.amount
        if self.tech_flow:
            d['techFlow'] = self.tech_flow.to_dict()
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'TechFlowValue':
        tech_flow_value = TechFlowValue()
        if v := d.get('amount'):
            tech_flow_value.amount = v
        if v := d.get('techFlow'):
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
        if self.allocation:
            d['allocation'] = self.allocation.value
        if self.amount:
            d['amount'] = self.amount
        if self.flow_property:
            d['flowProperty'] = self.flow_property.to_dict()
        if self.impact_method:
            d['impactMethod'] = self.impact_method.to_dict()
        if self.nw_set:
            d['nwSet'] = self.nw_set.to_dict()
        if self.parameters:
            d['parameters'] = [e.to_dict() for e in self.parameters]
        if self.target:
            d['target'] = self.target.to_dict()
        if self.unit:
            d['unit'] = self.unit.to_dict()
        if self.with_costs:
            d['withCosts'] = self.with_costs
        if self.with_regionalization:
            d['withRegionalization'] = self.with_regionalization
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'CalculationSetup':
        calculation_setup = CalculationSetup()
        if v := d.get('allocation'):
            calculation_setup.allocation = AllocationType.get(v)
        if v := d.get('amount'):
            calculation_setup.amount = v
        if v := d.get('flowProperty'):
            calculation_setup.flow_property = Ref.from_dict(v)
        if v := d.get('impactMethod'):
            calculation_setup.impact_method = Ref.from_dict(v)
        if v := d.get('nwSet'):
            calculation_setup.nw_set = Ref.from_dict(v)
        if v := d.get('parameters'):
            calculation_setup.parameters = [ParameterRedef.from_dict(e) for e in v]
        if v := d.get('target'):
            calculation_setup.target = Ref.from_dict(v)
        if v := d.get('unit'):
            calculation_setup.unit = Ref.from_dict(v)
        if v := d.get('withCosts'):
            calculation_setup.with_costs = v
        if v := d.get('withRegionalization'):
            calculation_setup.with_regionalization = v
        return calculation_setup


