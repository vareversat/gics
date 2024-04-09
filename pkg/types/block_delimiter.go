package types

import "github.com/vareversat/gics/pkg/registries"

type BlockDelimiterType string

const (
	VCALENDAR BlockDelimiterType = "VCALENDAR"
	VEVENT    BlockDelimiterType = "VEVENT"
	VTODO     BlockDelimiterType = "VTODO"
	VJOURNAL  BlockDelimiterType = "VJOURNAL"
	VFREEBUSY BlockDelimiterType = "VFREEBUSY"
	VTIMEZONE BlockDelimiterType = "VTIMEZONE"
	VALARM    BlockDelimiterType = "VALARM"
	STANDARD  BlockDelimiterType = "STANDARD"
	DAYLIGHT  BlockDelimiterType = "DAYLIGHT"
)

type BlockDelimiterValue struct {
	V
	Value BlockDelimiterType
}

func NewBlockDelimiterValue(value BlockDelimiterType) BlockDelimiterValue {
	return BlockDelimiterValue{
		V:     NewValue(registries.TEXT),
		Value: value,
	}
}
