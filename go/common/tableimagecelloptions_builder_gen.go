// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TableImageCellOptions] = (*TableImageCellOptionsBuilder)(nil)

// Json view cell options
type TableImageCellOptionsBuilder struct {
	internal *TableImageCellOptions
	errors   map[string]cog.BuildErrors
}

func NewTableImageCellOptionsBuilder() *TableImageCellOptionsBuilder {
	resource := NewTableImageCellOptions()
	builder := &TableImageCellOptionsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}
	builder.internal.Type = "image"

	return builder
}

func (builder *TableImageCellOptionsBuilder) Build() (TableImageCellOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return TableImageCellOptions{}, err
	}

	return *builder.internal, nil
}

func (builder *TableImageCellOptionsBuilder) Alt(alt string) *TableImageCellOptionsBuilder {
	builder.internal.Alt = &alt

	return builder
}

func (builder *TableImageCellOptionsBuilder) Title(title string) *TableImageCellOptionsBuilder {
	builder.internal.Title = &title

	return builder
}