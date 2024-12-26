package parser

import (
	"fmt"

	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/properties"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

// ParseProperty will try to parse a property based on rawPropertyName with the associated rawPropertyValue
func ParseProperty(
	propertyName string,
	propertyValue string,
	params ...parameters.Parameter,
) (properties.Property, error) {
	switch propertyName {
	case string(registries.ActionProp):
		return properties.NewActionPropertyFromString(propertyValue, params...), nil
	case string(registries.AttachmentProp):
		return properties.NewAttachmentProperty(propertyValue, params...), nil
	case string(registries.AttendeeProp):
		attendee, err := properties.NewAttendeePropertyFromString(propertyValue, params...)
		if err != nil {
			panic(err)
		}
		return attendee, nil
	case string(registries.BeginProp):
		return properties.NewBeginPropertyFromString(
			propertyValue,
		), nil
	case string(registries.CategoriesProp):
		return properties.NewCategoryPropertyFromString(propertyValue, params...), nil
	case string(registries.CalendarScaleProp):
		return properties.NewCalScaleProperty(params...), nil
	case string(registries.ClassProp):
		return properties.NewClassificationPropertyFromString(propertyValue, params...), nil
	case string(registries.CommentProp):
		return properties.NewCommentProperty(propertyValue, params...), nil
	case string(registries.CompletedProp):
		return properties.NewDateTimeCompletedPropertyFromString(propertyValue, params...), nil
	case string(registries.ContactProp):
		return properties.NewContactProperty(propertyValue, params...), nil
	case string(registries.DateTimeCreatedProp):
		return properties.NewDateTimeCreatedPropertyFromString(propertyValue, params...), nil
	case string(registries.DescriptionProp):
		return properties.NewDescriptionProperty(propertyValue, params...), nil
	case string(registries.DateTimeEndProp):
		return properties.NewDateTimeEndPropertyFromString(
			propertyValue,
			params...), nil
	case string(registries.DateTimeStartProp):
		return properties.NewDateTimeStartPropertyFromString(
			propertyValue,
			params...), nil
	case string(registries.DateTimeStampProp):
		return properties.NewDateTimeStampPropertyFromString(propertyValue, params...), nil
	case string(registries.DurationProp):
		return properties.NewDurationProperty(propertyValue)
	case string(registries.DateTimeDueProp):
		return properties.NewDateTimeDuePropertyFromString(
			propertyValue,
			params...), nil
	case string(registries.EndProp):
		return properties.NewEndPropertyFromString(
			propertyValue,
		), nil
	case string(registries.ExceptionDateTimesProp):
		return properties.NewExceptionDateTimePropertyFromString(
			propertyValue, params...,
		), nil
	case string(registries.ExceptionRuleProp):
		fmt.Printf("### omitting %s\n", propertyName)
	case string(registries.FreeBusyTimeProp):
		fmt.Printf("### omitting %s\n", propertyName)
	case string(registries.GeoProp):
		geo, err := properties.NewGeographicPositionPropertyFromString(propertyValue, params...)
		if err != nil {
			panic(err)
		}
		return geo, nil
	case string(registries.LastModifiedProp):
		return properties.NewLastModifiedPropertyFromString(propertyValue), nil
	case string(registries.LocationProp):
		return properties.NewLocationProperty(propertyValue, params...), nil
	case string(registries.MethodProp):
		return properties.NewMethodProperty(propertyValue, params...), nil
	case string(registries.OrganizerProp):
		organize, err := properties.NewOrganizerPropertyFromString(propertyValue, params...)
		if err != nil {
			panic(err)
		}
		return organize, nil
	case string(registries.PercentCompleteProp):
		percentage, err := properties.NewPercentCompletePropertyFromString(propertyValue, params...)
		if err != nil {
			panic(err)
		}
		return percentage, nil
	case string(registries.PriorityProp):
		priority, err := properties.NewPriorityPropertyFromString(propertyValue, params...)
		if err != nil {
			panic(err)
		}
		return priority, nil
	case string(registries.TimeZoneIdProp):
		return properties.NewTimeZoneIdProperty(propertyValue), nil
	case string(registries.ProductIdentifierProp):
		return properties.NewProductIdProperty(propertyValue, params...), nil
	case string(registries.RecurrenceDateTimesProp):
		fmt.Printf("### omitting %s\n", propertyName)
	case string(registries.RecurrenceIdProp):
		return properties.NewRecurrenceIdPropertyFromString(
			propertyValue,
			types.WithUtcTime,
			params...), nil
	case string(registries.RelatedToProp):
		return properties.NewRelatedToProperty(propertyValue, params...), nil
	case string(registries.RepeatProp):
		repeat, err := properties.NewRepeatCountPropertyFromString(propertyValue, params...)
		if err != nil {
			panic(err)
		}
		return repeat, nil
	case string(registries.RequestStatusProp):
		status, err := properties.NewRequestStatusPropertyFromString(propertyValue, params...)
		if err != nil {
			panic(err)
		}
		return status, nil
	case string(registries.ResourcesProp):
		return properties.NewResourcesPropertyFromString(propertyValue, params...), nil
	case string(registries.RecurrenceRuleProp):
		fmt.Printf("### omitting %s\n", propertyName)
	case string(registries.SequenceProp):
		sequence, err := properties.NewSequencePropertyFromString(propertyValue, params...)
		if err != nil {
			panic(err)
		}
		return sequence, nil
	case string(registries.StatusProp):
		return properties.NewStatusPropertyFromString(propertyValue, params...), nil
	case string(registries.SummaryProp):
		return properties.NewSummaryProperty(propertyValue, params...), nil
	case string(registries.TimeTransparencyProp):
		return properties.NewTimeTransparencyPropertyFromString(propertyValue, params...), nil
	case string(registries.TriggerProp):
		fmt.Printf("### omitting %s\n", propertyName)
	case string(registries.TimeZoneNameProp):
		return properties.NewTimeZoneNameProperty(propertyValue, params...), nil
	case string(registries.TimeZoneOffsetFromProp):
		return properties.NewTimeZoneOffsetFromProperty(propertyValue, params...), nil
	case string(registries.TimeZoneOffsetToProp):
		return properties.NewTimeZoneOffsetToProperty(propertyValue, params...), nil
	case string(registries.TimeZoneUrlProp):
		url, err := properties.NewTimeZoneUrlPropertyFromString(propertyValue, params...)
		if err != nil {
			panic(err)
		}
		return url, nil
	case string(registries.UidProp):
		return properties.NewUidProperty(propertyValue, params...), nil
	case string(registries.UrlProp):
		url, err := properties.NewUrlPropertyFromString(propertyValue, params...)
		if err != nil {
			panic(err)
		}
		return url, nil
	case string(registries.VersionProp):
		return properties.NewVersionProperty(propertyValue, params...), nil
	// Non standard properties
	default:
		return properties.NewNonStandardProperty(propertyName, propertyValue, params...), nil
	}
	return nil, nil
}
