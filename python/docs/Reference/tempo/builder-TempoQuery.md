---
title: <span class="badge builder"></span> TempoQuery
---
# <span class="badge builder"></span> TempoQuery

## Constructor

```python
TempoQuery()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> tempo.TempoQuery
```

### <span class="badge object-method"></span> datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```python
def datasource(datasource: dashboard.DataSourceRef) -> typing.Self
```

### <span class="badge object-method"></span> filters

```python
def filters(filters: list[cogbuilder.Builder[tempo.TraceqlFilter]]) -> typing.Self
```

### <span class="badge object-method"></span> hide

true if query is disabled (ie should not be returned to the dashboard)

Note this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```python
def hide(hide: bool) -> typing.Self
```

### <span class="badge object-method"></span> limit

Defines the maximum number of traces that are returned from Tempo

```python
def limit(limit: int) -> typing.Self
```

### <span class="badge object-method"></span> max_duration

@deprecated Define the maximum duration to select traces. Use duration format, for example: 1.2s, 100ms

```python
def max_duration(max_duration: str) -> typing.Self
```

### <span class="badge object-method"></span> min_duration

@deprecated Define the minimum duration to select traces. Use duration format, for example: 1.2s, 100ms

```python
def min_duration(min_duration: str) -> typing.Self
```

### <span class="badge object-method"></span> query

TraceQL query or trace ID

```python
def query(query: str) -> typing.Self
```

### <span class="badge object-method"></span> query_type

Specify the query flavor

TODO make this required and give it a default

```python
def query_type(query_type: str) -> typing.Self
```

### <span class="badge object-method"></span> ref_id

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```python
def ref_id(ref_id: str) -> typing.Self
```

### <span class="badge object-method"></span> search

@deprecated Logfmt query to filter traces by their tags. Example: http.status_code=200 error=true

```python
def search(search: str) -> typing.Self
```

### <span class="badge object-method"></span> service_map_include_namespace

Use service.namespace in addition to service.name to uniquely identify a service.

```python
def service_map_include_namespace(service_map_include_namespace: bool) -> typing.Self
```

### <span class="badge object-method"></span> service_map_query

Filters to be included in a PromQL query to select data for the service graph. Example: {client="app",service="app"}

```python
def service_map_query(service_map_query: str) -> typing.Self
```

### <span class="badge object-method"></span> service_name

@deprecated Query traces by service name

```python
def service_name(service_name: str) -> typing.Self
```

### <span class="badge object-method"></span> span_name

@deprecated Query traces by span name

```python
def span_name(span_name: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [TempoQuery](./object-TempoQuery.md)