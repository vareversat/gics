package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.3

import (
	"time"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type LastModifiedProperty interface {
	DateTimePropertyType
}

func NewLastModifiedProperty(value time.Time) LastModifiedProperty {
	return &dateTimePropertyType{
		PropName: registries.LASTMODIFIED,
		Value:    types.NewDateTimeValue(value, types.WithUtcTime),
	}
}

func NewLastModifiedPropertyFromString(value string) LastModifiedProperty {
	return &dateTimePropertyType{
		PropName: registries.LASTMODIFIED,
		Value:    types.NewDateTimeValueFromString(value, types.WithUtcTime),
	}
}
