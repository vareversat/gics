package registries

// ParameterRegistry contains all parameter names of the [RFC-5545]
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-8.3.3
type ParameterRegistry string

const (
	// AlternateTextRepresentation [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.1
	AlternateTextRepresentation ParameterRegistry = "ALTREP"
	// CommonName [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.2
	CommonName ParameterRegistry = "CN"
	// CalendarUserType [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.3
	CalendarUserType ParameterRegistry = "CUTYPE"
	// DelegatedFrom [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.4
	DelegatedFrom ParameterRegistry = "DELEGATED-FROM"
	// DelegatedTo [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.5
	DelegatedTo ParameterRegistry = "DELEGATED-TO"
	// DirectoryEntryReference [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.6
	DirectoryEntryReference ParameterRegistry = "DIR"
	// Encoding [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.7
	Encoding ParameterRegistry = "ENCODING"
	// FormatType [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.8
	FormatType ParameterRegistry = "FMTTYPE"
	// FreeBusyTimeType [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.9
	FreeBusyTimeType ParameterRegistry = "FBTYPE"
	// Language [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.10
	Language ParameterRegistry = "LANGUAGE"
	// Member [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.11
	Member ParameterRegistry = "MEMBER"
	// ParticipationStatus [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.12
	ParticipationStatus ParameterRegistry = "PARTSTAT"
	// Range [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.13
	Range ParameterRegistry = "RANGE"
	// Related [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.14
	Related ParameterRegistry = "RELATED"
	// RelationshipType [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.15
	RelationshipType ParameterRegistry = "RELTYPE"
	// Role [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.16
	Role ParameterRegistry = "ROLE"
	// RsvpExpectation [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.17
	RsvpExpectation ParameterRegistry = "RSVP"
	// SentBy [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.18
	SentBy ParameterRegistry = "SENT-BY"
	// TimeZoneIdProperty [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.19
	TimeZoneId ParameterRegistry = "TZID"
	// Value [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.20
	Value ParameterRegistry = "VALUE"
)
