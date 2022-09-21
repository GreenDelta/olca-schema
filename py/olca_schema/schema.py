# DO NOT CHANGE THIS CODE AS THIS IS GENERATED AUTOMATICALLY

# This module contains a Python API for reading and writing data sets in
# the JSON based openLCA data exchange format. For more information see
# http://greendelta.github.io/olca-schema

import datetime
import json
import uuid

from enum import Enum
from dataclasses import dataclass
from typing import Any, Dict, List, Optional, Union


class AllocationType(Enum):

    PHYSICAL_ALLOCATION = 'PHYSICAL_ALLOCATION'
    ECONOMIC_ALLOCATION = 'ECONOMIC_ALLOCATION'
    CAUSAL_ALLOCATION = 'CAUSAL_ALLOCATION'
    USE_DEFAULT_ALLOCATION = 'USE_DEFAULT_ALLOCATION'
    NO_ALLOCATION = 'NO_ALLOCATION'

    def get(v: Union[str, 'AllocationType'],
            default: Optional['AllocationType'] = None) -> 'AllocationType':
        for i in AllocationType:
            if i == v or i.value == v or i.name == v:
                return i
        return default


class Direction(Enum):

    INPUT = 'INPUT'
    OUTPUT = 'OUTPUT'

    def get(v: Union[str, 'Direction'],
            default: Optional['Direction'] = None) -> 'Direction':
        for i in Direction:
            if i == v or i.value == v or i.name == v:
                return i
        return default


class FlowPropertyType(Enum):

    ECONOMIC_QUANTITY = 'ECONOMIC_QUANTITY'
    PHYSICAL_QUANTITY = 'PHYSICAL_QUANTITY'

    def get(v: Union[str, 'FlowPropertyType'],
            default: Optional['FlowPropertyType'] = None) -> 'FlowPropertyType':
        for i in FlowPropertyType:
            if i == v or i.value == v or i.name == v:
                return i
        return default


class FlowType(Enum):

    ELEMENTARY_FLOW = 'ELEMENTARY_FLOW'
    PRODUCT_FLOW = 'PRODUCT_FLOW'
    WASTE_FLOW = 'WASTE_FLOW'

    def get(v: Union[str, 'FlowType'],
            default: Optional['FlowType'] = None) -> 'FlowType':
        for i in FlowType:
            if i == v or i.value == v or i.name == v:
                return i
        return default


class ModelType(Enum):

    ACTOR = 'ACTOR'
    CATEGORY = 'CATEGORY'
    CURRENCY = 'CURRENCY'
    DQ_SYSTEM = 'DQ_SYSTEM'
    EPD = 'EPD'
    FLOW = 'FLOW'
    FLOW_PROPERTY = 'FLOW_PROPERTY'
    IMPACT_CATEGORY = 'IMPACT_CATEGORY'
    IMPACT_METHOD = 'IMPACT_METHOD'
    LOCATION = 'LOCATION'
    PARAMETER = 'PARAMETER'
    PROCESS = 'PROCESS'
    PRODUCT_SYSTEM = 'PRODUCT_SYSTEM'
    PROJECT = 'PROJECT'
    RESULT = 'RESULT'
    SOCIAL_INDICATOR = 'SOCIAL_INDICATOR'
    SOURCE = 'SOURCE'
    UNIT_GROUP = 'UNIT_GROUP'

    def get(v: Union[str, 'ModelType'],
            default: Optional['ModelType'] = None) -> 'ModelType':
        for i in ModelType:
            if i == v or i.value == v or i.name == v:
                return i
        return default


class ParameterScope(Enum):

    PROCESS_SCOPE = 'PROCESS_SCOPE'
    IMPACT_SCOPE = 'IMPACT_SCOPE'
    GLOBAL_SCOPE = 'GLOBAL_SCOPE'

    def get(v: Union[str, 'ParameterScope'],
            default: Optional['ParameterScope'] = None) -> 'ParameterScope':
        for i in ParameterScope:
            if i == v or i.value == v or i.name == v:
                return i
        return default


class ProcessType(Enum):

    LCI_RESULT = 'LCI_RESULT'
    UNIT_PROCESS = 'UNIT_PROCESS'

    def get(v: Union[str, 'ProcessType'],
            default: Optional['ProcessType'] = None) -> 'ProcessType':
        for i in ProcessType:
            if i == v or i.value == v or i.name == v:
                return i
        return default


class RiskLevel(Enum):

    NO_OPPORTUNITY = 'NO_OPPORTUNITY'
    HIGH_OPPORTUNITY = 'HIGH_OPPORTUNITY'
    MEDIUM_OPPORTUNITY = 'MEDIUM_OPPORTUNITY'
    LOW_OPPORTUNITY = 'LOW_OPPORTUNITY'
    NO_RISK = 'NO_RISK'
    VERY_LOW_RISK = 'VERY_LOW_RISK'
    LOW_RISK = 'LOW_RISK'
    MEDIUM_RISK = 'MEDIUM_RISK'
    HIGH_RISK = 'HIGH_RISK'
    VERY_HIGH_RISK = 'VERY_HIGH_RISK'
    NO_DATA = 'NO_DATA'
    NOT_APPLICABLE = 'NOT_APPLICABLE'

    def get(v: Union[str, 'RiskLevel'],
            default: Optional['RiskLevel'] = None) -> 'RiskLevel':
        for i in RiskLevel:
            if i == v or i.value == v or i.name == v:
                return i
        return default


class UncertaintyType(Enum):

    LOG_NORMAL_DISTRIBUTION = 'LOG_NORMAL_DISTRIBUTION'
    NORMAL_DISTRIBUTION = 'NORMAL_DISTRIBUTION'
    TRIANGLE_DISTRIBUTION = 'TRIANGLE_DISTRIBUTION'
    UNIFORM_DISTRIBUTION = 'UNIFORM_DISTRIBUTION'

    def get(v: Union[str, 'UncertaintyType'],
            default: Optional['UncertaintyType'] = None) -> 'UncertaintyType':
        for i in UncertaintyType:
            if i == v or i.value == v or i.name == v:
                return i
        return default


@dataclass
class DQScore:

    description: Optional[str] = None
    label: Optional[str] = None
    position: Optional[int] = None
    uncertainty: Optional[float] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.description:
            d['description'] = self.description
        if self.label:
            d['label'] = self.label
        if self.position:
            d['position'] = self.position
        if self.uncertainty:
            d['uncertainty'] = self.uncertainty
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'DQScore':
        d_q_score = DQScore()
        if v := d.get('@type'):
            d_q_score.schema_type = v
        if v := d.get('description'):
            d_q_score.description = v
        if v := d.get('label'):
            d_q_score.label = v
        if v := d.get('position'):
            d_q_score.position = v
        if v := d.get('uncertainty'):
            d_q_score.uncertainty = v
        return d_q_score


@dataclass
class DQIndicator:

    name: Optional[str] = None
    position: Optional[int] = None
    scores: Optional[List[DQScore]] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.name:
            d['name'] = self.name
        if self.position:
            d['position'] = self.position
        if self.scores:
            d['scores'] = [e.to_dict() for e in self.scores]
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'DQIndicator':
        d_q_indicator = DQIndicator()
        if v := d.get('@type'):
            d_q_indicator.schema_type = v
        if v := d.get('name'):
            d_q_indicator.name = v
        if v := d.get('position'):
            d_q_indicator.position = v
        if v := d.get('scores'):
            d_q_indicator.scores = [DQScore.from_dict(e) for e in v]
        return d_q_indicator


@dataclass
class ExchangeRef:

    internal_id: Optional[int] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.internal_id:
            d['internalId'] = self.internal_id
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'ExchangeRef':
        exchange_ref = ExchangeRef()
        if v := d.get('@type'):
            exchange_ref.schema_type = v
        if v := d.get('internalId'):
            exchange_ref.internal_id = v
        return exchange_ref


@dataclass
class Ref:

    id: Optional[str] = None
    category: Optional[str] = None
    description: Optional[str] = None
    flow_type: Optional[FlowType] = None
    location: Optional[str] = None
    name: Optional[str] = None
    process_type: Optional[ProcessType] = None
    ref_unit: Optional[str] = None
    model_type: str = ''

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        d['@type'] = self.model_type
        if self.id:
            d['@id'] = self.id
        if self.category:
            d['category'] = self.category
        if self.description:
            d['description'] = self.description
        if self.flow_type:
            d['flowType'] = self.flow_type.value
        if self.location:
            d['location'] = self.location
        if self.name:
            d['name'] = self.name
        if self.process_type:
            d['processType'] = self.process_type.value
        if self.ref_unit:
            d['refUnit'] = self.ref_unit
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'Ref':
        ref = Ref()
        ref.model_type = d.get('@type', '')
        if v := d.get('@type'):
            ref.schema_type = v
        if v := d.get('@id'):
            ref.id = v
        if v := d.get('category'):
            ref.category = v
        if v := d.get('description'):
            ref.description = v
        if v := d.get('flowType'):
            ref.flow_type = FlowType.get(v)
        if v := d.get('location'):
            ref.location = v
        if v := d.get('name'):
            ref.name = v
        if v := d.get('processType'):
            ref.process_type = ProcessType.get(v)
        if v := d.get('refUnit'):
            ref.ref_unit = v
        return ref


@dataclass
class Actor:

    id: Optional[str] = None
    address: Optional[str] = None
    category: Optional[str] = None
    city: Optional[str] = None
    country: Optional[str] = None
    description: Optional[str] = None
    email: Optional[str] = None
    last_change: Optional[str] = None
    name: Optional[str] = None
    tags: Optional[List[str]] = None
    telefax: Optional[str] = None
    telephone: Optional[str] = None
    version: Optional[str] = None
    website: Optional[str] = None
    zip_code: Optional[str] = None

    def __post_init__(self):
        if self.id is None:
            self.id = str(uuid.uuid4())
        if self.version is None:
            self.version = '01.00.000'
        if self.last_change is None:
            self.last_change = datetime.datetime.utcnow().isoformat() + 'Z'

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        d['@type'] = 'Actor'
        if self.id:
            d['@id'] = self.id
        if self.address:
            d['address'] = self.address
        if self.category:
            d['category'] = self.category
        if self.city:
            d['city'] = self.city
        if self.country:
            d['country'] = self.country
        if self.description:
            d['description'] = self.description
        if self.email:
            d['email'] = self.email
        if self.last_change:
            d['lastChange'] = self.last_change
        if self.name:
            d['name'] = self.name
        if self.tags:
            d['tags'] = self.tags
        if self.telefax:
            d['telefax'] = self.telefax
        if self.telephone:
            d['telephone'] = self.telephone
        if self.version:
            d['version'] = self.version
        if self.website:
            d['website'] = self.website
        if self.zip_code:
            d['zipCode'] = self.zip_code
        return d

    def to_json(self) -> str:
        return json.dumps(self.to_dict(), indent=2)

    def to_ref(self) -> 'Ref':
        ref = Ref(id=self.id, name=self.name)
        ref.category = self.category
        ref.model_type = 'Actor'
        return ref

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'Actor':
        actor = Actor()
        if v := d.get('@type'):
            actor.schema_type = v
        if v := d.get('@id'):
            actor.id = v
        if v := d.get('address'):
            actor.address = v
        if v := d.get('category'):
            actor.category = v
        if v := d.get('city'):
            actor.city = v
        if v := d.get('country'):
            actor.country = v
        if v := d.get('description'):
            actor.description = v
        if v := d.get('email'):
            actor.email = v
        if v := d.get('lastChange'):
            actor.last_change = v
        if v := d.get('name'):
            actor.name = v
        if v := d.get('tags'):
            actor.tags = v
        if v := d.get('telefax'):
            actor.telefax = v
        if v := d.get('telephone'):
            actor.telephone = v
        if v := d.get('version'):
            actor.version = v
        if v := d.get('website'):
            actor.website = v
        if v := d.get('zipCode'):
            actor.zip_code = v
        return actor

    @staticmethod
    def from_json(data: Union[str, bytes]) -> 'Actor':
        return Actor.from_dict(json.loads(data))


