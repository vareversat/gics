package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.4

import (
	"time"

	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type RecurrenceIdProperty interface {
}

type recurrenceIdProperty struct {
	IANAToken     registries.Properties
	DateTimeValue values.DateTimeValue // Either Datetime or Date
	DateValue     values.DateValue
}

func NewRecurrenceIdProperty(value *time.Time) RecurrenceIdProperty {
	return &recurrenceIdProperty{
		IANAToken:     registries.RECURRENCEID,
		DateTimeValue: values.NewDateTimeValue(value),
		DateValue:     values.NewDateValue(value),
	}
}
