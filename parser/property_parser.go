package parser

import (
	"fmt"

	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/properties"
	"github.com/vareversat/gics/types"
)

// ParseProperty will try to parse a property based on rawPropertyName with the associated rawPropertyValue
func ParseProperty(
	propertyName string,
	propertyValue string,
	params ...parameters.Parameter,
) (properties.Property, error) {
	switch propertyName {
	case "ACTION":
		return properties.NewActionPropertyFromString(propertyValue, params...), nil
	case "ATTACH":
		return properties.NewAttachmentProperty(propertyValue, params...), nil
	case "ATTENDEE":
		attendee, err := properties.NewAttendeePropertyFromString(propertyValue, params...)
		if err != nil {
			panic(err)
		}
		return attendee, nil
	case "BEGIN":
		return properties.NewBeginPropertyFromString(
			propertyValue,
		), nil
	case "CATEGORIES":
		return properties.NewCategoryPropertyFromString(propertyValue, params...), nil
	case "CALSCALE":
		return properties.NewCalScaleProperty(params...), nil
	case "CLASS":
		return properties.NewClassificationPropertyFromString(propertyValue, params...), nil
	case "COMMENT":
		return properties.NewCommentProperty(propertyValue, params...), nil
	case "COMPLETED_PROP":
		return properties.NewDateTimeCompletedPropertyFromString(propertyValue, params...), nil
	case "CONTACT":
		return properties.NewContactProperty(propertyValue, params...), nil
	case "CREATED":
		return properties.NewDateTimeCreatedPropertyFromString(propertyValue, params...), nil
	case "DESCRIPTION":
		return properties.NewDescriptionProperty(propertyValue, params...), nil
	case "DTEND":
		return properties.NewDateTimeEndPropertyFromString(
			propertyValue,
			types.WithUtcTime,
			params...), nil
	case "DTSTART":
		return properties.NewDateTimeStartPropertyFromString(
			propertyValue,
			types.WithUtcTime,
			params...), nil
	case "DTSTAMP":
		return properties.NewDateTimeStampPropertyFromString(propertyValue, params...), nil
	case "DURATION":
		return properties.NewDurationProperty(propertyValue), nil
	case "DUE":
		return properties.NewDateTimeDuePropertyFromString(
			propertyValue,
			types.WithUtcTime,
			params...), nil
	case "END":
		return properties.NewEndPropertyFromString(
			propertyValue,
		), nil
	case "EXDATE":
		return properties.NewExceptionDateTimePropertyFromString(
			propertyValue,
			types.WithUtcTime, params...,
		), nil
	case "EXRULE":
		fmt.Printf("### omitting %s\n", propertyName)
	case "FREEBUSY":
		fmt.Printf("### omitting %s\n", propertyName)
	case "GEO":
		geo, err := properties.NewGeographicPositionPropertyFromString(propertyValue, params...)
		if err != nil {
			panic(err)
		}
		return geo, nil
	case "LAST-MODIFIED":
		return properties.NewLastModifiedPropertyFromString(propertyValue), nil
	case "LOCATION":
		return properties.NewLocationProperty(propertyValue, params...), nil
	case "METHOD":
		return properties.NewMethodProperty(propertyValue, params...), nil
	case "ORGANIZER":
		organize, err := properties.NewOrganizerPropertyFromString(propertyValue, params...)
		if err != nil {
			panic(err)
		}
		return organize, nil
	case "PERCENTCOMPLETE":
		percentage, err := properties.NewPercentCompletePropertyFromString(propertyValue, params...)
		if err != nil {
			panic(err)
		}
		return percentage, nil
	case "PRIORITY":
		priority, err := properties.NewPriorityPropertyFromString(propertyValue, params...)
		if err != nil {
			panic(err)
		}
		return priority, nil
	case "PROP_TZID":
		return properties.NewTimeZoneIdProperty(propertyValue), nil
	case "PRODID":
		return properties.NewProductIdProperty(propertyValue, params...), nil
	case "RDATE":
		fmt.Printf("### omitting %s\n", propertyName)
	case "RECURRENCE-ID":
		return properties.NewRecurrenceIdPropertyFromString(
			propertyValue,
			types.WithUtcTime,
			params...), nil
	case "RELATED-TO":
		return properties.NewRelatedToProperty(propertyValue, params...), nil
	case "REPEAT":
		repeat, err := properties.NewRepeatCountPropertyFromString(propertyValue, params...)
		if err != nil {
			panic(err)
		}
		return repeat, nil
	case "REQUEST-STATUS":
		status, err := properties.NewRequestStatusPropertyFromString(propertyValue, params...)
		if err != nil {
			panic(err)
		}
		return status, nil
	case "RESOURCES":
		return properties.NewResourcesPropertyFromString(propertyValue, params...), nil
	case "RRULE":
		fmt.Printf("### omitting %s\n", propertyName)
	case "SEQUENCE":
		sequence, err := properties.NewSequencePropertyFromString(propertyValue, params...)
		if err != nil {
			panic(err)
		}
		return sequence, nil
	case "STATUS":
		return properties.NewStatusPropertyFromString(propertyValue, params...), nil
	case "SUMMARY":
		return properties.NewSummaryProperty(propertyValue, params...), nil
	case "TRANSP":
		return properties.NewTimeTransparencyPropertyFromString(propertyValue, params...), nil
	case "TRIGGER":
		fmt.Printf("### omitting %s\n", propertyName)
	case "TZNAME":
		return properties.NewTimeZoneNameProperty(propertyValue, params...), nil
	case "TZOFFSETFROM":
		return properties.NewTimeZoneOffsetFromProperty(propertyValue, params...), nil
	case "TZOFFSETTO":
		return properties.NewTimeZoneOffsetToProperty(propertyValue, params...), nil
	case "TZURL":
		url, err := properties.NewTimeZoneUrlPropertyFromString(propertyValue, params...)
		if err != nil {
			panic(err)
		}
		return url, nil
	case "UID":
		return properties.NewUidProperty(propertyValue, params...), nil
	case "URL":
		url, err := properties.NewUrlPropertyFromString(propertyValue, params...)
		if err != nil {
			panic(err)
		}
		return url, nil
	case "VERSION":
		return properties.NewVersionProperty(propertyValue, params...), nil
	// Non standard properties
	default:
		return properties.NewNonStandardProperty(propertyName, propertyValue, params...), nil
	}
	return nil, nil
}