@dataclass
class AllocationFactor:

    allocation_type: Optional[AllocationType] = None
    exchange: Optional[ExchangeRef] = None
    formula: Optional[str] = None
    product: Optional[Ref] = None
    value: Optional[float] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.allocation_type:
            d['allocationType'] = self.allocation_type.value
        if self.exchange:
            d['exchange'] = self.exchange.to_dict()
        if self.formula:
            d['formula'] = self.formula
        if self.product:
            d['product'] = self.product.to_dict()
        if self.value:
            d['value'] = self.value
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'AllocationFactor':
        allocation_factor = AllocationFactor()
        if v := d.get('@type'):
            allocation_factor.schema_type = v
        if v := d.get('allocationType'):
            allocation_factor.allocation_type = AllocationType.get(v)
        if v := d.get('exchange'):
            allocation_factor.exchange = ExchangeRef.from_dict(v)
        if v := d.get('formula'):
            allocation_factor.formula = v
        if v := d.get('product'):
            allocation_factor.product = Ref.from_dict(v)
        if v := d.get('value'):
            allocation_factor.value = v
        return allocation_factor


@dataclass
class Currency:

    id: Optional[str] = None
    category: Optional[str] = None
    code: Optional[str] = None
    conversion_factor: Optional[float] = None
    description: Optional[str] = None
    last_change: Optional[str] = None
    name: Optional[str] = None
    ref_currency: Optional[Ref] = None
    tags: Optional[List[str]] = None
    version: Optional[str] = None

    def __post_init__(self):
        if self.id is None:
            self.id = str(uuid.uuid4())
        if self.version is None:
            self.version = '01.00.000'
        if self.last_change is None:
            self.last_change = datetime.datetime.utcnow().isoformat() + 'Z'

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        d['@type'] = 'Currency'
        if self.id:
            d['@id'] = self.id
        if self.category:
            d['category'] = self.category
        if self.code:
            d['code'] = self.code
        if self.conversion_factor:
            d['conversionFactor'] = self.conversion_factor
        if self.description:
            d['description'] = self.description
        if self.last_change:
            d['lastChange'] = self.last_change
        if self.name:
            d['name'] = self.name
        if self.ref_currency:
            d['refCurrency'] = self.ref_currency.to_dict()
        if self.tags:
            d['tags'] = self.tags
        if self.version:
            d['version'] = self.version
        return d

    def to_json(self) -> str:
        return json.dumps(self.to_dict(), indent=2)

    def to_ref(self) -> 'Ref':
        ref = Ref(id=self.id, name=self.name)
        ref.category = self.category
        ref.model_type = 'Currency'
        return ref

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'Currency':
        currency = Currency()
        if v := d.get('@type'):
            currency.schema_type = v
        if v := d.get('@id'):
            currency.id = v
        if v := d.get('category'):
            currency.category = v
        if v := d.get('code'):
            currency.code = v
        if v := d.get('conversionFactor'):
            currency.conversion_factor = v
        if v := d.get('description'):
            currency.description = v
        if v := d.get('lastChange'):
            currency.last_change = v
        if v := d.get('name'):
            currency.name = v
        if v := d.get('refCurrency'):
            currency.ref_currency = Ref.from_dict(v)
        if v := d.get('tags'):
            currency.tags = v
        if v := d.get('version'):
            currency.version = v
        return currency

    @staticmethod
    def from_json(data: Union[str, bytes]) -> 'Currency':
        return Currency.from_dict(json.loads(data))


@dataclass
class DQSystem:

    id: Optional[str] = None
    category: Optional[str] = None
    description: Optional[str] = None
    has_uncertainties: Optional[bool] = None
    indicators: Optional[List[DQIndicator]] = None
    last_change: Optional[str] = None
    name: Optional[str] = None
    source: Optional[Ref] = None
    tags: Optional[List[str]] = None
    version: Optional[str] = None

    def __post_init__(self):
        if self.id is None:
            self.id = str(uuid.uuid4())
        if self.version is None:
            self.version = '01.00.000'
        if self.last_change is None:
            self.last_change = datetime.datetime.utcnow().isoformat() + 'Z'

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        d['@type'] = 'DQSystem'
        if self.id:
            d['@id'] = self.id
        if self.category:
            d['category'] = self.category
        if self.description:
            d['description'] = self.description
        if self.has_uncertainties:
            d['hasUncertainties'] = self.has_uncertainties
        if self.indicators:
            d['indicators'] = [e.to_dict() for e in self.indicators]
        if self.last_change:
            d['lastChange'] = self.last_change
        if self.name:
            d['name'] = self.name
        if self.source:
            d['source'] = self.source.to_dict()
        if self.tags:
            d['tags'] = self.tags
        if self.version:
            d['version'] = self.version
        return d

    def to_json(self) -> str:
        return json.dumps(self.to_dict(), indent=2)

    def to_ref(self) -> 'Ref':
        ref = Ref(id=self.id, name=self.name)
        ref.category = self.category
        ref.model_type = 'DQSystem'
        return ref

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'DQSystem':
        d_q_system = DQSystem()
        if v := d.get('@type'):
            d_q_system.schema_type = v
        if v := d.get('@id'):
            d_q_system.id = v
        if v := d.get('category'):
            d_q_system.category = v
        if v := d.get('description'):
            d_q_system.description = v
        if v := d.get('hasUncertainties'):
            d_q_system.has_uncertainties = v
        if v := d.get('indicators'):
            d_q_system.indicators = [DQIndicator.from_dict(e) for e in v]
        if v := d.get('lastChange'):
            d_q_system.last_change = v
        if v := d.get('name'):
            d_q_system.name = v
        if v := d.get('source'):
            d_q_system.source = Ref.from_dict(v)
        if v := d.get('tags'):
            d_q_system.tags = v
        if v := d.get('version'):
            d_q_system.version = v
        return d_q_system

    @staticmethod
    def from_json(data: Union[str, bytes]) -> 'DQSystem':
        return DQSystem.from_dict(json.loads(data))


@dataclass
class EpdModule:

    multiplier: Optional[float] = None
    name: Optional[str] = None
    result: Optional[Ref] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.multiplier:
            d['multiplier'] = self.multiplier
        if self.name:
            d['name'] = self.name
        if self.result:
            d['result'] = self.result.to_dict()
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'EpdModule':
        epd_module = EpdModule()
        if v := d.get('@type'):
            epd_module.schema_type = v
        if v := d.get('multiplier'):
            epd_module.multiplier = v
        if v := d.get('name'):
            epd_module.name = v
        if v := d.get('result'):
            epd_module.result = Ref.from_dict(v)
        return epd_module


@dataclass
class EpdProduct:

    amount: Optional[float] = None
    flow: Optional[Ref] = None
    flow_property: Optional[Ref] = None
    unit: Optional[Ref] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.amount:
            d['amount'] = self.amount
        if self.flow:
            d['flow'] = self.flow.to_dict()
        if self.flow_property:
            d['flowProperty'] = self.flow_property.to_dict()
        if self.unit:
            d['unit'] = self.unit.to_dict()
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'EpdProduct':
        epd_product = EpdProduct()
        if v := d.get('@type'):
            epd_product.schema_type = v
        if v := d.get('amount'):
            epd_product.amount = v
        if v := d.get('flow'):
            epd_product.flow = Ref.from_dict(v)
        if v := d.get('flowProperty'):
            epd_product.flow_property = Ref.from_dict(v)
        if v := d.get('unit'):
            epd_product.unit = Ref.from_dict(v)
        return epd_product


@dataclass
class Epd:

    id: Optional[str] = None
    category: Optional[str] = None
    description: Optional[str] = None
    last_change: Optional[str] = None
    manufacturer: Optional[Ref] = None
    modules: Optional[List[EpdModule]] = None
    name: Optional[str] = None
    pcr: Optional[Ref] = None
    product: Optional[EpdProduct] = None
    program_operator: Optional[Ref] = None
    tags: Optional[List[str]] = None
    urn: Optional[str] = None
    verifier: Optional[Ref] = None
    version: Optional[str] = None

    def __post_init__(self):
        if self.id is None:
            self.id = str(uuid.uuid4())
        if self.version is None:
            self.version = '01.00.000'
        if self.last_change is None:
            self.last_change = datetime.datetime.utcnow().isoformat() + 'Z'

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        d['@type'] = 'Epd'
        if self.id:
            d['@id'] = self.id
        if self.category:
            d['category'] = self.category
        if self.description:
            d['description'] = self.description
        if self.last_change:
            d['lastChange'] = self.last_change
        if self.manufacturer:
            d['manufacturer'] = self.manufacturer.to_dict()
        if self.modules:
            d['modules'] = [e.to_dict() for e in self.modules]
        if self.name:
            d['name'] = self.name
        if self.pcr:
            d['pcr'] = self.pcr.to_dict()
        if self.product:
            d['product'] = self.product.to_dict()
        if self.program_operator:
            d['programOperator'] = self.program_operator.to_dict()
        if self.tags:
            d['tags'] = self.tags
        if self.urn:
            d['urn'] = self.urn
        if self.verifier:
            d['verifier'] = self.verifier.to_dict()
        if self.version:
            d['version'] = self.version
        return d

    def to_json(self) -> str:
        return json.dumps(self.to_dict(), indent=2)

    def to_ref(self) -> 'Ref':
        ref = Ref(id=self.id, name=self.name)
        ref.category = self.category
        ref.model_type = 'Epd'
        return ref

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'Epd':
        epd = Epd()
        if v := d.get('@type'):
            epd.schema_type = v
        if v := d.get('@id'):
            epd.id = v
        if v := d.get('category'):
            epd.category = v
        if v := d.get('description'):
            epd.description = v
        if v := d.get('lastChange'):
            epd.last_change = v
        if v := d.get('manufacturer'):
            epd.manufacturer = Ref.from_dict(v)
        if v := d.get('modules'):
            epd.modules = [EpdModule.from_dict(e) for e in v]
        if v := d.get('name'):
            epd.name = v
        if v := d.get('pcr'):
            epd.pcr = Ref.from_dict(v)
        if v := d.get('product'):
            epd.product = EpdProduct.from_dict(v)
        if v := d.get('programOperator'):
            epd.program_operator = Ref.from_dict(v)
        if v := d.get('tags'):
            epd.tags = v
        if v := d.get('urn'):
            epd.urn = v
        if v := d.get('verifier'):
            epd.verifier = Ref.from_dict(v)
        if v := d.get('version'):
            epd.version = v
        return epd

    @staticmethod
    def from_json(data: Union[str, bytes]) -> 'Epd':
        return Epd.from_dict(json.loads(data))


