package pkg

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.4

import "github.com/vareversat/gics/pkg/properties"

type Calendar struct {
	Begin     properties.ComponentBeginProperty // Mandatory
	Version   properties.VersionProperty        // Mandatory
	ProdId    properties.ProIdProperty          // Mandatory
	Method    properties.MethodProperty         // Optional
	CalScale  properties.CalScaleProperty       // Optional
	Calendars Calendars
	End       properties.ComponentEndProperty // Mandatory
}
