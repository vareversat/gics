package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.12

// Parameter used in these properties :
// - ATTENDEE

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type ParticipationStatusParam interface {
	TextParameter
}

func NewParticipationStatusParam(value ParticipantStatusType) ParticipationStatusParam {
	return &textParameter{
		ParamName: registries.PARTSTAT,
		Value:     types.NewTextValue(string(value)),
	}
}

type ParticipantStatusType string

const (
	NEEDACTION ParticipantStatusType = "NEEDS-ACTION"
	ACCEPTED   ParticipantStatusType = "ACCEPTED"
	DECLINED   ParticipantStatusType = "DECLINED"
	TENTATIVE  ParticipantStatusType = "TENTATIVE"
	DELEGATED  ParticipantStatusType = "DELEGATED"
	COMPLETED  ParticipantStatusType = "COMPLETED"
	INPROCESS  ParticipantStatusType = "IN-PROCESS"
)