@dataclass
class FlowMapRef:

    flow: Optional[Ref] = None
    flow_property: Optional[Ref] = None
    provider: Optional[Ref] = None
    unit: Optional[Ref] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.flow:
            d['flow'] = self.flow.to_dict()
        if self.flow_property:
            d['flowProperty'] = self.flow_property.to_dict()
        if self.provider:
            d['provider'] = self.provider.to_dict()
        if self.unit:
            d['unit'] = self.unit.to_dict()
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'FlowMapRef':
        flow_map_ref = FlowMapRef()
        if v := d.get('@type'):
            flow_map_ref.schema_type = v
        if v := d.get('flow'):
            flow_map_ref.flow = Ref.from_dict(v)
        if v := d.get('flowProperty'):
            flow_map_ref.flow_property = Ref.from_dict(v)
        if v := d.get('provider'):
            flow_map_ref.provider = Ref.from_dict(v)
        if v := d.get('unit'):
            flow_map_ref.unit = Ref.from_dict(v)
        return flow_map_ref


@dataclass
class FlowMapEntry:

    conversion_factor: Optional[float] = None
    from_: Optional[FlowMapRef] = None
    to: Optional[FlowMapRef] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.conversion_factor:
            d['conversionFactor'] = self.conversion_factor
        if self.from_:
            d['from'] = self.from_.to_dict()
        if self.to:
            d['to'] = self.to.to_dict()
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'FlowMapEntry':
        flow_map_entry = FlowMapEntry()
        if v := d.get('@type'):
            flow_map_entry.schema_type = v
        if v := d.get('conversionFactor'):
            flow_map_entry.conversion_factor = v
        if v := d.get('from'):
            flow_map_entry.from_ = FlowMapRef.from_dict(v)
        if v := d.get('to'):
            flow_map_entry.to = FlowMapRef.from_dict(v)
        return flow_map_entry


@dataclass
class FlowMap:

    id: Optional[str] = None
    category: Optional[str] = None
    description: Optional[str] = None
    last_change: Optional[str] = None
    mappings: Optional[List[FlowMapEntry]] = None
    name: Optional[str] = None
    source: Optional[Ref] = None
    tags: Optional[List[str]] = None
    target: Optional[Ref] = None
    version: Optional[str] = None

    def __post_init__(self):
        if self.id is None:
            self.id = str(uuid.uuid4())
        if self.version is None:
            self.version = '01.00.000'
        if self.last_change is None:
            self.last_change = datetime.datetime.utcnow().isoformat() + 'Z'

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        d['@type'] = 'FlowMap'
        if self.id:
            d['@id'] = self.id
        if self.category:
            d['category'] = self.category
        if self.description:
            d['description'] = self.description
        if self.last_change:
            d['lastChange'] = self.last_change
        if self.mappings:
            d['mappings'] = [e.to_dict() for e in self.mappings]
        if self.name:
            d['name'] = self.name
        if self.source:
            d['source'] = self.source.to_dict()
        if self.tags:
            d['tags'] = self.tags
        if self.target:
            d['target'] = self.target.to_dict()
        if self.version:
            d['version'] = self.version
        return d

    def to_json(self) -> str:
        return json.dumps(self.to_dict(), indent=2)

    def to_ref(self) -> 'Ref':
        ref = Ref(id=self.id, name=self.name)
        ref.category = self.category
        ref.model_type = 'FlowMap'
        return ref

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'FlowMap':
        flow_map = FlowMap()
        if v := d.get('@type'):
            flow_map.schema_type = v
        if v := d.get('@id'):
            flow_map.id = v
        if v := d.get('category'):
            flow_map.category = v
        if v := d.get('description'):
            flow_map.description = v
        if v := d.get('lastChange'):
            flow_map.last_change = v
        if v := d.get('mappings'):
            flow_map.mappings = [FlowMapEntry.from_dict(e) for e in v]
        if v := d.get('name'):
            flow_map.name = v
        if v := d.get('source'):
            flow_map.source = Ref.from_dict(v)
        if v := d.get('tags'):
            flow_map.tags = v
        if v := d.get('target'):
            flow_map.target = Ref.from_dict(v)
        if v := d.get('version'):
            flow_map.version = v
        return flow_map

    @staticmethod
    def from_json(data: Union[str, bytes]) -> 'FlowMap':
        return FlowMap.from_dict(json.loads(data))


@dataclass
class FlowProperty:

    id: Optional[str] = None
    category: Optional[str] = None
    description: Optional[str] = None
    flow_property_type: Optional[FlowPropertyType] = None
    last_change: Optional[str] = None
    name: Optional[str] = None
    tags: Optional[List[str]] = None
    unit_group: Optional[Ref] = None
    version: Optional[str] = None

    def __post_init__(self):
        if self.id is None:
            self.id = str(uuid.uuid4())
        if self.version is None:
            self.version = '01.00.000'
        if self.last_change is None:
            self.last_change = datetime.datetime.utcnow().isoformat() + 'Z'

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        d['@type'] = 'FlowProperty'
        if self.id:
            d['@id'] = self.id
        if self.category:
            d['category'] = self.category
        if self.description:
            d['description'] = self.description
        if self.flow_property_type:
            d['flowPropertyType'] = self.flow_property_type.value
        if self.last_change:
            d['lastChange'] = self.last_change
        if self.name:
            d['name'] = self.name
        if self.tags:
            d['tags'] = self.tags
        if self.unit_group:
            d['unitGroup'] = self.unit_group.to_dict()
        if self.version:
            d['version'] = self.version
        return d

    def to_json(self) -> str:
        return json.dumps(self.to_dict(), indent=2)

    def to_ref(self) -> 'Ref':
        ref = Ref(id=self.id, name=self.name)
        ref.category = self.category
        ref.model_type = 'FlowProperty'
        return ref

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'FlowProperty':
        flow_property = FlowProperty()
        if v := d.get('@type'):
            flow_property.schema_type = v
        if v := d.get('@id'):
            flow_property.id = v
        if v := d.get('category'):
            flow_property.category = v
        if v := d.get('description'):
            flow_property.description = v
        if v := d.get('flowPropertyType'):
            flow_property.flow_property_type = FlowPropertyType.get(v)
        if v := d.get('lastChange'):
            flow_property.last_change = v
        if v := d.get('name'):
            flow_property.name = v
        if v := d.get('tags'):
            flow_property.tags = v
        if v := d.get('unitGroup'):
            flow_property.unit_group = Ref.from_dict(v)
        if v := d.get('version'):
            flow_property.version = v
        return flow_property

    @staticmethod
    def from_json(data: Union[str, bytes]) -> 'FlowProperty':
        return FlowProperty.from_dict(json.loads(data))


@dataclass
class FlowPropertyFactor:

    conversion_factor: Optional[float] = None
    flow_property: Optional[Ref] = None
    is_ref_flow_property: Optional[bool] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.conversion_factor:
            d['conversionFactor'] = self.conversion_factor
        if self.flow_property:
            d['flowProperty'] = self.flow_property.to_dict()
        if self.is_ref_flow_property:
            d['isRefFlowProperty'] = self.is_ref_flow_property
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'FlowPropertyFactor':
        flow_property_factor = FlowPropertyFactor()
        if v := d.get('@type'):
            flow_property_factor.schema_type = v
        if v := d.get('conversionFactor'):
            flow_property_factor.conversion_factor = v
        if v := d.get('flowProperty'):
            flow_property_factor.flow_property = Ref.from_dict(v)
        if v := d.get('isRefFlowProperty'):
            flow_property_factor.is_ref_flow_property = v
        return flow_property_factor


@dataclass
class Flow:

    id: Optional[str] = None
    cas: Optional[str] = None
    category: Optional[str] = None
    description: Optional[str] = None
    flow_properties: Optional[List[FlowPropertyFactor]] = None
    flow_type: Optional[FlowType] = None
    formula: Optional[str] = None
    is_infrastructure_flow: Optional[bool] = None
    last_change: Optional[str] = None
    location: Optional[Ref] = None
    name: Optional[str] = None
    synonyms: Optional[str] = None
    tags: Optional[List[str]] = None
    version: Optional[str] = None

    def __post_init__(self):
        if self.id is None:
            self.id = str(uuid.uuid4())
        if self.version is None:
            self.version = '01.00.000'
        if self.last_change is None:
            self.last_change = datetime.datetime.utcnow().isoformat() + 'Z'

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        d['@type'] = 'Flow'
        if self.id:
            d['@id'] = self.id
        if self.cas:
            d['cas'] = self.cas
        if self.category:
            d['category'] = self.category
        if self.description:
            d['description'] = self.description
        if self.flow_properties:
            d['flowProperties'] = [e.to_dict() for e in self.flow_properties]
        if self.flow_type:
            d['flowType'] = self.flow_type.value
        if self.formula:
            d['formula'] = self.formula
        if self.is_infrastructure_flow:
            d['isInfrastructureFlow'] = self.is_infrastructure_flow
        if self.last_change:
            d['lastChange'] = self.last_change
        if self.location:
            d['location'] = self.location.to_dict()
        if self.name:
            d['name'] = self.name
        if self.synonyms:
            d['synonyms'] = self.synonyms
        if self.tags:
            d['tags'] = self.tags
        if self.version:
            d['version'] = self.version
        return d

    def to_json(self) -> str:
        return json.dumps(self.to_dict(), indent=2)

    def to_ref(self) -> 'Ref':
        ref = Ref(id=self.id, name=self.name)
        ref.category = self.category
        ref.model_type = 'Flow'
        return ref

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'Flow':
        flow = Flow()
        if v := d.get('@type'):
            flow.schema_type = v
        if v := d.get('@id'):
            flow.id = v
        if v := d.get('cas'):
            flow.cas = v
        if v := d.get('category'):
            flow.category = v
        if v := d.get('description'):
            flow.description = v
        if v := d.get('flowProperties'):
            flow.flow_properties = [FlowPropertyFactor.from_dict(e) for e in v]
        if v := d.get('flowType'):
            flow.flow_type = FlowType.get(v)
        if v := d.get('formula'):
            flow.formula = v
        if v := d.get('isInfrastructureFlow'):
            flow.is_infrastructure_flow = v
        if v := d.get('lastChange'):
            flow.last_change = v
        if v := d.get('location'):
            flow.location = Ref.from_dict(v)
        if v := d.get('name'):
            flow.name = v
        if v := d.get('synonyms'):
            flow.synonyms = v
        if v := d.get('tags'):
            flow.tags = v
        if v := d.get('version'):
            flow.version = v
        return flow

    @staticmethod
    def from_json(data: Union[str, bytes]) -> 'Flow':
        return Flow.from_dict(json.loads(data))


