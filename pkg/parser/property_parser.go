package parser

import (
	"fmt"

	"github.com/vareversat/gics/pkg/properties"
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

// ParseProperty will try to parse a property based on rawPropertyName with the associated rawPropertyValue
func ParseProperty(propertyName string, propertyValue string) (properties.Property, error) {
	switch propertyName {
	case "ACTION":
		return properties.NewActionPropertyFromString(propertyValue), nil
	case "ATTACH":
		return properties.NewAttachmentProperty(propertyValue), nil
	case "ATTENDEE":
		attendee, err := properties.NewAttendeePropertyFromString(propertyValue)
		if err != nil {
			panic(err)
		}
		return attendee, nil
	case "BEGIN":
		return properties.NewBlockDelimiterPropertyFromString(
			registries.BEGIN,
			propertyValue,
		), nil
	case "CATEGORIES":
		return properties.NewCategoryPropertyFromString(propertyValue), nil
	case "CALSCALE":
		return properties.NewCalScaleProperty(), nil
	case "CLASS":
		return properties.NewClassificationPropertyFromString(propertyValue), nil
	case "COMMENT":
		return properties.NewCommentProperty(propertyValue), nil
	case "COMPLETED_PROP":
		panic("Traitement pour COMPLETED_PROP")
	case "CONTACT":
		return properties.NewContactProperty(propertyValue), nil
	case "CREATED":
		panic("Traitement pour CREATED")
	case "DESCRIPTION":
		return properties.NewDescriptionProperty(propertyValue), nil
	case "DTEND":
		return properties.NewDateTimeEndPropertyFromString(propertyValue, types.WithUtcTime), nil
	case "DTSTART":
		panic("Traitement pour DTSTART")
	case "DTSTAMP":
		panic("Traitement pour DTSTAMP")
	case "DURATION":
		return properties.NewDurationProperty(propertyValue), nil
	case "DUE":
		panic("Traitement pour DUE")
	case "END":
		return properties.NewBlockDelimiterPropertyFromString(registries.END, propertyValue), nil
	case "EXDATE":
		return properties.NewExceptionDateTimePropertyFromString(
			propertyValue,
			types.WithUtcTime,
		), nil
	case "EXRULE":
		panic("Traitement pour EXRULE")
	case "FREEBUSY":
		panic("Traitement pour FREEBUSY")
	case "GEO":
		geo, err := properties.NewGeographicPositionPropertyFromString(propertyValue)
		if err != nil {
			panic(err)
		}
		return geo, nil
	case "LAST-MODIFIED":
		return properties.NewLastModifiedPropertyFromString(propertyValue), nil
	case "LOCATION":
		return properties.NewLocationProperty(propertyValue), nil
	case "METHOD":
		return properties.NewMethodProperty(propertyValue), nil
	case "ORGANIZER":
		organize, err := properties.NewOrganizerPropertyFromString(propertyValue)
		if err != nil {
			panic(err)
		}
		return organize, nil
	case "PERCENTCOMPLETE":
		percentage, err := properties.NewPercentCompletePropertyFromString(propertyValue)
		if err != nil {
			panic(err)
		}
		return percentage, nil
	case "PRIORITY":
		priority, err := properties.NewPriorityPropertyFromString(propertyValue)
		if err != nil {
			panic(err)
		}
		return priority, nil
	case "PROP_TZID":
		return properties.NewTimeZoneIdProperty(propertyValue), nil
	case "PRODID":
		return properties.NewProductIdProperty(propertyValue), nil
	case "RDATE":
		fmt.Println("RDATE TODO")
		return nil, nil
	case "RECURRENCE-ID":
		return properties.NewRecurrenceIdPropertyFromString(propertyValue, types.WithUtcTime), nil
	case "RELATED-TO":
		return properties.NewRelatedToProperty(propertyValue), nil
	case "REPEAT":
		repeat, err := properties.NewRepeatCountPropertyFromString(propertyValue)
		if err != nil {
			panic(err)
		}
		return repeat, nil
	case "REQUEST-STATUS":
		status, err := properties.NewRequestStatusPropertyFromString(propertyValue)
		if err != nil {
			panic(err)
		}
		return status, nil
	case "RESOURCES":
		return properties.NewResourcesPropertyFromString(propertyValue), nil
	case "RRULE":
		fmt.Println("RRULE TODO")
		return nil, nil
	case "SEQUENCE":
		sequence, err := properties.NewSequencePropertyFromString(propertyValue)
		if err != nil {
			panic(err)
		}
		return sequence, nil
	case "STATUS":
		return properties.NewStatusPropertyFromString(propertyValue), nil
	case "SUMMARY":
		return properties.NewSummaryProperty(propertyValue), nil
	case "TRANSP":
		return properties.NewTimeTransparencyPropertyFromString(propertyValue), nil
	case "TRIGGER":
		panic("Traitement pour TRIGGER")
	case "TZNAME":
		return properties.NewTimeZoneNameProperty(propertyValue), nil
	case "TZOFFSETFROM":
		return properties.NewTimeZoneOffsetFromProperty(propertyValue), nil
	case "TZOFFSETTO":
		return properties.NewTimeZoneOffsetToProperty(propertyValue), nil
	case "TZURL":
		url, err := properties.NewTimeZoneUrlPropertyFromString(propertyValue)
		if err != nil {
			panic(err)
		}
		return url, nil
	case "UID":
		properties.NewUidProperty(propertyValue)
	case "URL":
		url, err := properties.NewUrlPropertyFromString(propertyValue)
		if err != nil {
			panic(err)
		}
		return url, nil
	case "VERSION":
		return properties.NewVersionProperty(propertyValue), nil
	default:
		panic(propertyName)
	}
	return nil, nil
}
