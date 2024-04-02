package components

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.6.6

import (
	"github.com/vareversat/gics/pkg/properties"
	"github.com/vareversat/gics/pkg/registries"
)

type TimeZoneCalendarComponent interface {
}

type timeZoneCalendarComponent struct {
	// MANDATORY and MUST NOT occur more than once
	Begin      properties.BlockDelimiterProperty
	End        properties.BlockDelimiterProperty
	TimeZoneId properties.TimeZoneIdProperty

	// OPTIONAL and MUST NOT occur more than once
	LastModified properties.LastModifiedProperty
	TimeZoneUrl  properties.TimeZoneUrlProperty
}

type subBlock struct {
	// MANDATORY and MUST NOT occur more than once
	Begin              properties.BlockDelimiterProperty
	End                properties.BlockDelimiterProperty
	DateTimeStart      properties.DateTimeStartProperty
	TimeZoneOffsetTo   properties.TimeZoneOffsetToProperty
	TimeZoneOffsetFrom properties.TimeZoneOffsetFromProperty

	// OPTIONAL and SHOULD NOT occur more than once
	RecurrenceRule properties.RecurrenceRuleProperty

	// OPTIONAL and MUST NOT occur more than once
	Comment             []properties.AttachmentProperty
	RecurrenceDateTimes []properties.RecurrenceDateTimesProperty
	TimeZoneName        []properties.TimeZoneNameProperty
}

func NewTimeZoneCalendarComponent() TimeZoneCalendarComponent {
	return &timeZoneCalendarComponent{
		Begin: properties.NewBlockDelimiterProperty(registries.BEGIN, properties.VTIMEZONE),
		End:   properties.NewBlockDelimiterProperty(registries.END, properties.VTIMEZONE),
	}
}
