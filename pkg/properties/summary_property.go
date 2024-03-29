package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.12

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type SummaryProperty interface {
}

type summaryProperty struct {
	IANAToken registries.Properties
	Value     values.TextValue
}

func NewSummaryProperty(value string) SummaryProperty {
	return &summaryProperty{
		IANAToken: registries.SUMMARY,
		Value:     values.NewTextValue(value),
	}
}
