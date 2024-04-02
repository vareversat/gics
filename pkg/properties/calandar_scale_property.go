package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.7.1

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type CalendarScaleProperty interface {
	TextPropertyType
}

func NewCalScaleProperty() CalendarScaleProperty {
	return &textPropertyType{
		PropName: registries.CALSCALE,
		Value: types.NewTextValue(
			"GREGORIAN",
		), // This is the default value of the CALSCALE property
	}
}
