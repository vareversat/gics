package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.12

import (
	"github.com/vareversat/gics/pkg/parameters"
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type SummaryProperty interface {
	TextPropertyType
}

func NewSummaryProperty(
	summaryValue string, params ...parameters.Parameter,
) SummaryProperty {
	return &textPropertyType{
		PropName:   registries.SUMMARY,
		Value:      types.NewTextValue(summaryValue),
		Parameters: params,
	}
}
