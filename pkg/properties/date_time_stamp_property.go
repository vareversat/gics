package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.2

import (
	"time"

	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type DateTimeStampProperty interface {
}

type dateTimeStampProperty struct {
	IANAToken registries.Properties
	Value     values.DateTimeValue
}

func NewDateTimeStampProperty(value *time.Time) DateTimeStampProperty {
	return &dateTimeStampProperty{
		IANAToken: registries.DTSTAMP,
		Value:     values.NewDateTimeValue(value),
	}
}