@dataclass
class FlowResult:

    amount: Optional[float] = None
    description: Optional[str] = None
    flow: Optional[Ref] = None
    flow_property: Optional[Ref] = None
    is_input: Optional[bool] = None
    is_ref_flow: Optional[bool] = None
    location: Optional[Ref] = None
    unit: Optional[Ref] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.amount:
            d['amount'] = self.amount
        if self.description:
            d['description'] = self.description
        if self.flow:
            d['flow'] = self.flow.to_dict()
        if self.flow_property:
            d['flowProperty'] = self.flow_property.to_dict()
        if self.is_input:
            d['isInput'] = self.is_input
        if self.is_ref_flow:
            d['isRefFlow'] = self.is_ref_flow
        if self.location:
            d['location'] = self.location.to_dict()
        if self.unit:
            d['unit'] = self.unit.to_dict()
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'FlowResult':
        flow_result = FlowResult()
        if v := d.get('@type'):
            flow_result.schema_type = v
        if v := d.get('amount'):
            flow_result.amount = v
        if v := d.get('description'):
            flow_result.description = v
        if v := d.get('flow'):
            flow_result.flow = Ref.from_dict(v)
        if v := d.get('flowProperty'):
            flow_result.flow_property = Ref.from_dict(v)
        if v := d.get('isInput'):
            flow_result.is_input = v
        if v := d.get('isRefFlow'):
            flow_result.is_ref_flow = v
        if v := d.get('location'):
            flow_result.location = Ref.from_dict(v)
        if v := d.get('unit'):
            flow_result.unit = Ref.from_dict(v)
        return flow_result


@dataclass
class ImpactResult:

    amount: Optional[float] = None
    description: Optional[str] = None
    indicator: Optional[Ref] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.amount:
            d['amount'] = self.amount
        if self.description:
            d['description'] = self.description
        if self.indicator:
            d['indicator'] = self.indicator.to_dict()
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'ImpactResult':
        impact_result = ImpactResult()
        if v := d.get('@type'):
            impact_result.schema_type = v
        if v := d.get('amount'):
            impact_result.amount = v
        if v := d.get('description'):
            impact_result.description = v
        if v := d.get('indicator'):
            impact_result.indicator = Ref.from_dict(v)
        return impact_result


@dataclass
class Location:

    id: Optional[str] = None
    category: Optional[str] = None
    code: Optional[str] = None
    description: Optional[str] = None
    geometry: Optional[Dict[str, Any]] = None
    last_change: Optional[str] = None
    latitude: Optional[float] = None
    longitude: Optional[float] = None
    name: Optional[str] = None
    tags: Optional[List[str]] = None
    version: Optional[str] = None

    def __post_init__(self):
        if self.id is None:
            self.id = str(uuid.uuid4())
        if self.version is None:
            self.version = '01.00.000'
        if self.last_change is None:
            self.last_change = datetime.datetime.utcnow().isoformat() + 'Z'

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        d['@type'] = 'Location'
        if self.id:
            d['@id'] = self.id
        if self.category:
            d['category'] = self.category
        if self.code:
            d['code'] = self.code
        if self.description:
            d['description'] = self.description
        if self.geometry:
            d['geometry'] = self.geometry
        if self.last_change:
            d['lastChange'] = self.last_change
        if self.latitude:
            d['latitude'] = self.latitude
        if self.longitude:
            d['longitude'] = self.longitude
        if self.name:
            d['name'] = self.name
        if self.tags:
            d['tags'] = self.tags
        if self.version:
            d['version'] = self.version
        return d

    def to_json(self) -> str:
        return json.dumps(self.to_dict(), indent=2)

    def to_ref(self) -> 'Ref':
        ref = Ref(id=self.id, name=self.name)
        ref.category = self.category
        ref.model_type = 'Location'
        return ref

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'Location':
        location = Location()
        if v := d.get('@type'):
            location.schema_type = v
        if v := d.get('@id'):
            location.id = v
        if v := d.get('category'):
            location.category = v
        if v := d.get('code'):
            location.code = v
        if v := d.get('description'):
            location.description = v
        if v := d.get('geometry'):
            location.geometry = v
        if v := d.get('lastChange'):
            location.last_change = v
        if v := d.get('latitude'):
            location.latitude = v
        if v := d.get('longitude'):
            location.longitude = v
        if v := d.get('name'):
            location.name = v
        if v := d.get('tags'):
            location.tags = v
        if v := d.get('version'):
            location.version = v
        return location

    @staticmethod
    def from_json(data: Union[str, bytes]) -> 'Location':
        return Location.from_dict(json.loads(data))


@dataclass
class NwFactor:

    impact_category: Optional[Ref] = None
    normalisation_factor: Optional[float] = None
    weighting_factor: Optional[float] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.impact_category:
            d['impactCategory'] = self.impact_category.to_dict()
        if self.normalisation_factor:
            d['normalisationFactor'] = self.normalisation_factor
        if self.weighting_factor:
            d['weightingFactor'] = self.weighting_factor
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'NwFactor':
        nw_factor = NwFactor()
        if v := d.get('@type'):
            nw_factor.schema_type = v
        if v := d.get('impactCategory'):
            nw_factor.impact_category = Ref.from_dict(v)
        if v := d.get('normalisationFactor'):
            nw_factor.normalisation_factor = v
        if v := d.get('weightingFactor'):
            nw_factor.weighting_factor = v
        return nw_factor


@dataclass
class NwSet:

    id: Optional[str] = None
    description: Optional[str] = None
    factors: Optional[List[NwFactor]] = None
    name: Optional[str] = None
    weighted_score_unit: Optional[str] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.id:
            d['@id'] = self.id
        if self.description:
            d['description'] = self.description
        if self.factors:
            d['factors'] = [e.to_dict() for e in self.factors]
        if self.name:
            d['name'] = self.name
        if self.weighted_score_unit:
            d['weightedScoreUnit'] = self.weighted_score_unit
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'NwSet':
        nw_set = NwSet()
        if v := d.get('@type'):
            nw_set.schema_type = v
        if v := d.get('@id'):
            nw_set.id = v
        if v := d.get('description'):
            nw_set.description = v
        if v := d.get('factors'):
            nw_set.factors = [NwFactor.from_dict(e) for e in v]
        if v := d.get('name'):
            nw_set.name = v
        if v := d.get('weightedScoreUnit'):
            nw_set.weighted_score_unit = v
        return nw_set


@dataclass
class ImpactMethod:

    id: Optional[str] = None
    category: Optional[str] = None
    code: Optional[str] = None
    description: Optional[str] = None
    impact_categories: Optional[List[Ref]] = None
    last_change: Optional[str] = None
    name: Optional[str] = None
    nw_sets: Optional[List[NwSet]] = None
    source: Optional[Ref] = None
    tags: Optional[List[str]] = None
    version: Optional[str] = None

    def __post_init__(self):
        if self.id is None:
            self.id = str(uuid.uuid4())
        if self.version is None:
            self.version = '01.00.000'
        if self.last_change is None:
            self.last_change = datetime.datetime.utcnow().isoformat() + 'Z'

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        d['@type'] = 'ImpactMethod'
        if self.id:
            d['@id'] = self.id
        if self.category:
            d['category'] = self.category
        if self.code:
            d['code'] = self.code
        if self.description:
            d['description'] = self.description
        if self.impact_categories:
            d['impactCategories'] = [e.to_dict() for e in self.impact_categories]
        if self.last_change:
            d['lastChange'] = self.last_change
        if self.name:
            d['name'] = self.name
        if self.nw_sets:
            d['nwSets'] = [e.to_dict() for e in self.nw_sets]
        if self.source:
            d['source'] = self.source.to_dict()
        if self.tags:
            d['tags'] = self.tags
        if self.version:
            d['version'] = self.version
        return d

    def to_json(self) -> str:
        return json.dumps(self.to_dict(), indent=2)

    def to_ref(self) -> 'Ref':
        ref = Ref(id=self.id, name=self.name)
        ref.category = self.category
        ref.model_type = 'ImpactMethod'
        return ref

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'ImpactMethod':
        impact_method = ImpactMethod()
        if v := d.get('@type'):
            impact_method.schema_type = v
        if v := d.get('@id'):
            impact_method.id = v
        if v := d.get('category'):
            impact_method.category = v
        if v := d.get('code'):
            impact_method.code = v
        if v := d.get('description'):
            impact_method.description = v
        if v := d.get('impactCategories'):
            impact_method.impact_categories = [Ref.from_dict(e) for e in v]
        if v := d.get('lastChange'):
            impact_method.last_change = v
        if v := d.get('name'):
            impact_method.name = v
        if v := d.get('nwSets'):
            impact_method.nw_sets = [NwSet.from_dict(e) for e in v]
        if v := d.get('source'):
            impact_method.source = Ref.from_dict(v)
        if v := d.get('tags'):
            impact_method.tags = v
        if v := d.get('version'):
            impact_method.version = v
        return impact_method

    @staticmethod
    def from_json(data: Union[str, bytes]) -> 'ImpactMethod':
        return ImpactMethod.from_dict(json.loads(data))


