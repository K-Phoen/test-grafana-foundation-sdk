---
title: <span class="badge object-type-struct"></span> TableSparklineCellOptions
---
# <span class="badge object-type-struct"></span> TableSparklineCellOptions

Sparkline cell options

## Definition

```go
type TableSparklineCellOptions struct {
    Type string `json:"type"`
    DrawStyle *common.GraphDrawStyle `json:"drawStyle,omitempty"`
    GradientMode *common.GraphGradientMode `json:"gradientMode,omitempty"`
    ThresholdsStyle *common.GraphThresholdsStyleConfig `json:"thresholdsStyle,omitempty"`
    LineColor *string `json:"lineColor,omitempty"`
    LineWidth *float64 `json:"lineWidth,omitempty"`
    LineInterpolation *common.LineInterpolation `json:"lineInterpolation,omitempty"`
    LineStyle *common.LineStyle `json:"lineStyle,omitempty"`
    FillColor *string `json:"fillColor,omitempty"`
    FillOpacity *float64 `json:"fillOpacity,omitempty"`
    ShowPoints *common.VisibilityMode `json:"showPoints,omitempty"`
    PointSize *float64 `json:"pointSize,omitempty"`
    PointColor *string `json:"pointColor,omitempty"`
    AxisPlacement *common.AxisPlacement `json:"axisPlacement,omitempty"`
    AxisColorMode *common.AxisColorMode `json:"axisColorMode,omitempty"`
    AxisLabel *string `json:"axisLabel,omitempty"`
    AxisWidth *float64 `json:"axisWidth,omitempty"`
    AxisSoftMin *float64 `json:"axisSoftMin,omitempty"`
    AxisSoftMax *float64 `json:"axisSoftMax,omitempty"`
    AxisGridShow *bool `json:"axisGridShow,omitempty"`
    ScaleDistribution *common.ScaleDistributionConfig `json:"scaleDistribution,omitempty"`
    BarAlignment *common.BarAlignment `json:"barAlignment,omitempty"`
    BarWidthFactor *float64 `json:"barWidthFactor,omitempty"`
    Stacking *common.StackingConfig `json:"stacking,omitempty"`
    HideFrom *common.HideSeriesConfig `json:"hideFrom,omitempty"`
    Transform *common.GraphTransform `json:"transform,omitempty"`
    // Indicate if null values should be treated as gaps or connected.
    // When the value is a number, it represents the maximum delta in the
    // X axis that should be considered connected.  For timeseries, this is milliseconds
    SpanNulls *common.BoolOrFloat64 `json:"spanNulls,omitempty"`
    FillBelowTo *string `json:"fillBelowTo,omitempty"`
    PointSymbol *string `json:"pointSymbol,omitempty"`
    AxisCenteredZero *bool `json:"axisCenteredZero,omitempty"`
    BarMaxWidth *float64 `json:"barMaxWidth,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TableSparklineCellOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (tableSparklineCellOptions *TableSparklineCellOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TableSparklineCellOptions` objects.

```go
func (tableSparklineCellOptions *TableSparklineCellOptions) Equals(other TableSparklineCellOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TableSparklineCellOptions` fields for violations and returns them.

```go
func (tableSparklineCellOptions *TableSparklineCellOptions) Validate() error
```

## See also

 * <span class="badge builder"></span> [TableSparklineCellOptionsBuilder](./builder-TableSparklineCellOptionsBuilder.md)