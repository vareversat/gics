package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.1

import (
	"time"

	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type DateTimeCreatedProperty interface {
}

type dateTimeCreatedProperty struct {
	IANAToken registries.Properties
	Value     values.DateTimeValue
}

func NewDateTimeCreatedProperty(value *time.Time) DateTimeCreatedProperty {
	return &dateTimeCreatedProperty{
		IANAToken: registries.CREATED,
		Value:     values.NewDateTimeValue(value),
	}
}