@dataclass
class ProcessDocumentation:

    completeness_description: Optional[str] = None
    creation_date: Optional[str] = None
    data_collection_description: Optional[str] = None
    data_documentor: Optional[Ref] = None
    data_generator: Optional[Ref] = None
    data_selection_description: Optional[str] = None
    data_set_owner: Optional[Ref] = None
    data_treatment_description: Optional[str] = None
    geography_description: Optional[str] = None
    intended_application: Optional[str] = None
    inventory_method_description: Optional[str] = None
    is_copyright_protected: Optional[bool] = None
    modeling_constants_description: Optional[str] = None
    project_description: Optional[str] = None
    publication: Optional[Ref] = None
    restrictions_description: Optional[str] = None
    review_details: Optional[str] = None
    reviewer: Optional[Ref] = None
    sampling_description: Optional[str] = None
    sources: Optional[List[Ref]] = None
    technology_description: Optional[str] = None
    time_description: Optional[str] = None
    valid_from: Optional[str] = None
    valid_until: Optional[str] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.completeness_description:
            d['completenessDescription'] = self.completeness_description
        if self.creation_date:
            d['creationDate'] = self.creation_date
        if self.data_collection_description:
            d['dataCollectionDescription'] = self.data_collection_description
        if self.data_documentor:
            d['dataDocumentor'] = self.data_documentor.to_dict()
        if self.data_generator:
            d['dataGenerator'] = self.data_generator.to_dict()
        if self.data_selection_description:
            d['dataSelectionDescription'] = self.data_selection_description
        if self.data_set_owner:
            d['dataSetOwner'] = self.data_set_owner.to_dict()
        if self.data_treatment_description:
            d['dataTreatmentDescription'] = self.data_treatment_description
        if self.geography_description:
            d['geographyDescription'] = self.geography_description
        if self.intended_application:
            d['intendedApplication'] = self.intended_application
        if self.inventory_method_description:
            d['inventoryMethodDescription'] = self.inventory_method_description
        if self.is_copyright_protected:
            d['isCopyrightProtected'] = self.is_copyright_protected
        if self.modeling_constants_description:
            d['modelingConstantsDescription'] = self.modeling_constants_description
        if self.project_description:
            d['projectDescription'] = self.project_description
        if self.publication:
            d['publication'] = self.publication.to_dict()
        if self.restrictions_description:
            d['restrictionsDescription'] = self.restrictions_description
        if self.review_details:
            d['reviewDetails'] = self.review_details
        if self.reviewer:
            d['reviewer'] = self.reviewer.to_dict()
        if self.sampling_description:
            d['samplingDescription'] = self.sampling_description
        if self.sources:
            d['sources'] = [e.to_dict() for e in self.sources]
        if self.technology_description:
            d['technologyDescription'] = self.technology_description
        if self.time_description:
            d['timeDescription'] = self.time_description
        if self.valid_from:
            d['validFrom'] = self.valid_from
        if self.valid_until:
            d['validUntil'] = self.valid_until
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'ProcessDocumentation':
        process_documentation = ProcessDocumentation()
        if v := d.get('@type'):
            process_documentation.schema_type = v
        if v := d.get('completenessDescription'):
            process_documentation.completeness_description = v
        if v := d.get('creationDate'):
            process_documentation.creation_date = v
        if v := d.get('dataCollectionDescription'):
            process_documentation.data_collection_description = v
        if v := d.get('dataDocumentor'):
            process_documentation.data_documentor = Ref.from_dict(v)
        if v := d.get('dataGenerator'):
            process_documentation.data_generator = Ref.from_dict(v)
        if v := d.get('dataSelectionDescription'):
            process_documentation.data_selection_description = v
        if v := d.get('dataSetOwner'):
            process_documentation.data_set_owner = Ref.from_dict(v)
        if v := d.get('dataTreatmentDescription'):
            process_documentation.data_treatment_description = v
        if v := d.get('geographyDescription'):
            process_documentation.geography_description = v
        if v := d.get('intendedApplication'):
            process_documentation.intended_application = v
        if v := d.get('inventoryMethodDescription'):
            process_documentation.inventory_method_description = v
        if v := d.get('isCopyrightProtected'):
            process_documentation.is_copyright_protected = v
        if v := d.get('modelingConstantsDescription'):
            process_documentation.modeling_constants_description = v
        if v := d.get('projectDescription'):
            process_documentation.project_description = v
        if v := d.get('publication'):
            process_documentation.publication = Ref.from_dict(v)
        if v := d.get('restrictionsDescription'):
            process_documentation.restrictions_description = v
        if v := d.get('reviewDetails'):
            process_documentation.review_details = v
        if v := d.get('reviewer'):
            process_documentation.reviewer = Ref.from_dict(v)
        if v := d.get('samplingDescription'):
            process_documentation.sampling_description = v
        if v := d.get('sources'):
            process_documentation.sources = [Ref.from_dict(e) for e in v]
        if v := d.get('technologyDescription'):
            process_documentation.technology_description = v
        if v := d.get('timeDescription'):
            process_documentation.time_description = v
        if v := d.get('validFrom'):
            process_documentation.valid_from = v
        if v := d.get('validUntil'):
            process_documentation.valid_until = v
        return process_documentation


@dataclass
class ProcessLink:

    exchange: Optional[ExchangeRef] = None
    flow: Optional[Ref] = None
    process: Optional[Ref] = None
    provider: Optional[Ref] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.exchange:
            d['exchange'] = self.exchange.to_dict()
        if self.flow:
            d['flow'] = self.flow.to_dict()
        if self.process:
            d['process'] = self.process.to_dict()
        if self.provider:
            d['provider'] = self.provider.to_dict()
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'ProcessLink':
        process_link = ProcessLink()
        if v := d.get('@type'):
            process_link.schema_type = v
        if v := d.get('exchange'):
            process_link.exchange = ExchangeRef.from_dict(v)
        if v := d.get('flow'):
            process_link.flow = Ref.from_dict(v)
        if v := d.get('process'):
            process_link.process = Ref.from_dict(v)
        if v := d.get('provider'):
            process_link.provider = Ref.from_dict(v)
        return process_link


@dataclass
class Result:

    id: Optional[str] = None
    category: Optional[str] = None
    description: Optional[str] = None
    flow_results: Optional[List[FlowResult]] = None
    impact_method: Optional[Ref] = None
    impact_results: Optional[List[ImpactResult]] = None
    last_change: Optional[str] = None
    name: Optional[str] = None
    product_system: Optional[Ref] = None
    tags: Optional[List[str]] = None
    version: Optional[str] = None

    def __post_init__(self):
        if self.id is None:
            self.id = str(uuid.uuid4())
        if self.version is None:
            self.version = '01.00.000'
        if self.last_change is None:
            self.last_change = datetime.datetime.utcnow().isoformat() + 'Z'

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        d['@type'] = 'Result'
        if self.id:
            d['@id'] = self.id
        if self.category:
            d['category'] = self.category
        if self.description:
            d['description'] = self.description
        if self.flow_results:
            d['flowResults'] = [e.to_dict() for e in self.flow_results]
        if self.impact_method:
            d['impactMethod'] = self.impact_method.to_dict()
        if self.impact_results:
            d['impactResults'] = [e.to_dict() for e in self.impact_results]
        if self.last_change:
            d['lastChange'] = self.last_change
        if self.name:
            d['name'] = self.name
        if self.product_system:
            d['productSystem'] = self.product_system.to_dict()
        if self.tags:
            d['tags'] = self.tags
        if self.version:
            d['version'] = self.version
        return d

    def to_json(self) -> str:
        return json.dumps(self.to_dict(), indent=2)

    def to_ref(self) -> 'Ref':
        ref = Ref(id=self.id, name=self.name)
        ref.category = self.category
        ref.model_type = 'Result'
        return ref

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'Result':
        result = Result()
        if v := d.get('@type'):
            result.schema_type = v
        if v := d.get('@id'):
            result.id = v
        if v := d.get('category'):
            result.category = v
        if v := d.get('description'):
            result.description = v
        if v := d.get('flowResults'):
            result.flow_results = [FlowResult.from_dict(e) for e in v]
        if v := d.get('impactMethod'):
            result.impact_method = Ref.from_dict(v)
        if v := d.get('impactResults'):
            result.impact_results = [ImpactResult.from_dict(e) for e in v]
        if v := d.get('lastChange'):
            result.last_change = v
        if v := d.get('name'):
            result.name = v
        if v := d.get('productSystem'):
            result.product_system = Ref.from_dict(v)
        if v := d.get('tags'):
            result.tags = v
        if v := d.get('version'):
            result.version = v
        return result

    @staticmethod
    def from_json(data: Union[str, bytes]) -> 'Result':
        return Result.from_dict(json.loads(data))


@dataclass
class SocialAspect:

    activity_value: Optional[float] = None
    comment: Optional[str] = None
    quality: Optional[str] = None
    raw_amount: Optional[str] = None
    risk_level: Optional[RiskLevel] = None
    social_indicator: Optional[Ref] = None
    source: Optional[Ref] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.activity_value:
            d['activityValue'] = self.activity_value
        if self.comment:
            d['comment'] = self.comment
        if self.quality:
            d['quality'] = self.quality
        if self.raw_amount:
            d['rawAmount'] = self.raw_amount
        if self.risk_level:
            d['riskLevel'] = self.risk_level.value
        if self.social_indicator:
            d['socialIndicator'] = self.social_indicator.to_dict()
        if self.source:
            d['source'] = self.source.to_dict()
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'SocialAspect':
        social_aspect = SocialAspect()
        if v := d.get('@type'):
            social_aspect.schema_type = v
        if v := d.get('activityValue'):
            social_aspect.activity_value = v
        if v := d.get('comment'):
            social_aspect.comment = v
        if v := d.get('quality'):
            social_aspect.quality = v
        if v := d.get('rawAmount'):
            social_aspect.raw_amount = v
        if v := d.get('riskLevel'):
            social_aspect.risk_level = RiskLevel.get(v)
        if v := d.get('socialIndicator'):
            social_aspect.social_indicator = Ref.from_dict(v)
        if v := d.get('source'):
            social_aspect.source = Ref.from_dict(v)
        return social_aspect


@dataclass
class SocialIndicator:

    id: Optional[str] = None
    activity_quantity: Optional[Ref] = None
    activity_unit: Optional[Ref] = None
    activity_variable: Optional[str] = None
    category: Optional[str] = None
    description: Optional[str] = None
    evaluation_scheme: Optional[str] = None
    last_change: Optional[str] = None
    name: Optional[str] = None
    tags: Optional[List[str]] = None
    unit_of_measurement: Optional[str] = None
    version: Optional[str] = None

    def __post_init__(self):
        if self.id is None:
            self.id = str(uuid.uuid4())
        if self.version is None:
            self.version = '01.00.000'
        if self.last_change is None:
            self.last_change = datetime.datetime.utcnow().isoformat() + 'Z'

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        d['@type'] = 'SocialIndicator'
        if self.id:
            d['@id'] = self.id
        if self.activity_quantity:
            d['activityQuantity'] = self.activity_quantity.to_dict()
        if self.activity_unit:
            d['activityUnit'] = self.activity_unit.to_dict()
        if self.activity_variable:
            d['activityVariable'] = self.activity_variable
        if self.category:
            d['category'] = self.category
        if self.description:
            d['description'] = self.description
        if self.evaluation_scheme:
            d['evaluationScheme'] = self.evaluation_scheme
        if self.last_change:
            d['lastChange'] = self.last_change
        if self.name:
            d['name'] = self.name
        if self.tags:
            d['tags'] = self.tags
        if self.unit_of_measurement:
            d['unitOfMeasurement'] = self.unit_of_measurement
        if self.version:
            d['version'] = self.version
        return d

    def to_json(self) -> str:
        return json.dumps(self.to_dict(), indent=2)

    def to_ref(self) -> 'Ref':
        ref = Ref(id=self.id, name=self.name)
        ref.category = self.category
        ref.model_type = 'SocialIndicator'
        return ref

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'SocialIndicator':
        social_indicator = SocialIndicator()
        if v := d.get('@type'):
            social_indicator.schema_type = v
        if v := d.get('@id'):
            social_indicator.id = v
        if v := d.get('activityQuantity'):
            social_indicator.activity_quantity = Ref.from_dict(v)
        if v := d.get('activityUnit'):
            social_indicator.activity_unit = Ref.from_dict(v)
        if v := d.get('activityVariable'):
            social_indicator.activity_variable = v
        if v := d.get('category'):
            social_indicator.category = v
        if v := d.get('description'):
            social_indicator.description = v
        if v := d.get('evaluationScheme'):
            social_indicator.evaluation_scheme = v
        if v := d.get('lastChange'):
            social_indicator.last_change = v
        if v := d.get('name'):
            social_indicator.name = v
        if v := d.get('tags'):
            social_indicator.tags = v
        if v := d.get('unitOfMeasurement'):
            social_indicator.unit_of_measurement = v
        if v := d.get('version'):
            social_indicator.version = v
        return social_indicator

    @staticmethod
    def from_json(data: Union[str, bytes]) -> 'SocialIndicator':
        return SocialIndicator.from_dict(json.loads(data))


