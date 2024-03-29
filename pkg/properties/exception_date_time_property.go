package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.5.1

import (
	"time"

	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type ExceptionDateTimeProperty interface {
}

type exceptionDateTimeProperty struct {
	IANAToken     registries.Properties
	DateTimeValue values.DateTimeValue // Either Datetime or Date
	DateValue     values.DateValue
}

func NewExceptionDateTimeProperty(value *time.Time) ExceptionDateTimeProperty {
	return &exceptionDateTimeProperty{
		IANAToken:     registries.EXDATE,
		DateTimeValue: values.NewDateTimeValue(value),
		DateValue:     values.NewDateValue(value),
	}
}
