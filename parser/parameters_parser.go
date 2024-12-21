package parser

import (
	"github.com/vareversat/gics/parameters"
)

// ParseParameter will try to parse a property based on parameterName with the associated parameterValue
func ParseParameter(parameterName string, parameterValue string) (parameters.Parameter, error) {
	switch parameterName {
	case "ALTREP":
		return parameters.NewAlternateTextRepresentationParamFromString(parameterValue)
	case "CN":
		return parameters.NewCommonNameParam(parameterValue), nil
	case "CUTYPE":
		return parameters.NewCalendarUserTypeParamFromString(parameterValue), nil
	case "DELEGATED-FROM":
		return parameters.NewDelegatedFromParamFromString(parameterValue)
	case "DELEGATED-TO":
		return parameters.NewDelegatedToParamFromString(parameterValue)
	case "DIR":
		return parameters.NewDirectoryEntryParamFromString(parameterValue)
	case "ENCODING":
		return parameters.NewInlineEncodingParamFromString(parameterValue), nil
	case "FBTYPE":
		return parameters.NewFreeBusyTimeParamFromString(parameterValue), nil
	case "FMTTYPE":
		return parameters.NewFormatTypeParam(parameterValue), nil
	case "LANGUAGE":
		return parameters.NewLanguageParam(parameterValue), nil
	case "MEMBER":
		return parameters.NewMemberParamFromString(parameterValue)
	case "PARTSTAT":
		return parameters.NewParticipationRoleParamFromString(parameterValue), nil
	case "RELATED":
		return parameters.NewAlarmTriggerRelationshipParamFromString(parameterValue), nil
	case "RANGE":
		return parameters.NewRangeParam(), nil
	case "RELTYPE":
		return parameters.NewRelationshipParamFromString(parameterValue), nil
	case "ROLE":
		return parameters.NewParticipationRoleParamFromString(parameterValue), nil
	case "RSVP":
		return parameters.NewRSVPParamFromString(parameterValue)
	case "SENT-BY":
		return parameters.NewSentByParamFromString(parameterValue)
	case "TZID":
		return parameters.NewTimeZoneIdentifierParam(parameterValue), nil
	case "VALUE":
		return parameters.NewValueDataTypesParamFromString(parameterValue), nil
	// Non standard properties
	default:
		return parameters.NewNonStandardParameter(parameterName, parameterValue), nil

	}
}