@dataclass
class Source:

    id: Optional[str] = None
    category: Optional[str] = None
    description: Optional[str] = None
    external_file: Optional[str] = None
    last_change: Optional[str] = None
    name: Optional[str] = None
    tags: Optional[List[str]] = None
    text_reference: Optional[str] = None
    url: Optional[str] = None
    version: Optional[str] = None
    year: Optional[int] = None

    def __post_init__(self):
        if self.id is None:
            self.id = str(uuid.uuid4())
        if self.version is None:
            self.version = '01.00.000'
        if self.last_change is None:
            self.last_change = datetime.datetime.utcnow().isoformat() + 'Z'

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        d['@type'] = 'Source'
        if self.id:
            d['@id'] = self.id
        if self.category:
            d['category'] = self.category
        if self.description:
            d['description'] = self.description
        if self.external_file:
            d['externalFile'] = self.external_file
        if self.last_change:
            d['lastChange'] = self.last_change
        if self.name:
            d['name'] = self.name
        if self.tags:
            d['tags'] = self.tags
        if self.text_reference:
            d['textReference'] = self.text_reference
        if self.url:
            d['url'] = self.url
        if self.version:
            d['version'] = self.version
        if self.year:
            d['year'] = self.year
        return d

    def to_json(self) -> str:
        return json.dumps(self.to_dict(), indent=2)

    def to_ref(self) -> 'Ref':
        ref = Ref(id=self.id, name=self.name)
        ref.category = self.category
        ref.model_type = 'Source'
        return ref

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'Source':
        source = Source()
        if v := d.get('@type'):
            source.schema_type = v
        if v := d.get('@id'):
            source.id = v
        if v := d.get('category'):
            source.category = v
        if v := d.get('description'):
            source.description = v
        if v := d.get('externalFile'):
            source.external_file = v
        if v := d.get('lastChange'):
            source.last_change = v
        if v := d.get('name'):
            source.name = v
        if v := d.get('tags'):
            source.tags = v
        if v := d.get('textReference'):
            source.text_reference = v
        if v := d.get('url'):
            source.url = v
        if v := d.get('version'):
            source.version = v
        if v := d.get('year'):
            source.year = v
        return source

    @staticmethod
    def from_json(data: Union[str, bytes]) -> 'Source':
        return Source.from_dict(json.loads(data))


@dataclass
class Uncertainty:

    distribution_type: Optional[UncertaintyType] = None
    geom_mean: Optional[float] = None
    geom_sd: Optional[float] = None
    maximum: Optional[float] = None
    mean: Optional[float] = None
    minimum: Optional[float] = None
    mode: Optional[float] = None
    sd: Optional[float] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.distribution_type:
            d['distributionType'] = self.distribution_type.value
        if self.geom_mean:
            d['geomMean'] = self.geom_mean
        if self.geom_sd:
            d['geomSd'] = self.geom_sd
        if self.maximum:
            d['maximum'] = self.maximum
        if self.mean:
            d['mean'] = self.mean
        if self.minimum:
            d['minimum'] = self.minimum
        if self.mode:
            d['mode'] = self.mode
        if self.sd:
            d['sd'] = self.sd
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'Uncertainty':
        uncertainty = Uncertainty()
        if v := d.get('@type'):
            uncertainty.schema_type = v
        if v := d.get('distributionType'):
            uncertainty.distribution_type = UncertaintyType.get(v)
        if v := d.get('geomMean'):
            uncertainty.geom_mean = v
        if v := d.get('geomSd'):
            uncertainty.geom_sd = v
        if v := d.get('maximum'):
            uncertainty.maximum = v
        if v := d.get('mean'):
            uncertainty.mean = v
        if v := d.get('minimum'):
            uncertainty.minimum = v
        if v := d.get('mode'):
            uncertainty.mode = v
        if v := d.get('sd'):
            uncertainty.sd = v
        return uncertainty


@dataclass
class Exchange:

    amount: Optional[float] = None
    amount_formula: Optional[str] = None
    base_uncertainty: Optional[float] = None
    cost_formula: Optional[str] = None
    cost_value: Optional[float] = None
    currency: Optional[Ref] = None
    default_provider: Optional[Ref] = None
    description: Optional[str] = None
    dq_entry: Optional[str] = None
    flow: Optional[Ref] = None
    flow_property: Optional[Ref] = None
    internal_id: Optional[int] = None
    is_avoided_product: Optional[bool] = None
    is_input: Optional[bool] = None
    is_quantitative_reference: Optional[bool] = None
    location: Optional[Ref] = None
    uncertainty: Optional[Uncertainty] = None
    unit: Optional[Ref] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.amount:
            d['amount'] = self.amount
        if self.amount_formula:
            d['amountFormula'] = self.amount_formula
        if self.base_uncertainty:
            d['baseUncertainty'] = self.base_uncertainty
        if self.cost_formula:
            d['costFormula'] = self.cost_formula
        if self.cost_value:
            d['costValue'] = self.cost_value
        if self.currency:
            d['currency'] = self.currency.to_dict()
        if self.default_provider:
            d['defaultProvider'] = self.default_provider.to_dict()
        if self.description:
            d['description'] = self.description
        if self.dq_entry:
            d['dqEntry'] = self.dq_entry
        if self.flow:
            d['flow'] = self.flow.to_dict()
        if self.flow_property:
            d['flowProperty'] = self.flow_property.to_dict()
        if self.internal_id:
            d['internalId'] = self.internal_id
        if self.is_avoided_product:
            d['isAvoidedProduct'] = self.is_avoided_product
        if self.is_input:
            d['isInput'] = self.is_input
        if self.is_quantitative_reference:
            d['isQuantitativeReference'] = self.is_quantitative_reference
        if self.location:
            d['location'] = self.location.to_dict()
        if self.uncertainty:
            d['uncertainty'] = self.uncertainty.to_dict()
        if self.unit:
            d['unit'] = self.unit.to_dict()
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'Exchange':
        exchange = Exchange()
        if v := d.get('@type'):
            exchange.schema_type = v
        if v := d.get('amount'):
            exchange.amount = v
        if v := d.get('amountFormula'):
            exchange.amount_formula = v
        if v := d.get('baseUncertainty'):
            exchange.base_uncertainty = v
        if v := d.get('costFormula'):
            exchange.cost_formula = v
        if v := d.get('costValue'):
            exchange.cost_value = v
        if v := d.get('currency'):
            exchange.currency = Ref.from_dict(v)
        if v := d.get('defaultProvider'):
            exchange.default_provider = Ref.from_dict(v)
        if v := d.get('description'):
            exchange.description = v
        if v := d.get('dqEntry'):
            exchange.dq_entry = v
        if v := d.get('flow'):
            exchange.flow = Ref.from_dict(v)
        if v := d.get('flowProperty'):
            exchange.flow_property = Ref.from_dict(v)
        if v := d.get('internalId'):
            exchange.internal_id = v
        if v := d.get('isAvoidedProduct'):
            exchange.is_avoided_product = v
        if v := d.get('isInput'):
            exchange.is_input = v
        if v := d.get('isQuantitativeReference'):
            exchange.is_quantitative_reference = v
        if v := d.get('location'):
            exchange.location = Ref.from_dict(v)
        if v := d.get('uncertainty'):
            exchange.uncertainty = Uncertainty.from_dict(v)
        if v := d.get('unit'):
            exchange.unit = Ref.from_dict(v)
        return exchange


@dataclass
class ImpactFactor:

    flow: Optional[Ref] = None
    flow_property: Optional[Ref] = None
    formula: Optional[str] = None
    location: Optional[Ref] = None
    uncertainty: Optional[Uncertainty] = None
    unit: Optional[Ref] = None
    value: Optional[float] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.flow:
            d['flow'] = self.flow.to_dict()
        if self.flow_property:
            d['flowProperty'] = self.flow_property.to_dict()
        if self.formula:
            d['formula'] = self.formula
        if self.location:
            d['location'] = self.location.to_dict()
        if self.uncertainty:
            d['uncertainty'] = self.uncertainty.to_dict()
        if self.unit:
            d['unit'] = self.unit.to_dict()
        if self.value:
            d['value'] = self.value
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'ImpactFactor':
        impact_factor = ImpactFactor()
        if v := d.get('@type'):
            impact_factor.schema_type = v
        if v := d.get('flow'):
            impact_factor.flow = Ref.from_dict(v)
        if v := d.get('flowProperty'):
            impact_factor.flow_property = Ref.from_dict(v)
        if v := d.get('formula'):
            impact_factor.formula = v
        if v := d.get('location'):
            impact_factor.location = Ref.from_dict(v)
        if v := d.get('uncertainty'):
            impact_factor.uncertainty = Uncertainty.from_dict(v)
        if v := d.get('unit'):
            impact_factor.unit = Ref.from_dict(v)
        if v := d.get('value'):
            impact_factor.value = v
        return impact_factor


@dataclass
class Parameter:

    id: Optional[str] = None
    category: Optional[str] = None
    description: Optional[str] = None
    formula: Optional[str] = None
    is_input_parameter: Optional[bool] = None
    last_change: Optional[str] = None
    name: Optional[str] = None
    parameter_scope: Optional[ParameterScope] = None
    tags: Optional[List[str]] = None
    uncertainty: Optional[Uncertainty] = None
    value: Optional[float] = None
    version: Optional[str] = None

    def __post_init__(self):
        if self.id is None:
            self.id = str(uuid.uuid4())
        if self.version is None:
            self.version = '01.00.000'
        if self.last_change is None:
            self.last_change = datetime.datetime.utcnow().isoformat() + 'Z'

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        d['@type'] = 'Parameter'
        if self.id:
            d['@id'] = self.id
        if self.category:
            d['category'] = self.category
        if self.description:
            d['description'] = self.description
        if self.formula:
            d['formula'] = self.formula
        if self.is_input_parameter:
            d['isInputParameter'] = self.is_input_parameter
        if self.last_change:
            d['lastChange'] = self.last_change
        if self.name:
            d['name'] = self.name
        if self.parameter_scope:
            d['parameterScope'] = self.parameter_scope.value
        if self.tags:
            d['tags'] = self.tags
        if self.uncertainty:
            d['uncertainty'] = self.uncertainty.to_dict()
        if self.value:
            d['value'] = self.value
        if self.version:
            d['version'] = self.version
        return d

    def to_json(self) -> str:
        return json.dumps(self.to_dict(), indent=2)

    def to_ref(self) -> 'Ref':
        ref = Ref(id=self.id, name=self.name)
        ref.category = self.category
        ref.model_type = 'Parameter'
        return ref

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'Parameter':
        parameter = Parameter()
        if v := d.get('@type'):
            parameter.schema_type = v
        if v := d.get('@id'):
            parameter.id = v
        if v := d.get('category'):
            parameter.category = v
        if v := d.get('description'):
            parameter.description = v
        if v := d.get('formula'):
            parameter.formula = v
        if v := d.get('isInputParameter'):
            parameter.is_input_parameter = v
        if v := d.get('lastChange'):
            parameter.last_change = v
        if v := d.get('name'):
            parameter.name = v
        if v := d.get('parameterScope'):
            parameter.parameter_scope = ParameterScope.get(v)
        if v := d.get('tags'):
            parameter.tags = v
        if v := d.get('uncertainty'):
            parameter.uncertainty = Uncertainty.from_dict(v)
        if v := d.get('value'):
            parameter.value = v
        if v := d.get('version'):
            parameter.version = v
        return parameter

    @staticmethod
    def from_json(data: Union[str, bytes]) -> 'Parameter':
        return Parameter.from_dict(json.loads(data))


