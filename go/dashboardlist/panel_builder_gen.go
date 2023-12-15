// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardlist

import (
	"errors"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	cogvariants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

var _ cog.Builder[dashboard.Panel] = (*PanelBuilder)(nil)

// Dashboard panels are the basic visualization building blocks.
type PanelBuilder struct {
	internal *dashboard.Panel
	errors   map[string]cog.BuildErrors
}

func NewPanelBuilder() *PanelBuilder {
	resource := &dashboard.Panel{}
	builder := &PanelBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "dashboardlist"

	return builder
}

func (builder *PanelBuilder) Build() (dashboard.Panel, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("Panel", err)...)
	}

	if len(errs) != 0 {
		return dashboard.Panel{}, errs
	}

	return *builder.internal, nil
}

// Depends on the panel plugin. See the plugin documentation for details.
func (builder *PanelBuilder) WithTarget(targets cog.Builder[cogvariants.Dataquery]) *PanelBuilder {
	targetsResource, err := targets.Build()
	if err != nil {
		builder.errors["targets"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Targets = append(builder.internal.Targets, targetsResource)

	return builder
}

// Panel title.
func (builder *PanelBuilder) Title(title string) *PanelBuilder {
	builder.internal.Title = &title

	return builder
}

// Panel description.
func (builder *PanelBuilder) Description(description string) *PanelBuilder {
	builder.internal.Description = &description

	return builder
}

// Whether to display the panel without a background.
func (builder *PanelBuilder) Transparent(transparent bool) *PanelBuilder {
	builder.internal.Transparent = &transparent

	return builder
}

// The datasource used in all targets.
func (builder *PanelBuilder) Datasource(datasource dashboard.DataSourceRef) *PanelBuilder {
	builder.internal.Datasource = &datasource

	return builder
}

// Panel height. The height is the number of rows from the top edge of the panel.
func (builder *PanelBuilder) Height(h uint32) *PanelBuilder {
	if !(h > 0) {
		builder.errors["h"] = cog.MakeBuildErrors("h", errors.New("h must be > 0"))
		return builder
	}
	if builder.internal.GridPos == nil {
		builder.internal.GridPos = &dashboard.GridPos{}
	}
	builder.internal.GridPos.H = h

	return builder
}

// Panel width. The width is the number of columns from the left edge of the panel.
func (builder *PanelBuilder) Span(w uint32) *PanelBuilder {
	if !(w > 0) {
		builder.errors["w"] = cog.MakeBuildErrors("w", errors.New("w must be > 0"))
		return builder
	}
	if !(w <= 24) {
		builder.errors["w"] = cog.MakeBuildErrors("w", errors.New("w must be <= 24"))
		return builder
	}
	if builder.internal.GridPos == nil {
		builder.internal.GridPos = &dashboard.GridPos{}
	}
	builder.internal.GridPos.W = w

	return builder
}

// Panel links.
func (builder *PanelBuilder) Links(links []cog.Builder[dashboard.DashboardLink]) *PanelBuilder {
	linksResources := make([]dashboard.DashboardLink, 0, len(links))
	for _, r := range links {
		linksResource, err := r.Build()
		if err != nil {
			builder.errors["links"] = err.(cog.BuildErrors)
			return builder
		}
		linksResources = append(linksResources, linksResource)
	}
	builder.internal.Links = linksResources

	return builder
}

// Name of template variable to repeat for.
func (builder *PanelBuilder) Repeat(repeat string) *PanelBuilder {
	builder.internal.Repeat = &repeat

	return builder
}

// Direction to repeat in if 'repeat' is set.
// `h` for horizontal, `v` for vertical.
func (builder *PanelBuilder) RepeatDirection(repeatDirection dashboard.PanelRepeatDirection) *PanelBuilder {
	builder.internal.RepeatDirection = &repeatDirection

	return builder
}

// Option for repeated panels that controls max items per row
// Only relevant for horizontally repeated panels
func (builder *PanelBuilder) MaxPerRow(maxPerRow float64) *PanelBuilder {
	builder.internal.MaxPerRow = &maxPerRow

	return builder
}

// The maximum number of data points that the panel queries are retrieving.
func (builder *PanelBuilder) MaxDataPoints(maxDataPoints float64) *PanelBuilder {
	builder.internal.MaxDataPoints = &maxDataPoints

	return builder
}

// List of transformations that are applied to the panel data before rendering.
// When there are multiple transformations, Grafana applies them in the order they are listed.
// Each transformation creates a result set that then passes on to the next transformation in the processing pipeline.
func (builder *PanelBuilder) WithTransformation(transformations dashboard.DataTransformerConfig) *PanelBuilder {
	builder.internal.Transformations = append(builder.internal.Transformations, transformations)

	return builder
}

// The min time interval setting defines a lower limit for the $__interval and $__interval_ms variables.
// This value must be formatted as a number followed by a valid time
// identifier like: "40s", "3d", etc.
// See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
func (builder *PanelBuilder) Interval(interval string) *PanelBuilder {
	builder.internal.Interval = &interval

	return builder
}

// Overrides the relative time range for individual panels,
// which causes them to be different than what is selected in
// the dashboard time picker in the top-right corner of the dashboard. You can use this to show metrics from different
// time periods or days on the same dashboard.
// The value is formatted as time operation like: `now-5m` (Last 5 minutes), `now/d` (the day so far),
// `now-5d/d`(Last 5 days), `now/w` (This week so far), `now-2y/y` (Last 2 years).
// Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
// See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
func (builder *PanelBuilder) TimeFrom(timeFrom string) *PanelBuilder {
	builder.internal.TimeFrom = &timeFrom

	return builder
}

// Overrides the time range for individual panels by shifting its start and end relative to the time picker.
// For example, you can shift the time range for the panel to be two hours earlier than the dashboard time picker setting `2h`.
// Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
// See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
func (builder *PanelBuilder) TimeShift(timeShift string) *PanelBuilder {
	builder.internal.TimeShift = &timeShift

	return builder
}

// Controls if the timeFrom or timeShift overrides are shown in the panel header
func (builder *PanelBuilder) HideTimeOverride(hideTimeOverride bool) *PanelBuilder {
	builder.internal.HideTimeOverride = &hideTimeOverride

	return builder
}

// Dynamically load the panel
func (builder *PanelBuilder) LibraryPanel(libraryPanel dashboard.LibraryPanelRef) *PanelBuilder {
	builder.internal.LibraryPanel = &libraryPanel

	return builder
}

// The display value for this field.  This supports template variables blank is auto
func (builder *PanelBuilder) DisplayName(displayName string) *PanelBuilder {
	if builder.internal.FieldConfig == nil {
		builder.internal.FieldConfig = &dashboard.FieldConfigSource{}
	}
	builder.internal.FieldConfig.Defaults.DisplayName = &displayName

	return builder
}

// Unit a field should use. The unit you select is applied to all fields except time.
// You can use the units ID availables in Grafana or a custom unit.
// Available units in Grafana: https://github.com/grafana/grafana/blob/main/packages/grafana-data/src/valueFormats/categories.ts
// As custom unit, you can use the following formats:
// `suffix:<suffix>` for custom unit that should go after value.
// `prefix:<prefix>` for custom unit that should go before value.
// `time:<format>` For custom date time formats type for example `time:YYYY-MM-DD`.
// `si:<base scale><unit characters>` for custom SI units. For example: `si: mF`. This one is a bit more advanced as you can specify both a unit and the source data scale. So if your source data is represented as milli (thousands of) something prefix the unit with that SI scale character.
// `count:<unit>` for a custom count unit.
// `currency:<unit>` for custom a currency unit.
func (builder *PanelBuilder) Unit(unit string) *PanelBuilder {
	if builder.internal.FieldConfig == nil {
		builder.internal.FieldConfig = &dashboard.FieldConfigSource{}
	}
	builder.internal.FieldConfig.Defaults.Unit = &unit

	return builder
}

// Specify the number of decimals Grafana includes in the rendered value.
// If you leave this field blank, Grafana automatically truncates the number of decimals based on the value.
// For example 1.1234 will display as 1.12 and 100.456 will display as 100.
// To display all decimals, set the unit to `String`.
func (builder *PanelBuilder) Decimals(decimals float64) *PanelBuilder {
	if builder.internal.FieldConfig == nil {
		builder.internal.FieldConfig = &dashboard.FieldConfigSource{}
	}
	builder.internal.FieldConfig.Defaults.Decimals = &decimals

	return builder
}

// The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
func (builder *PanelBuilder) Min(min float64) *PanelBuilder {
	if builder.internal.FieldConfig == nil {
		builder.internal.FieldConfig = &dashboard.FieldConfigSource{}
	}
	builder.internal.FieldConfig.Defaults.Min = &min

	return builder
}

// The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
func (builder *PanelBuilder) Max(max float64) *PanelBuilder {
	if builder.internal.FieldConfig == nil {
		builder.internal.FieldConfig = &dashboard.FieldConfigSource{}
	}
	builder.internal.FieldConfig.Defaults.Max = &max

	return builder
}

// Convert input values into a display string
func (builder *PanelBuilder) Mappings(mappings []dashboard.ValueMapping) *PanelBuilder {
	if builder.internal.FieldConfig == nil {
		builder.internal.FieldConfig = &dashboard.FieldConfigSource{}
	}
	builder.internal.FieldConfig.Defaults.Mappings = mappings

	return builder
}

// Map numeric values to states
func (builder *PanelBuilder) Thresholds(thresholds cog.Builder[dashboard.ThresholdsConfig]) *PanelBuilder {
	if builder.internal.FieldConfig == nil {
		builder.internal.FieldConfig = &dashboard.FieldConfigSource{}
	}
	thresholdsResource, err := thresholds.Build()
	if err != nil {
		builder.errors["fieldConfig.defaults.thresholds"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.FieldConfig.Defaults.Thresholds = &thresholdsResource

	return builder
}

// Alternative to empty string
func (builder *PanelBuilder) NoValue(noValue string) *PanelBuilder {
	if builder.internal.FieldConfig == nil {
		builder.internal.FieldConfig = &dashboard.FieldConfigSource{}
	}
	builder.internal.FieldConfig.Defaults.NoValue = &noValue

	return builder
}

// Defaults are the options applied to all fields.
func (builder *PanelBuilder) Defaults(defaults dashboard.FieldConfig) *PanelBuilder {
	if builder.internal.FieldConfig == nil {
		builder.internal.FieldConfig = &dashboard.FieldConfigSource{}
	}
	builder.internal.FieldConfig.Defaults = defaults

	return builder
}

// Overrides are the options applied to specific fields overriding the defaults.
func (builder *PanelBuilder) WithOverride(matcher dashboard.MatcherConfig, properties []dashboard.DynamicConfigValue) *PanelBuilder {
	if builder.internal.FieldConfig == nil {
		builder.internal.FieldConfig = &dashboard.FieldConfigSource{}
	}
	builder.internal.FieldConfig.Overrides = append(builder.internal.FieldConfig.Overrides, struct {
		Matcher    dashboard.MatcherConfig        `json:"matcher"`
		Properties []dashboard.DynamicConfigValue `json:"properties"`
	}{
		Matcher:    matcher,
		Properties: properties,
	})

	return builder
}

func (builder *PanelBuilder) KeepTime(keepTime bool) *PanelBuilder {
	if builder.internal.Options == nil {
		builder.internal.Options = &Options{}
	}
	builder.internal.Options.(*Options).KeepTime = keepTime

	return builder
}

func (builder *PanelBuilder) IncludeVars(includeVars bool) *PanelBuilder {
	if builder.internal.Options == nil {
		builder.internal.Options = &Options{}
	}
	builder.internal.Options.(*Options).IncludeVars = includeVars

	return builder
}

func (builder *PanelBuilder) ShowStarred(showStarred bool) *PanelBuilder {
	if builder.internal.Options == nil {
		builder.internal.Options = &Options{}
	}
	builder.internal.Options.(*Options).ShowStarred = showStarred

	return builder
}

func (builder *PanelBuilder) ShowRecentlyViewed(showRecentlyViewed bool) *PanelBuilder {
	if builder.internal.Options == nil {
		builder.internal.Options = &Options{}
	}
	builder.internal.Options.(*Options).ShowRecentlyViewed = showRecentlyViewed

	return builder
}

func (builder *PanelBuilder) ShowSearch(showSearch bool) *PanelBuilder {
	if builder.internal.Options == nil {
		builder.internal.Options = &Options{}
	}
	builder.internal.Options.(*Options).ShowSearch = showSearch

	return builder
}

func (builder *PanelBuilder) ShowHeadings(showHeadings bool) *PanelBuilder {
	if builder.internal.Options == nil {
		builder.internal.Options = &Options{}
	}
	builder.internal.Options.(*Options).ShowHeadings = showHeadings

	return builder
}

func (builder *PanelBuilder) MaxItems(maxItems int64) *PanelBuilder {
	if builder.internal.Options == nil {
		builder.internal.Options = &Options{}
	}
	builder.internal.Options.(*Options).MaxItems = maxItems

	return builder
}

func (builder *PanelBuilder) Query(query string) *PanelBuilder {
	if builder.internal.Options == nil {
		builder.internal.Options = &Options{}
	}
	builder.internal.Options.(*Options).Query = query

	return builder
}

// folderId is deprecated, and migrated to folderUid on panel init
func (builder *PanelBuilder) FolderId(folderId int64) *PanelBuilder {
	if builder.internal.Options == nil {
		builder.internal.Options = &Options{}
	}
	builder.internal.Options.(*Options).FolderId = &folderId

	return builder
}

func (builder *PanelBuilder) FolderUID(folderUID string) *PanelBuilder {
	if builder.internal.Options == nil {
		builder.internal.Options = &Options{}
	}
	builder.internal.Options.(*Options).FolderUID = &folderUID

	return builder
}

func (builder *PanelBuilder) applyDefaults() {
	builder.Transparent(false)
	builder.Height(9)
	builder.Span(12)
	builder.KeepTime(false)
	builder.IncludeVars(false)
	builder.ShowStarred(true)
	builder.ShowRecentlyViewed(false)
	builder.ShowSearch(false)
	builder.ShowHeadings(true)
	builder.MaxItems(10)
	builder.Query("")
}