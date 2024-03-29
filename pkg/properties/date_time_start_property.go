package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.2

import (
	"time"

	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type DateTimeStartProperty interface {
}

type dateTimeStartProperty struct {
	IANAToken     registries.Properties
	DateTimeValue values.DateTimeValue // Either Datetime or Date
	DateValue     values.DateValue
}

func NewDateTimeStartProperty(value *time.Time) DateTimeStartProperty {
	return &dateTimeStartProperty{
		IANAToken:     registries.DTSTART,
		DateTimeValue: values.NewDateTimeValue(value),
		DateValue:     values.NewDateValue(value),
	}
}
