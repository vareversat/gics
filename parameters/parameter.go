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
	GetParamName() registries.Parameters
	// GetParamValue return the parameter name in (string typed)
	GetParamValue() string
}

// Parameters is an array of Parameter
type Parameters []Parameter

type TextParameter interface {
	Parameter
}

type UriParameter interface {
	Parameter
}

type BooleanParameter interface {
	Parameter
}

type CalendarUserAddressParameter interface {
	Parameter
}

type textParameter struct {
	ParamName registries.Parameters
	Value     types.TextValue
}

type uriParameter struct {
	ParamName registries.Parameters
	Value     types.UriValue
}

type booleanParameter struct {
	ParamName registries.Parameters
	Value     types.BooleanValue
}

type calendarUserAddressParameter struct {
	ParamName registries.Parameters
	Value     types.CalendarUserAddressValue
}

func (tP *textParameter) ToICalendarParamFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", tP.ParamName, tP.GetParamValue())))
}

func (tP *textParameter) GetParamName() registries.Parameters {
	return tP.ParamName
}

func (tP *textParameter) GetParamValue() string {
	return tP.Value.S
}

func (uP *uriParameter) ToICalendarParamFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", uP.ParamName, uP.GetParamValue())))
}

func (uP *uriParameter) GetParamName() registries.Parameters {
	return uP.ParamName
}

func (uP *uriParameter) GetParamValue() string {
	return fmt.Sprintf("\"%s\"", uP.Value.GetValue())
}

func (bP *booleanParameter) ToICalendarParamFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", bP.ParamName, bP.GetParamValue())))
}

func (bP *booleanParameter) GetParamName() registries.Parameters {
	return bP.ParamName
}

func (bP *booleanParameter) GetParamValue() string {
	return bP.Value.GetValue()
}

func (cP *calendarUserAddressParameter) ToICalendarParamFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", cP.ParamName, cP.GetParamValue())))
}

func (cP *calendarUserAddressParameter) GetParamName() registries.Parameters {
	return cP.ParamName
}

func (cP *calendarUserAddressParameter) GetParamValue() string {
	return fmt.Sprintf("\"%s\"", cP.Value.GetValue())
}
