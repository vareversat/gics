package types

import "github.com/vareversat/gics/pkg/registries"

type StatusType string

const (
	// For VEVENT
	VEVENT_TENTATIVE StatusType = "TENTATIVE"
	VEVENT_CONFIRMED StatusType = "CONFIRMED"
	VEVENT_ANCELLED  StatusType = "CANCELLED"

	// For VTODO
	VTODO_NEEDSACTION StatusType = "NEEDS-ACTION"
	VTODO_COMPLETED   StatusType = "COMPLETED"
	VTODO_INPROGRESS  StatusType = "IN-PROCESS"
	VTODO_CANCELLED   StatusType = "CANCELLED"

	// For VJOURNAL
	VJOURNAL_DRAFT    StatusType = "DRAFT"
	VJOURNAL_FINAL    StatusType = "FINAL"
	VJOURNAL_CANCELED StatusType = "IN-PROCESS"
)

type StatusValue struct {
	V
	Value StatusType
}

func NewStatusValue(value StatusType) StatusValue {
	return StatusValue{
		V:     NewValue(registries.TEXT),
		Value: value,
	}
}
