package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.1

import (
	"time"

	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type DateTimeCompletedProperty interface {
}

type dateTimeCompletedProperty struct {
	IANAToken registries.Properties
	Value     values.DateTimeValue
}

func NewDateTimeCompletedProperty(value *time.Time) DateTimeCompletedProperty {
	return &dateTimeCompletedProperty{
		IANAToken: registries.COMPLETED_PROP,
		Value:     values.NewDateTimeValue(value),
	}
}
