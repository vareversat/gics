package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.5

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type DurationProperty interface {
	DurationPropertyType
}

func NewDurationProperty(duration string) DurationProperty {
	return &durationPropertyType{
		PropName: registries.DURATION_PROP,
		Value:    types.NewDurationValue(duration),
	}
}
