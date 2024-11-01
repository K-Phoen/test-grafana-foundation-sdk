// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExprTypeResampleResultAssertions] = (*ExprTypeResampleResultAssertionsBuilder)(nil)

type ExprTypeResampleResultAssertionsBuilder struct {
	internal *ExprTypeResampleResultAssertions
	errors   map[string]cog.BuildErrors
}

func NewExprTypeResampleResultAssertionsBuilder() *ExprTypeResampleResultAssertionsBuilder {
	resource := &ExprTypeResampleResultAssertions{}
	builder := &ExprTypeResampleResultAssertionsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ExprTypeResampleResultAssertionsBuilder) Build() (ExprTypeResampleResultAssertions, error) {
	if err := builder.internal.Validate(); err != nil {
		return ExprTypeResampleResultAssertions{}, err
	}

	return *builder.internal, nil
}

// Maximum frame count
func (builder *ExprTypeResampleResultAssertionsBuilder) MaxFrames(maxFrames int64) *ExprTypeResampleResultAssertionsBuilder {
	builder.internal.MaxFrames = &maxFrames

	return builder
}

// Type asserts that the frame matches a known type structure.
// Possible enum values:
//   - `""`
//   - `"timeseries-wide"`
//   - `"timeseries-long"`
//   - `"timeseries-many"`
//   - `"timeseries-multi"`
//   - `"directory-listing"`
//   - `"table"`
//   - `"numeric-wide"`
//   - `"numeric-multi"`
//   - `"numeric-long"`
//   - `"log-lines"`
func (builder *ExprTypeResampleResultAssertionsBuilder) Type(typeArg TypeResampleType) *ExprTypeResampleResultAssertionsBuilder {
	builder.internal.Type = &typeArg

	return builder
}

// TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
// contract documentation https://grafana.github.io/dataplane/contract/.
func (builder *ExprTypeResampleResultAssertionsBuilder) TypeVersion(typeVersion []int64) *ExprTypeResampleResultAssertionsBuilder {
	builder.internal.TypeVersion = typeVersion

	return builder
}

func (builder *ExprTypeResampleResultAssertionsBuilder) applyDefaults() {
}