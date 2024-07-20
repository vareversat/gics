package parameters

import (
	"fmt"
	"io"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

// Parameter is the interface definition of a property parameter. See [RFC-5545] for more info
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2
type Parameter interface {
	// ToICalendarParamFormat format output to the icalendar specs
	ToICalendarParamFormat(output io.Writer)
	// GetParamName return the parameter name
	GetParamName() registries.ParameterRegistry
	// GetParamValue return the parameter name in (string typed)
	GetParamValue() string
}

// Parameters is an array of Parameter
type Parameters []Parameter

// TextTypeParameter is the interface representation of a text typed parameter
// Go representation: string
type TextTypeParameter interface {
	Parameter
}

type textParameter struct {
	ParamName registries.ParameterRegistry
	Value     types.TextValue
}

func (tP *textParameter) ToICalendarParamFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", tP.ParamName, tP.GetParamValue())))
}

func (tP *textParameter) GetParamName() registries.ParameterRegistry {
	return tP.ParamName
}

func (tP *textParameter) GetParamValue() string {
	return tP.Value.S
}

// UriTypeParameter is the interface representation of an uri typed parameter
// Go representation: url.URL
type UriTypeParameter interface {
	Parameter
}

type uriParameter struct {
	ParamName registries.ParameterRegistry
	Value     types.UriValue
}

func (uP *uriParameter) ToICalendarParamFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", uP.ParamName, uP.GetParamValue())))
}

func (uP *uriParameter) GetParamName() registries.ParameterRegistry {
	return uP.ParamName
}

func (uP *uriParameter) GetParamValue() string {
	return fmt.Sprintf("\"%s\"", uP.Value.GetValue())
}

// BooleanTypeParameter is the interface representation of a boolean typed parameter
// Go representation: bool
type BooleanTypeParameter interface {
	Parameter
}

type booleanParameter struct {
	ParamName registries.ParameterRegistry
	Value     types.BooleanValue
}

func (bP *booleanParameter) ToICalendarParamFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", bP.ParamName, bP.GetParamValue())))
}

func (bP *booleanParameter) GetParamName() registries.ParameterRegistry {
	return bP.ParamName
}

func (bP *booleanParameter) GetParamValue() string {
	return bP.Value.GetValue()
}

// CalendarUserAddressTypeParameter is the interface representation of a calendar user address typed parameter
// Go representation: url.URL
type CalendarUserAddressTypeParameter interface {
	Parameter
}

type calendarUserAddressParameter struct {
	ParamName registries.ParameterRegistry
	Value     types.CalendarUserAddressValue
}

func (cP *calendarUserAddressParameter) ToICalendarParamFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", cP.ParamName, cP.GetParamValue())))
}

func (cP *calendarUserAddressParameter) GetParamName() registries.ParameterRegistry {
	return cP.ParamName
}

func (cP *calendarUserAddressParameter) GetParamValue() string {
	return fmt.Sprintf("\"%s\"", cP.Value.GetValue())
}
