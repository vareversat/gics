package registries

// AlarmTriggerRelationshipType contains all component names of the [RFC-5545]
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.14
type AlarmTriggerRelationshipType string

const (
	// Start [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.14
	Start AlarmTriggerRelationshipType = "START"
	// End [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.14
	End AlarmTriggerRelationshipType = "END"
)
