package registries

// ClassificationRegistry contains all calendar user type registry names of the [RFC-5545]
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-8.3.11
type ClassificationRegistry string

const (
	// Public [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.3
	Public ClassificationRegistry = "PUBLIC"
	// Private [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.3
	Private ClassificationRegistry = "PRIVATE"
	// Confidential [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.3
	Confidential ClassificationRegistry = "CONFIDENTIAL"
)
