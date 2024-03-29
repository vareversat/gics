package registries

// RFC5545 section  https://https://datatracker.ietf.org/doc/html/rfc5545#section-8.3.9

type ParticipationRole string

const (
	CHAIR          ParticipationRole = "CHAIR"
	REQPARTICIPANT ParticipationRole = "REQ-PARTICIPANT"
	OPTPARTICIPANT ParticipationRole = "OPT-PARTICIPANT"
	NONPARTICIPANT ParticipationRole = "NON-PARTICIPANT"
)
