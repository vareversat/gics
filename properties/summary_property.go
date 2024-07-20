package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.12

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registry"
	"github.com/vareversat/gics/types"
)

type SummaryProperty interface {
	TextPropertyType
}

func NewSummaryProperty(
	summaryValue string, params ...parameters.Parameter,
) SummaryProperty {
	return &textPropertyType{
		PropName:   registry.SUMMARY,
		Value:      types.NewTextValue(summaryValue),
		Parameters: params,
	}
}
