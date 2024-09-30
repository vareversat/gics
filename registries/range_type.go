package registries

// RangeTypeRegistry contains all participant role names of the [RFC-5545]
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.13
type RangeTypeRegistry string

const (
	// Deprecated: As stated in the [RFC-5545], registries.RecurrenceIdRangeParam  is deprecated
	// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.13
	ThisAndFuture RangeTypeRegistry = "THISANDFUTURE"
)
