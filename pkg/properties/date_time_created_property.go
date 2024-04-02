package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.1

import (
	"time"

	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type DateTimeCreatedProperty interface {
	DateTimePropertyType
}

func NewDateTimeCreatedProperty(
	timeValue time.Time,
	format types.DTFormat,
) DateTimeCreatedProperty {
	return &dateTimePropertyType{
		PropName: registries.CREATED,
		Value:    types.NewDateTimeValue(timeValue, format),
	}
}
