package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.2

import (
	"time"

	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type DateTimeEndProperty interface {
}

type dateTimeEndProperty struct {
	IANAToken     registries.Properties
	DateTimeValue values.DateTimeValue // Either Datetime or Date
	DateValue     values.DateValue
}

func NewDateTimeEndProperty(value *time.Time) DateTimeEndProperty {
	return &dateTimeEndProperty{
		IANAToken:     registries.DTEND,
		DateTimeValue: values.NewDateTimeValue(value),
		DateValue:     values.NewDateValue(value),
	}
}
