package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.3.4

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type TimeZoneOffsetTo interface {
}

type timeZoneOffsetTo struct {
	IANAToken registries.Properties
	Value     values.TextValue
}

func NewTimeZoneOffsetTo(value string) TimeZoneOffsetTo {
	return &timeZoneOffsetTo{
		IANAToken: registries.TZOFFSETTO,
		Value:     values.NewUtcOffsetValue(value),
	}
}
