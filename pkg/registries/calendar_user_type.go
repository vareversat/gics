package registries

// RFC5545 section https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.3

type CalendarUserType string

const (
	GROUP      CalendarUserType = "GROUP"
	INDIVIDUAL CalendarUserType = "INDIVIDUAL"
	RESOURCE   CalendarUserType = "RESOURCE"
	ROOM       CalendarUserType = "ROOM"
	UNKNOWN    CalendarUserType = "UNKNOWN"
)
