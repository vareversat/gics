package registries

// ParticipantRoleRegistry contains all participant role names of the [RFC-5545]
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-8.3.9
type ParticipantRoleRegistry string

const (
	// Chair [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.16
	Chair ParticipantRoleRegistry = "CHAIR"
	// RequiredParticipant [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.16
	RequiredParticipant ParticipantRoleRegistry = "REQ-PARTICIPANT"
	// OptionalParticipant [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.16
	OptionalParticipant ParticipantRoleRegistry = "OPT-PARTICIPANT"
	// NonParticipant [See RFC-5545 spec]
	// [See RFC-5545 spec]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.16
	NonParticipant ParticipantRoleRegistry = "NON-PARTICIPANT"
)
