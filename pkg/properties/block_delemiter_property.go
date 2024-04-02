package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.3

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type BlockDelimiterProperty interface {
	TextPropertyType
}

func NewBlockDelimiterProperty(
	block registries.PropertyNames,
	component Component,
) BlockDelimiterProperty {
	return &textPropertyType{
		PropName: block,
		Value:    types.NewTextValue(string(component)),
	}
}

type Component string

const (
	VCALENDAR Component = "VCALENDAR"
	VEVENT    Component = "VEVENT"
	VTODO     Component = "VTODO"
	VJOURNAL  Component = "VJOURNAL"
	VFREEBUSY Component = "VFREEBUSY"
	VTIMEZONE Component = "VTIMEZONE"
	VALARM    Component = "VALARM"
	STANDARD  Component = "STANDARD"
	DAYLIGHT  Component = "DAYLIGHT"
)
