package registry

// ParameterRegistry contains all parameter names of the [RFC-5545]
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-8.3.3
type ParameterRegistry string

const (
	// Altrep [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.1
	Altrep ParameterRegistry = "ALTREP"
	// Cn [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.2
	Cn ParameterRegistry = "CN"
	// CuType [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.3
	CuType ParameterRegistry = "CUTYPE"
	// DelegatedFrom [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.4
	DelegatedFrom ParameterRegistry = "DELEGATED-FROM"
	// DelegatedTo [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.5
	DelegatedTo ParameterRegistry = "DELEGATED-TO"
	// Dir [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.6
	Dir ParameterRegistry = "DIR"
	// Encoding [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.7
	Encoding ParameterRegistry = "ENCODING"
	// FmtType [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.8
	FmtType ParameterRegistry = "FMTTYPE"
	// FbType [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.9
	FbType ParameterRegistry = "FBTYPE"
	// Language [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.10
	Language ParameterRegistry = "LANGUAGE"
	// Member [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.11
	Member ParameterRegistry = "MEMBER"
	// PartStat [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.12
	PartStat ParameterRegistry = "PARTSTAT"
	// Range [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.13
	Range ParameterRegistry = "RANGE"
	// Related [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.14
	Related ParameterRegistry = "RELATED"
	// RelType [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.15
	RelType ParameterRegistry = "RELTYPE"
	// Role [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.16
	Role ParameterRegistry = "ROLE"
	// Rsvp [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.17
	Rsvp ParameterRegistry = "RSVP"
	// SentBy [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.18
	SentBy ParameterRegistry = "SENT-BY"
	// TZIDParameter [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.19
	TZIDParameter ParameterRegistry = "TZID"
	// Value [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.20
	Value ParameterRegistry = "VALUE"
)
