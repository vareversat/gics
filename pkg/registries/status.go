package registries

// RFC5545 section hhttps://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.11

type Status string

const (
	// For VEVENT
	VEVENT_TENTATIVE Status = "TENTATIVE"
	VEVENT_CONFIRMED Status = "CONFIRMED"
	VEVENT_ANCELLED  Status = "CANCELLED"

	// For VTODO
	VTODO_NEEDSACTION Status = "NEEDS-ACTION"
	VTODO_COMPLETED   Status = "COMPLETED"
	VTODO_INPROGRESS  Status = "IN-PROCESS"
	VTODO_CANCELLED   Status = "CANCELLED"

	// For VJOURNAL
	VJOURNAL_DRAFT    Status = "DRAFT"
	VJOURNAL_FINAL    Status = "FINAL"
	VJOURNAL_CANCELED Status = "IN-PROCESS"
)
