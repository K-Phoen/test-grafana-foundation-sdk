---
title: <span class="badge builder"></span> DataQueryBuilder
---
# <span class="badge builder"></span> DataQueryBuilder

## Constructor

```typescript
new DataQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```typescript
datasource(datasource: any)
```

### <span class="badge object-method"></span> hide

true if query is disabled (ie should not be returned to the dashboard)

Note this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```typescript
hide(hide: boolean)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```typescript
queryType(queryType: string)
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```typescript
refId(refId: string)
```

## See also

 * <span class="badge object-type-interface"></span> [DataQuery](./object-DataQuery.md)