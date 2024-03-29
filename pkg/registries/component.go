package registries

// RFC5545 section  https://datatracker.ietf.org/doc/html/rfc5545#section-8.3.1

type Component string

const (
	VCALENDAR Component = "VCALENDAR"
	VEVENT    Component = "VEVENT"
	VTODO     Component = "VTODO"
	VJOURNAL  Component = "VJOURNAL"
	VFREEBUSY Component = "VFREEBUSY"
	VTIMEZONE Component = "VTIMEZONE"
	VALARM    Component = "VALARM"
	STANDARD  Component = "STANDARD"
	DAYLIGHT  Component = "DAYLIGHT"
)
