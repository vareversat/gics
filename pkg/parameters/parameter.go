package parameters

import (
	"fmt"
	"io"

	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type Parameter interface {
	ToICalendarParamFormat(output io.Writer)
	GetParamName() registries.Parameters
	GetParamValue() string
}

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
	output.Write([]byte(fmt.Sprintf("%s=%s", tP.ParamName, tP.Value.S)))
}

func (tP *textParameter) GetParamName() registries.Parameters {
	return tP.ParamName
}

func (tP *textParameter) GetParamValue() string {
	return tP.Value.S
}

func (uP *uriParameter) ToICalendarParamFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", uP.ParamName, uP.Value.GetValue())))
}

func (tP *uriParameter) GetParamName() registries.Parameters {
	return tP.ParamName
}

func (tP *uriParameter) GetParamValue() string {
	return tP.Value.GetValue()
}

func (bP *booleanParameter) ToICalendarParamFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", bP.ParamName, bP.Value.GetValue())))
}

func (bP *booleanParameter) GetParamName() registries.Parameters {
	return bP.ParamName
}

func (bP *booleanParameter) GetParamValue() string {
	return bP.Value.GetValue()
}

func (cP *calendarUserAddressParameter) ToICalendarParamFormat(output io.Writer) {
	output.Write([]byte(fmt.Sprintf("%s=%s", cP.ParamName, cP.Value.GetValue())))
}

func (cP *calendarUserAddressParameter) GetParamName() registries.Parameters {
	return cP.ParamName
}

func (cP *calendarUserAddressParameter) GetParamValue() string {
	return cP.Value.GetValue()
}
