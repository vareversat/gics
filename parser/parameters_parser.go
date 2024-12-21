package parser

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
)

// ParseParameter will try to parse a property based on parameterName with the associated parameterValue
func ParseParameter(parameterName string, parameterValue string) (parameters.Parameter, error) {
	switch parameterName {
	case string(registries.AlternateTextRepresentationParam):
		return parameters.NewAlternateTextRepresentationParamFromString(parameterValue)
	case string(registries.CommonNameParam):
		return parameters.NewCommonNameParam(parameterValue), nil
	case string(registries.CalendarUserTypeParam):
		return parameters.NewCalendarUserTypeParamFromString(parameterValue), nil
	case string(registries.DelegatedFromParam):
		return parameters.NewDelegatedFromParamFromString(parameterValue)
	case string(registries.DelegatedToParam):
		return parameters.NewDelegatedToParamFromString(parameterValue)
	case string(registries.DirectoryEntryReferenceParam):
		return parameters.NewDirectoryEntryParamFromString(parameterValue)
	case string(registries.EncodingParam):
		return parameters.NewInlineEncodingParamFromString(parameterValue), nil
	case string(registries.FreeBusyTimeTypeParam):
		return parameters.NewFreeBusyTimeParamFromString(parameterValue), nil
	case string(registries.FormatTypeParam):
		return parameters.NewFormatTypeParam(parameterValue), nil
	case string(registries.LanguageParam):
		return parameters.NewLanguageParam(parameterValue), nil
	case string(registries.MemberParam):
		return parameters.NewMemberParamFromString(parameterValue)
	case string(registries.ParticipationStatusParam):
		return parameters.NewParticipationRoleParamFromString(parameterValue), nil
	case string(registries.AlarmTriggerRelationshipParam):
		return parameters.NewAlarmTriggerRelationshipParamFromString(parameterValue), nil
	case string(registries.RecurrenceIdRangeParam):
		return parameters.NewRecurrenceIdRangeParam(), nil
	case string(registries.RelationshipTypeParam):
		return parameters.NewRelationshipParamFromString(parameterValue), nil
	case string(registries.ParticipantRoleParam):
		return parameters.NewParticipationRoleParamFromString(parameterValue), nil
	case string(registries.RsvpExpectationParam):
		return parameters.NewRSVPParamFromString(parameterValue)
	case string(registries.SentByParam):
		return parameters.NewSentByParamFromString(parameterValue)
	case string(registries.TimeZoneIdParam):
		return parameters.NewTimeZoneIdentifierParam(parameterValue)
	case string(registries.ValueDataTypesParam):
		return parameters.NewValueDataTypesParamFromString(parameterValue), nil
	// Non standard properties
	default:
		return parameters.NewNonStandardParameter(parameterName, parameterValue), nil

	}
}