@dataclass
class ImpactCategory:

    id: Optional[str] = None
    category: Optional[str] = None
    code: Optional[str] = None
    description: Optional[str] = None
    direction: Optional[Direction] = None
    impact_factors: Optional[List[ImpactFactor]] = None
    last_change: Optional[str] = None
    name: Optional[str] = None
    parameters: Optional[List[Parameter]] = None
    ref_unit: Optional[str] = None
    source: Optional[Ref] = None
    tags: Optional[List[str]] = None
    version: Optional[str] = None

    def __post_init__(self):
        if self.id is None:
            self.id = str(uuid.uuid4())
        if self.version is None:
            self.version = '01.00.000'
        if self.last_change is None:
            self.last_change = datetime.datetime.utcnow().isoformat() + 'Z'

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        d['@type'] = 'ImpactCategory'
        if self.id:
            d['@id'] = self.id
        if self.category:
            d['category'] = self.category
        if self.code:
            d['code'] = self.code
        if self.description:
            d['description'] = self.description
        if self.direction:
            d['direction'] = self.direction.value
        if self.impact_factors:
            d['impactFactors'] = [e.to_dict() for e in self.impact_factors]
        if self.last_change:
            d['lastChange'] = self.last_change
        if self.name:
            d['name'] = self.name
        if self.parameters:
            d['parameters'] = [e.to_dict() for e in self.parameters]
        if self.ref_unit:
            d['refUnit'] = self.ref_unit
        if self.source:
            d['source'] = self.source.to_dict()
        if self.tags:
            d['tags'] = self.tags
        if self.version:
            d['version'] = self.version
        return d

    def to_json(self) -> str:
        return json.dumps(self.to_dict(), indent=2)

    def to_ref(self) -> 'Ref':
        ref = Ref(id=self.id, name=self.name)
        ref.category = self.category
        ref.model_type = 'ImpactCategory'
        return ref

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'ImpactCategory':
        impact_category = ImpactCategory()
        if v := d.get('@type'):
            impact_category.schema_type = v
        if v := d.get('@id'):
            impact_category.id = v
        if v := d.get('category'):
            impact_category.category = v
        if v := d.get('code'):
            impact_category.code = v
        if v := d.get('description'):
            impact_category.description = v
        if v := d.get('direction'):
            impact_category.direction = Direction.get(v)
        if v := d.get('impactFactors'):
            impact_category.impact_factors = [ImpactFactor.from_dict(e) for e in v]
        if v := d.get('lastChange'):
            impact_category.last_change = v
        if v := d.get('name'):
            impact_category.name = v
        if v := d.get('parameters'):
            impact_category.parameters = [Parameter.from_dict(e) for e in v]
        if v := d.get('refUnit'):
            impact_category.ref_unit = v
        if v := d.get('source'):
            impact_category.source = Ref.from_dict(v)
        if v := d.get('tags'):
            impact_category.tags = v
        if v := d.get('version'):
            impact_category.version = v
        return impact_category

    @staticmethod
    def from_json(data: Union[str, bytes]) -> 'ImpactCategory':
        return ImpactCategory.from_dict(json.loads(data))


@dataclass
class ParameterRedef:

    context: Optional[Ref] = None
    description: Optional[str] = None
    is_protected: Optional[bool] = None
    name: Optional[str] = None
    uncertainty: Optional[Uncertainty] = None
    value: Optional[float] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.context:
            d['context'] = self.context.to_dict()
        if self.description:
            d['description'] = self.description
        if self.is_protected:
            d['isProtected'] = self.is_protected
        if self.name:
            d['name'] = self.name
        if self.uncertainty:
            d['uncertainty'] = self.uncertainty.to_dict()
        if self.value:
            d['value'] = self.value
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'ParameterRedef':
        parameter_redef = ParameterRedef()
        if v := d.get('@type'):
            parameter_redef.schema_type = v
        if v := d.get('context'):
            parameter_redef.context = Ref.from_dict(v)
        if v := d.get('description'):
            parameter_redef.description = v
        if v := d.get('isProtected'):
            parameter_redef.is_protected = v
        if v := d.get('name'):
            parameter_redef.name = v
        if v := d.get('uncertainty'):
            parameter_redef.uncertainty = Uncertainty.from_dict(v)
        if v := d.get('value'):
            parameter_redef.value = v
        return parameter_redef


@dataclass
class ParameterRedefSet:

    description: Optional[str] = None
    is_baseline: Optional[bool] = None
    name: Optional[str] = None
    parameters: Optional[List[ParameterRedef]] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.description:
            d['description'] = self.description
        if self.is_baseline:
            d['isBaseline'] = self.is_baseline
        if self.name:
            d['name'] = self.name
        if self.parameters:
            d['parameters'] = [e.to_dict() for e in self.parameters]
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'ParameterRedefSet':
        parameter_redef_set = ParameterRedefSet()
        if v := d.get('@type'):
            parameter_redef_set.schema_type = v
        if v := d.get('description'):
            parameter_redef_set.description = v
        if v := d.get('isBaseline'):
            parameter_redef_set.is_baseline = v
        if v := d.get('name'):
            parameter_redef_set.name = v
        if v := d.get('parameters'):
            parameter_redef_set.parameters = [ParameterRedef.from_dict(e) for e in v]
        return parameter_redef_set


@dataclass
class Process:

    id: Optional[str] = None
    allocation_factors: Optional[List[AllocationFactor]] = None
    category: Optional[str] = None
    default_allocation_method: Optional[AllocationType] = None
    description: Optional[str] = None
    dq_entry: Optional[str] = None
    dq_system: Optional[Ref] = None
    exchange_dq_system: Optional[Ref] = None
    exchanges: Optional[List[Exchange]] = None
    is_infrastructure_process: Optional[bool] = None
    last_change: Optional[str] = None
    last_internal_id: Optional[int] = None
    location: Optional[Ref] = None
    name: Optional[str] = None
    parameters: Optional[List[Parameter]] = None
    process_documentation: Optional[ProcessDocumentation] = None
    process_type: Optional[ProcessType] = None
    social_aspects: Optional[List[SocialAspect]] = None
    social_dq_system: Optional[Ref] = None
    tags: Optional[List[str]] = None
    version: Optional[str] = None

    def __post_init__(self):
        if self.id is None:
            self.id = str(uuid.uuid4())
        if self.version is None:
            self.version = '01.00.000'
        if self.last_change is None:
            self.last_change = datetime.datetime.utcnow().isoformat() + 'Z'

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        d['@type'] = 'Process'
        if self.id:
            d['@id'] = self.id
        if self.allocation_factors:
            d['allocationFactors'] = [e.to_dict() for e in self.allocation_factors]
        if self.category:
            d['category'] = self.category
        if self.default_allocation_method:
            d['defaultAllocationMethod'] = self.default_allocation_method.value
        if self.description:
            d['description'] = self.description
        if self.dq_entry:
            d['dqEntry'] = self.dq_entry
        if self.dq_system:
            d['dqSystem'] = self.dq_system.to_dict()
        if self.exchange_dq_system:
            d['exchangeDqSystem'] = self.exchange_dq_system.to_dict()
        if self.exchanges:
            d['exchanges'] = [e.to_dict() for e in self.exchanges]
        if self.is_infrastructure_process:
            d['isInfrastructureProcess'] = self.is_infrastructure_process
        if self.last_change:
            d['lastChange'] = self.last_change
        if self.last_internal_id:
            d['lastInternalId'] = self.last_internal_id
        if self.location:
            d['location'] = self.location.to_dict()
        if self.name:
            d['name'] = self.name
        if self.parameters:
            d['parameters'] = [e.to_dict() for e in self.parameters]
        if self.process_documentation:
            d['processDocumentation'] = self.process_documentation.to_dict()
        if self.process_type:
            d['processType'] = self.process_type.value
        if self.social_aspects:
            d['socialAspects'] = [e.to_dict() for e in self.social_aspects]
        if self.social_dq_system:
            d['socialDqSystem'] = self.social_dq_system.to_dict()
        if self.tags:
            d['tags'] = self.tags
        if self.version:
            d['version'] = self.version
        return d

    def to_json(self) -> str:
        return json.dumps(self.to_dict(), indent=2)

    def to_ref(self) -> 'Ref':
        ref = Ref(id=self.id, name=self.name)
        ref.category = self.category
        ref.model_type = 'Process'
        return ref

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'Process':
        process = Process()
        if v := d.get('@type'):
            process.schema_type = v
        if v := d.get('@id'):
            process.id = v
        if v := d.get('allocationFactors'):
            process.allocation_factors = [AllocationFactor.from_dict(e) for e in v]
        if v := d.get('category'):
            process.category = v
        if v := d.get('defaultAllocationMethod'):
            process.default_allocation_method = AllocationType.get(v)
        if v := d.get('description'):
            process.description = v
        if v := d.get('dqEntry'):
            process.dq_entry = v
        if v := d.get('dqSystem'):
            process.dq_system = Ref.from_dict(v)
        if v := d.get('exchangeDqSystem'):
            process.exchange_dq_system = Ref.from_dict(v)
        if v := d.get('exchanges'):
            process.exchanges = [Exchange.from_dict(e) for e in v]
        if v := d.get('isInfrastructureProcess'):
            process.is_infrastructure_process = v
        if v := d.get('lastChange'):
            process.last_change = v
        if v := d.get('lastInternalId'):
            process.last_internal_id = v
        if v := d.get('location'):
            process.location = Ref.from_dict(v)
        if v := d.get('name'):
            process.name = v
        if v := d.get('parameters'):
            process.parameters = [Parameter.from_dict(e) for e in v]
        if v := d.get('processDocumentation'):
            process.process_documentation = ProcessDocumentation.from_dict(v)
        if v := d.get('processType'):
            process.process_type = ProcessType.get(v)
        if v := d.get('socialAspects'):
            process.social_aspects = [SocialAspect.from_dict(e) for e in v]
        if v := d.get('socialDqSystem'):
            process.social_dq_system = Ref.from_dict(v)
        if v := d.get('tags'):
            process.tags = v
        if v := d.get('version'):
            process.version = v
        return process

    @staticmethod
    def from_json(data: Union[str, bytes]) -> 'Process':
        return Process.from_dict(json.loads(data))


