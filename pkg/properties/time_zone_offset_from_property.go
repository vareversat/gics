package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.3.3

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type TimeZoneOffsetFrom interface {
}

type timeZoneOffsetFrom struct {
	IANAToken registries.Properties
	Value     values.TextValue
}

func NewTimeZoneOffsetFrom(value string) TimeZoneOffsetFrom {
	return &timeZoneOffsetTo{
		IANAToken: registries.TZOFFSETFROM,
		Value:     values.NewUtcOffsetValue(value),
	}
}
