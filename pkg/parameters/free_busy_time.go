package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.9

// Parameter used in these properties :
// - FREEBUSY

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type FreeBusyTimeParam interface {
	TextParameter
}

func NewFreeBusyTimeParam(value FreeBusyTimeType) FreeBusyTimeParam {
	return &textParameter{
		ParamName: registries.FBTYPE,
		Value:     types.NewTextValue(string(value)),
	}
}

type FreeBusyTimeType string

const (
	FREE            FreeBusyTimeType = "FREE"
	BUSSY           FreeBusyTimeType = "BUSSY"
	BUSYUNAVAILABLE FreeBusyTimeType = "BUSY-UNAVAILABLE"
	BUSYTENTATIVE   FreeBusyTimeType = "BUSY-TENTATIVE"
)
