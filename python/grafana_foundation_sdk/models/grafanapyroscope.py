# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import enum
from ..cog import variants as cogvariants
import typing
from ..cog import runtime as cogruntime


class PyroscopeQueryType(enum.StrEnum):
    METRICS = "metrics"
    PROFILE = "profile"
    BOTH = "both"


class Dataquery(cogvariants.Dataquery):
    # Specifies the query label selectors.
    label_selector: typing.Optional[str]
    # Specifies the query span selectors.
    span_selector: typing.Optional[list[str]]
    # Specifies the type of profile to query.
    profile_type_id: typing.Optional[str]
    # Allows to group the results.
    group_by: typing.Optional[list[str]]
    # Sets the maximum number of nodes in the flamegraph.
    max_nodes: typing.Optional[int]
    # A unique identifier for the query within the list of targets.
    # In server side expressions, the refId is used as a variable name to identify results.
    # By default, the UI will assign A->Z; however setting meaningful names may be useful.
    ref_id: typing.Optional[str]
    # If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
    hide: typing.Optional[bool]
    # Specify the query flavor
    # TODO make this required and give it a default
    query_type: typing.Optional[str]
    # For mixed data sources the selected datasource is on the query level.
    # For non mixed scenarios this is undefined.
    # TODO find a better way to do this ^ that's friendly to schema
    # TODO this shouldn't be unknown but DataSourceRef | null
    datasource: typing.Optional[object]

    def __init__(self, label_selector: typing.Optional[str] = "{}", span_selector: typing.Optional[list[str]] = None, profile_type_id: typing.Optional[str] = None, group_by: typing.Optional[list[str]] = None, max_nodes: typing.Optional[int] = None, ref_id: typing.Optional[str] = None, hide: typing.Optional[bool] = None, query_type: typing.Optional[str] = None, datasource: typing.Optional[object] = None):
        self.label_selector = label_selector
        self.span_selector = span_selector
        self.profile_type_id = profile_type_id
        self.group_by = group_by
        self.max_nodes = max_nodes
        self.ref_id = ref_id
        self.hide = hide
        self.query_type = query_type
        self.datasource = datasource

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.label_selector is not None:
            payload["labelSelector"] = self.label_selector
        if self.span_selector is not None:
            payload["spanSelector"] = self.span_selector
        if self.profile_type_id is not None:
            payload["profileTypeId"] = self.profile_type_id
        if self.group_by is not None:
            payload["groupBy"] = self.group_by
        if self.max_nodes is not None:
            payload["maxNodes"] = self.max_nodes
        if self.ref_id is not None:
            payload["refId"] = self.ref_id
        if self.hide is not None:
            payload["hide"] = self.hide
        if self.query_type is not None:
            payload["queryType"] = self.query_type
        if self.datasource is not None:
            payload["datasource"] = self.datasource
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "labelSelector" in data:
            args["label_selector"] = data["labelSelector"]
        if "spanSelector" in data:
            args["span_selector"] = data["spanSelector"]
        if "profileTypeId" in data:
            args["profile_type_id"] = data["profileTypeId"]
        if "groupBy" in data:
            args["group_by"] = data["groupBy"]
        if "maxNodes" in data:
            args["max_nodes"] = data["maxNodes"]
        if "refId" in data:
            args["ref_id"] = data["refId"]
        if "hide" in data:
            args["hide"] = data["hide"]
        if "queryType" in data:
            args["query_type"] = data["queryType"]
        if "datasource" in data:
            args["datasource"] = data["datasource"]        

        return cls(**args)


def variant_config() -> cogruntime.DataqueryConfig:
    return cogruntime.DataqueryConfig(
        identifier="grafanapyroscope",
        from_json_hook=Dataquery.from_json,
    )


