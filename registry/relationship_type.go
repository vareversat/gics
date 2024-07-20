package registry

// RelationshipTypeRegistry contains all relationship names of the [RFC-5545]
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-8.3.8
type RelationshipTypeRegistry string

const (
	// Child [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.15
	Child ParticipantStatusRegistry = "CHILD"
	// Parent [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.15
	Parent ParticipantStatusRegistry = "PARENT"
	// Sibling [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.15
	Sibling ParticipantStatusRegistry = "SIBLING"
)