@dataclass
class ProductSystem:

    id: Optional[str] = None
    category: Optional[str] = None
    description: Optional[str] = None
    last_change: Optional[str] = None
    name: Optional[str] = None
    parameter_sets: Optional[List[ParameterRedefSet]] = None
    process_links: Optional[List[ProcessLink]] = None
    processes: Optional[List[Ref]] = None
    ref_exchange: Optional[ExchangeRef] = None
    ref_process: Optional[Ref] = None
    tags: Optional[List[str]] = None
    target_amount: Optional[float] = None
    target_flow_property: Optional[Ref] = None
    target_unit: Optional[Ref] = None
    version: Optional[str] = None

    def __post_init__(self):
        if self.id is None:
            self.id = str(uuid.uuid4())
        if self.version is None:
            self.version = '01.00.000'
        if self.last_change is None:
            self.last_change = datetime.datetime.utcnow().isoformat() + 'Z'

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        d['@type'] = 'ProductSystem'
        if self.id:
            d['@id'] = self.id
        if self.category:
            d['category'] = self.category
        if self.description:
            d['description'] = self.description
        if self.last_change:
            d['lastChange'] = self.last_change
        if self.name:
            d['name'] = self.name
        if self.parameter_sets:
            d['parameterSets'] = [e.to_dict() for e in self.parameter_sets]
        if self.process_links:
            d['processLinks'] = [e.to_dict() for e in self.process_links]
        if self.processes:
            d['processes'] = [e.to_dict() for e in self.processes]
        if self.ref_exchange:
            d['refExchange'] = self.ref_exchange.to_dict()
        if self.ref_process:
            d['refProcess'] = self.ref_process.to_dict()
        if self.tags:
            d['tags'] = self.tags
        if self.target_amount:
            d['targetAmount'] = self.target_amount
        if self.target_flow_property:
            d['targetFlowProperty'] = self.target_flow_property.to_dict()
        if self.target_unit:
            d['targetUnit'] = self.target_unit.to_dict()
        if self.version:
            d['version'] = self.version
        return d

    def to_json(self) -> str:
        return json.dumps(self.to_dict(), indent=2)

    def to_ref(self) -> 'Ref':
        ref = Ref(id=self.id, name=self.name)
        ref.category = self.category
        ref.model_type = 'ProductSystem'
        return ref

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'ProductSystem':
        product_system = ProductSystem()
        if v := d.get('@type'):
            product_system.schema_type = v
        if v := d.get('@id'):
            product_system.id = v
        if v := d.get('category'):
            product_system.category = v
        if v := d.get('description'):
            product_system.description = v
        if v := d.get('lastChange'):
            product_system.last_change = v
        if v := d.get('name'):
            product_system.name = v
        if v := d.get('parameterSets'):
            product_system.parameter_sets = [ParameterRedefSet.from_dict(e) for e in v]
        if v := d.get('processLinks'):
            product_system.process_links = [ProcessLink.from_dict(e) for e in v]
        if v := d.get('processes'):
            product_system.processes = [Ref.from_dict(e) for e in v]
        if v := d.get('refExchange'):
            product_system.ref_exchange = ExchangeRef.from_dict(v)
        if v := d.get('refProcess'):
            product_system.ref_process = Ref.from_dict(v)
        if v := d.get('tags'):
            product_system.tags = v
        if v := d.get('targetAmount'):
            product_system.target_amount = v
        if v := d.get('targetFlowProperty'):
            product_system.target_flow_property = Ref.from_dict(v)
        if v := d.get('targetUnit'):
            product_system.target_unit = Ref.from_dict(v)
        if v := d.get('version'):
            product_system.version = v
        return product_system

    @staticmethod
    def from_json(data: Union[str, bytes]) -> 'ProductSystem':
        return ProductSystem.from_dict(json.loads(data))


@dataclass
class ProjectVariant:

    allocation_method: Optional[AllocationType] = None
    amount: Optional[float] = None
    description: Optional[str] = None
    is_disabled: Optional[bool] = None
    name: Optional[str] = None
    parameter_redefs: Optional[List[ParameterRedef]] = None
    product_system: Optional[Ref] = None
    unit: Optional[Ref] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.allocation_method:
            d['allocationMethod'] = self.allocation_method.value
        if self.amount:
            d['amount'] = self.amount
        if self.description:
            d['description'] = self.description
        if self.is_disabled:
            d['isDisabled'] = self.is_disabled
        if self.name:
            d['name'] = self.name
        if self.parameter_redefs:
            d['parameterRedefs'] = [e.to_dict() for e in self.parameter_redefs]
        if self.product_system:
            d['productSystem'] = self.product_system.to_dict()
        if self.unit:
            d['unit'] = self.unit.to_dict()
        return d

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'ProjectVariant':
        project_variant = ProjectVariant()
        if v := d.get('@type'):
            project_variant.schema_type = v
        if v := d.get('allocationMethod'):
            project_variant.allocation_method = AllocationType.get(v)
        if v := d.get('amount'):
            project_variant.amount = v
        if v := d.get('description'):
            project_variant.description = v
        if v := d.get('isDisabled'):
            project_variant.is_disabled = v
        if v := d.get('name'):
            project_variant.name = v
        if v := d.get('parameterRedefs'):
            project_variant.parameter_redefs = [ParameterRedef.from_dict(e) for e in v]
        if v := d.get('productSystem'):
            project_variant.product_system = Ref.from_dict(v)
        if v := d.get('unit'):
            project_variant.unit = Ref.from_dict(v)
        return project_variant


@dataclass
class Project:

    id: Optional[str] = None
    category: Optional[str] = None
    description: Optional[str] = None
    impact_method: Optional[Ref] = None
    is_with_costs: Optional[bool] = None
    is_with_regionalization: Optional[bool] = None
    last_change: Optional[str] = None
    name: Optional[str] = None
    nw_set: Optional[NwSet] = None
    tags: Optional[List[str]] = None
    variants: Optional[List[ProjectVariant]] = None
    version: Optional[str] = None

    def __post_init__(self):
        if self.id is None:
            self.id = str(uuid.uuid4())
        if self.version is None:
            self.version = '01.00.000'
        if self.last_change is None:
            self.last_change = datetime.datetime.utcnow().isoformat() + 'Z'

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        d['@type'] = 'Project'
        if self.id:
            d['@id'] = self.id
        if self.category:
            d['category'] = self.category
        if self.description:
            d['description'] = self.description
        if self.impact_method:
            d['impactMethod'] = self.impact_method.to_dict()
        if self.is_with_costs:
            d['isWithCosts'] = self.is_with_costs
        if self.is_with_regionalization:
            d['isWithRegionalization'] = self.is_with_regionalization
        if self.last_change:
            d['lastChange'] = self.last_change
        if self.name:
            d['name'] = self.name
        if self.nw_set:
            d['nwSet'] = self.nw_set.to_dict()
        if self.tags:
            d['tags'] = self.tags
        if self.variants:
            d['variants'] = [e.to_dict() for e in self.variants]
        if self.version:
            d['version'] = self.version
        return d

    def to_json(self) -> str:
        return json.dumps(self.to_dict(), indent=2)

    def to_ref(self) -> 'Ref':
        ref = Ref(id=self.id, name=self.name)
        ref.category = self.category
        ref.model_type = 'Project'
        return ref

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'Project':
        project = Project()
        if v := d.get('@type'):
            project.schema_type = v
        if v := d.get('@id'):
            project.id = v
        if v := d.get('category'):
            project.category = v
        if v := d.get('description'):
            project.description = v
        if v := d.get('impactMethod'):
            project.impact_method = Ref.from_dict(v)
        if v := d.get('isWithCosts'):
            project.is_with_costs = v
        if v := d.get('isWithRegionalization'):
            project.is_with_regionalization = v
        if v := d.get('lastChange'):
            project.last_change = v
        if v := d.get('name'):
            project.name = v
        if v := d.get('nwSet'):
            project.nw_set = NwSet.from_dict(v)
        if v := d.get('tags'):
            project.tags = v
        if v := d.get('variants'):
            project.variants = [ProjectVariant.from_dict(e) for e in v]
        if v := d.get('version'):
            project.version = v
        return project

    @staticmethod
    def from_json(data: Union[str, bytes]) -> 'Project':
        return Project.from_dict(json.loads(data))


@dataclass
class Unit:

    id: Optional[str] = None
    conversion_factor: Optional[float] = None
    description: Optional[str] = None
    is_ref_unit: Optional[bool] = None
    name: Optional[str] = None
    synonyms: Optional[List[str]] = None

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        if self.id:
            d['@id'] = self.id
        if self.conversion_factor:
            d['conversionFactor'] = self.conversion_factor
        if self.description:
            d['description'] = self.description
        if self.is_ref_unit:
            d['isRefUnit'] = self.is_ref_unit
        if self.name:
            d['name'] = self.name
        if self.synonyms:
            d['synonyms'] = self.synonyms
        return d

    def to_ref(self) -> 'Ref':
        ref = Ref(id=self.id, name=self.name)
        ref.model_type = 'Unit'
        return ref

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'Unit':
        unit = Unit()
        if v := d.get('@type'):
            unit.schema_type = v
        if v := d.get('@id'):
            unit.id = v
        if v := d.get('conversionFactor'):
            unit.conversion_factor = v
        if v := d.get('description'):
            unit.description = v
        if v := d.get('isRefUnit'):
            unit.is_ref_unit = v
        if v := d.get('name'):
            unit.name = v
        if v := d.get('synonyms'):
            unit.synonyms = v
        return unit


@dataclass
class UnitGroup:

    id: Optional[str] = None
    category: Optional[str] = None
    default_flow_property: Optional[Ref] = None
    description: Optional[str] = None
    last_change: Optional[str] = None
    name: Optional[str] = None
    tags: Optional[List[str]] = None
    units: Optional[List[Unit]] = None
    version: Optional[str] = None

    def __post_init__(self):
        if self.id is None:
            self.id = str(uuid.uuid4())
        if self.version is None:
            self.version = '01.00.000'
        if self.last_change is None:
            self.last_change = datetime.datetime.utcnow().isoformat() + 'Z'

    def to_dict(self) -> Dict[str, Any]:
        d: Dict[str, Any] = {}
        d['@type'] = 'UnitGroup'
        if self.id:
            d['@id'] = self.id
        if self.category:
            d['category'] = self.category
        if self.default_flow_property:
            d['defaultFlowProperty'] = self.default_flow_property.to_dict()
        if self.description:
            d['description'] = self.description
        if self.last_change:
            d['lastChange'] = self.last_change
        if self.name:
            d['name'] = self.name
        if self.tags:
            d['tags'] = self.tags
        if self.units:
            d['units'] = [e.to_dict() for e in self.units]
        if self.version:
            d['version'] = self.version
        return d

    def to_json(self) -> str:
        return json.dumps(self.to_dict(), indent=2)

    def to_ref(self) -> 'Ref':
        ref = Ref(id=self.id, name=self.name)
        ref.category = self.category
        ref.model_type = 'UnitGroup'
        return ref

    @staticmethod
    def from_dict(d: Dict[str, Any]) -> 'UnitGroup':
        unit_group = UnitGroup()
        if v := d.get('@type'):
            unit_group.schema_type = v
        if v := d.get('@id'):
            unit_group.id = v
        if v := d.get('category'):
            unit_group.category = v
        if v := d.get('defaultFlowProperty'):
            unit_group.default_flow_property = Ref.from_dict(v)
        if v := d.get('description'):
            unit_group.description = v
        if v := d.get('lastChange'):
            unit_group.last_change = v
        if v := d.get('name'):
            unit_group.name = v
        if v := d.get('tags'):
            unit_group.tags = v
        if v := d.get('units'):
            unit_group.units = [Unit.from_dict(e) for e in v]
        if v := d.get('version'):
            unit_group.version = v
        return unit_group

    @staticmethod
    def from_json(data: Union[str, bytes]) -> 'UnitGroup':
        return UnitGroup.from_dict(json.loads(data))


RootEntity = Union[
    Actor,
    Currency,
    DQSystem,
    Epd,
    Flow,
    FlowMap,
    FlowProperty,
    ImpactCategory,
    ImpactMethod,
    Location,
    Parameter,
    Process,
    ProductSystem,
    Project,
    Result,
    SocialIndicator,
    Source,
    UnitGroup,
]
