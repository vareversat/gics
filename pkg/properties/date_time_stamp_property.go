package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.2

import (
	"time"

	"github.com/vareversat/gics/pkg/parameters"

	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type DateTimeStampProperty interface {
	DateTimePropertyType
}

func NewDateTimeStampProperty(
	timeValue time.Time,
	format types.DTFormat,
	params ...parameters.Parameter,
) DateTimeStampProperty {
	return &dateTimePropertyType{
		PropName:   registries.DTSTAMP,
		Value:      types.NewDateTimeValue(timeValue, format),
		Parameters: params,
	}
}
