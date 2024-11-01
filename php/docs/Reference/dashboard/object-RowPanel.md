---
title: <span class="badge object-type-class"></span> RowPanel
---
# <span class="badge object-type-class"></span> RowPanel

Row panel

## Definition

```php
class RowPanel implements \JsonSerializable
{
    /**
     * The panel type
     */
    public string $type;

    /**
     * Whether this row should be collapsed or not.
     */
    public bool $collapsed;

    /**
     * Row title
     */
    public ?string $title;

    /**
     * Name of default datasource for the row
     */
    public ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource;

    /**
     * Row grid position
     */
    public ?\Grafana\Foundation\Dashboard\GridPos $gridPos;

    /**
     * Unique identifier of the panel. Generated by Grafana when creating a new panel. It must be unique within a dashboard, but not globally.
     */
    public int $id;

    /**
     * List of panels in the row
     * @var array<\Grafana\Foundation\Dashboard\Panel>
     */
    public array $panels;

    /**
     * Name of template variable to repeat for.
     */
    public ?string $repeat;

}
```
## Methods

### <span class="badge object-method"></span> fromArray

Builds this object from an array.

This function is meant to be used with the return value of `json_decode($json, true)`.

```php
static fromArray(array $inputData)
```

### <span class="badge object-method"></span> jsonSerialize

Returns the data representing this object, preparing it for JSON serialization with `json_encode()`.

```php
jsonSerialize()
```

## See also

 * <span class="badge builder"></span> [RowBuilder](./builder-RowBuilder.md)