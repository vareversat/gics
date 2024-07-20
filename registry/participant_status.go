package registry

// ParticipantStatusRegistry contains all participant status names of the [RFC-5545]
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-8.3.7
type ParticipantStatusRegistry string

const (
	// NeedsAction (used for registry.Vtodo component) [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.12
	NeedsAction ParticipantStatusRegistry = "NEEDS-ACTION"
	// Accepted (used for registry.Vtodo component) [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.12
	Accepted ParticipantStatusRegistry = "ACCEPTED"
	// Declined (used for registry.Vtodo component) [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.12
	Declined ParticipantStatusRegistry = "DECLINED"
	// Tentative (used for registry.Vevent component) [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.12
	Tentative ParticipantStatusRegistry = "TENTATIVE"
	// Delegated (used for registry.Vevent component) [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.12
	Delegated ParticipantStatusRegistry = "DELEGATED"
	// Completed (used for registry.Vevent component) [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.12
	Completed ParticipantStatusRegistry = "COMPLETED"
	// InProgress (used for registry.Vevent component) [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.12
	InProgress ParticipantStatusRegistry = "IN-PROGRESS"
)
