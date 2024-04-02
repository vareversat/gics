package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.3

import (
	"time"

	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
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
