package registries

// RFC5545 section  https://datatracker.ietf.org/doc/html/rfc5545#section-8.3.7

type ParticipantStatus string

const (
	NEEDACTION ParticipantStatus = "NEEDS-ACTION"
	ACCEPTED   ParticipantStatus = "ACCEPTED"
	DECLINED   ParticipantStatus = "DECLINED"
	TENTATIVE  ParticipantStatus = "TENTATIVE"
	DELEGATED  ParticipantStatus = "DELEGATED"
	COMPLETED  ParticipantStatus = "COMPLETED"
	INPROCESS  ParticipantStatus = "IN-PROCESS"
)
