package registries

// ParameterRegistry contains all parameter names of the [RFC-5545]
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-8.3.3
type ParameterRegistry string

const (
	// AlternateTextRepresentationParam [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.1
	AlternateTextRepresentationParam ParameterRegistry = "ALTREP"
	// CommonNameParam [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.2
	CommonNameParam ParameterRegistry = "CN"
	// CalendarUserTypeParam [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.3
	CalendarUserTypeParam ParameterRegistry = "CUTYPE"
	// DelegatedFromParam [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.4
	DelegatedFromParam ParameterRegistry = "DELEGATED-FROM"
	// DelegatedToParam [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.5
	DelegatedToParam ParameterRegistry = "DELEGATED-TO"
	// DirectoryEntryReferenceParam [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.6
	DirectoryEntryReferenceParam ParameterRegistry = "DIR"
	// EncodingParam [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.7
	EncodingParam ParameterRegistry = "ENCODING"
	// FormatTypeParam [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.8
	FormatTypeParam ParameterRegistry = "FMTTYPE"
	// FreeBusyTimeTypeParam [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.9
	FreeBusyTimeTypeParam ParameterRegistry = "FBTYPE"
	// LanguageParam [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.10
	LanguageParam ParameterRegistry = "LANGUAGE"
	// MemberParam [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.11
	MemberParam ParameterRegistry = "MEMBER"
	// ParticipationStatusParam [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.12
	ParticipationStatusParam ParameterRegistry = "PARTSTAT"
	// RecurrenceIdRangeParam [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.13
	RecurrenceIdRangeParam ParameterRegistry = "RANGE"
	// AlarmTriggerRelationshipParam [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.14
	AlarmTriggerRelationshipParam ParameterRegistry = "RELATED"
	// RelationshipTypeParam [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.15
	RelationshipTypeParam ParameterRegistry = "RELTYPE"
	// ParticipantRoleParam [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.16
	ParticipantRoleParam ParameterRegistry = "ROLE"
	// RsvpExpectationParam [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.17
	RsvpExpectationParam ParameterRegistry = "RSVP"
	// SentByParam [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.18
	SentByParam ParameterRegistry = "SENT-BY"
	// TimeZoneIdParam [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.19
	TimeZoneIdParam ParameterRegistry = "TZID"
	// ValueDataTypesParam [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.20
	ValueDataTypesParam ParameterRegistry = "VALUE"
	// NonStandardParam [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2
	NonStandardParam ParameterRegistry = "NON-STANDARD-PARAMETER"
)
