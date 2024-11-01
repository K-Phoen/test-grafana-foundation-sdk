---
title: <span class="badge object-type-class"></span> Dashboard
---
# <span class="badge object-type-class"></span> Dashboard

## Definition

```python
class Dashboard:
    # Unique numeric identifier for the dashboard.
    # `id` is internal to a specific Grafana instance. `uid` should be used to identify a dashboard across Grafana instances.
    id_val: typing.Optional[int]
    # Unique dashboard identifier that can be generated by anyone. string (8-40)
    uid: typing.Optional[str]
    # Title of dashboard.
    title: typing.Optional[str]
    # Description of dashboard.
    description: typing.Optional[str]
    # This property should only be used in dashboards defined by plugins.  It is a quick check
    # to see if the version has changed since the last time.
    revision: typing.Optional[int]
    # ID of a dashboard imported from the https://grafana.com/grafana/dashboards/ portal
    gnet_id: typing.Optional[str]
    # Tags associated with dashboard.
    tags: typing.Optional[list[str]]
    # Theme of dashboard.
    # Default value: dark.
    style: typing.Literal["light", "dark"]
    # Timezone of dashboard. Accepted values are IANA TZDB zone ID or "browser" or "utc".
    timezone: typing.Optional[str]
    # Whether a dashboard is editable or not.
    editable: bool
    # Configuration of dashboard cursor sync behavior.
    # Accepted values are 0 (sync turned off), 1 (shared crosshair), 2 (shared crosshair and tooltip).
    graph_tooltip: dashboard.DashboardCursorSync
    # Time range for dashboard.
    # Accepted values are relative time strings like {from: 'now-6h', to: 'now'} or absolute time strings like {from: '2020-07-10T08:00:00.000Z', to: '2020-07-10T14:00:00.000Z'}.
    time: typing.Optional[dashboard.DashboardDashboardTime]
    # Configuration of the time picker shown at the top of a dashboard.
    timepicker: typing.Optional[dashboard.TimePicker]
    # The month that the fiscal year starts on.  0 = January, 11 = December
    fiscal_year_start_month: typing.Optional[int]
    # When set to true, the dashboard will redraw panels at an interval matching the pixel width.
    # This will keep data "moving left" regardless of the query refresh rate. This setting helps
    # avoid dashboards presenting stale live data
    live_now: typing.Optional[bool]
    # Day when the week starts. Expressed by the name of the day in lowercase, e.g. "monday".
    week_start: typing.Optional[str]
    # Refresh rate of dashboard. Represented via interval string, e.g. "5s", "1m", "1h", "1d".
    refresh: typing.Optional[typing.Union[str, typing.Literal[False]]]
    # Version of the JSON schema, incremented each time a Grafana update brings
    # changes to said schema.
    schema_version: int
    # Version of the dashboard, incremented each time the dashboard is updated.
    version: typing.Optional[int]
    # List of dashboard panels
    panels: typing.Optional[list[typing.Union[dashboard.Panel, dashboard.RowPanel]]]
    # Configured template variables
    templating: dashboard.DashboardDashboardTemplating
    # Contains the list of annotations that are associated with the dashboard.
    # Annotations are used to overlay event markers and overlay event tags on graphs.
    # Grafana comes with a native annotation store and the ability to add annotation events directly from the graph panel or via the HTTP API.
    # See https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/annotate-visualizations/
    annotations: dashboard.AnnotationContainer
    # Links with references to other dashboards or external websites.
    links: typing.Optional[list[dashboard.DashboardLink]]
    # Snapshot options. They are present only if the dashboard is a snapshot.
    snapshot: typing.Optional[dashboard.Snapshot]
```
## Methods

### <span class="badge object-method"></span> to_json

Converts this object into a representation that can easily be encoded to JSON.

```python
def to_json() -> dict[str, object]
```

### <span class="badge object-method"></span> from_json

Builds this object from a JSON-decoded dict.

```python
@classmethod
def from_json(data: dict[str, typing.Any]) -> typing.Self
```

## Examples

### Marshalling to JSON

```python
from grafana_foundation_sdk.cog.encoder import JSONEncoder
from grafana_foundation_sdk.models.dashboard import Dashboard


if __name__ == '__main__':
    dashboard = Dashboard()

    encoder = JSONEncoder(sort_keys=True, indent=2)
    print(encoder.encode(dashboard))
```

### Unmarshalling from JSON

```python
import json

from grafana_foundation_sdk.cog.plugins import register_default_plugins
from grafana_foundation_sdk.models.dashboard import Dashboard as DashboardModel


if __name__ == '__main__':
    # Required to correctly unmarshal panels and dataqueries
    register_default_plugins()

    with open("dashboard.json", "r") as f:
        decoded_dashboard = DashboardModel.from_json(json.load(f))
        print(decoded_dashboard)
```
## See also

 * <span class="badge builder"></span> [Dashboard](./builder-Dashboard.md)