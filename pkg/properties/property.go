package properties

import (
	"bytes"
	"fmt"
	"io"

	"github.com/vareversat/gics/pkg/parameters"
	"github.com/vareversat/gics/pkg/types/recurrence_rule"

	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type Property interface {
	ToICalendarPropFormat(output io.Writer)
}

type Properties []Property

type TextPropertyType interface {
	Property
}

type IntegerPropertyType interface {
	Property
}

type DateTimePropertyType interface {
	Property
}

type DatePropertyType interface {
	Property
}

type DurationPropertyType interface {
	Property
}

type PeriodPropertyType interface {
	Property
}

type GeoPropertyType interface {
	Property
}

type CalendarUserAddressPropertyType interface {
	Property
}

type UtcOffsetPropertyType interface {
	Property
}

type UriPropertyType interface {
	Property
}

type RequestStatusPropertyType interface {
	Property
}

type RecurrenceRulePropertyType interface {
	Property
}

type ActionPropertyType interface {
	Property
}

type ClassificationPropertyType interface {
	Property
}

type StatusPropertyType interface {
	Property
}

type TimeTransparencyPropertyType interface {
	Property
}

type textPropertyType struct {
	PropName   registries.Properties
	Value      types.TextValue
	Values     []types.TextValue
	Parameters parameters.Parameters
}

type integerPropertyType struct {
	PropName   registries.Properties
	Value      types.IntegerValue
	Parameters parameters.Parameters
}

type dateTimePropertyType struct {
	PropName   registries.Properties
	Value      types.DateTimeValue
	Values     []types.DateTimeValue
	Parameters parameters.Parameters
}

type datePropertyType struct {
	PropName   registries.Properties
	Value      types.DateValue
	Values     []types.DateValue
	Parameters parameters.Parameters
}

type periodPropertyType struct {
	PropName   registries.Properties
	Value      types.PeriodValue
	Values     []types.PeriodValue
	Parameters parameters.Parameters
}

type durationPropertyType struct {
	PropName   registries.Properties
	Value      types.DurationValue
	Parameters parameters.Parameters
}

type geoPropertyType struct {
	PropName   registries.Properties
	Longitude  types.FloatValue
	Latitude   types.FloatValue
	Parameters parameters.Parameters
}

type calendarUserAddressPropertyType struct {
	PropName   registries.Properties
	Value      types.CalendarUserAddressValue
	Parameters parameters.Parameters
}

type utcOffsetPropertyType struct {
	PropName   registries.Properties
	Value      types.UtcOffsetValue
	Parameters parameters.Parameters
}

type uriPropertyType struct {
	PropName   registries.Properties
	Value      types.UriValue
	Parameters parameters.Parameters
}
type requestStatusPropertyType struct {
	PropName          registries.Properties
	StatusCode        types.TextValue
	StatusDescription types.TextValue
	ExtraData         types.TextValue
	Parameters        parameters.Parameters
}

type recurrenceRulePropertyType struct {
	PropName registries.Properties
	Value    recurrence_rule.RecurrenceRuleValue
}

type actionPropertyType struct {
	PropName registries.Properties
	Value    types.ActionValue
}

type classificationPropertyType struct {
	PropName registries.Properties
	Value    types.ClassificationValue
}

type statusPropertyType struct {
	PropName registries.Properties
	Value    types.StatusValue
}

type timeTransparencyPropertyType struct {
	PropName registries.Properties
	Value    types.TimeTransparencyValue
}

func computeParameters(paramsOutput io.Writer, params parameters.Parameters) {
	for i := 0; i < len(params); i++ {
		paramsOutput.Write([]byte(fmt.Sprint(";")))
		params[i].ToICalendarParamFormat(paramsOutput)
	}
}

func foldOutput(unfoldedOutput *bytes.Buffer) {
	escapeCharacter := "\r\n"
	foldingIndex := 75
	foldingCharacter := escapeCharacter + " "
	for i := foldingIndex; i < unfoldedOutput.Len(); i += foldingIndex + 1 {
		var afterFoldingBlock = bytes.Clone(unfoldedOutput.Bytes()[i:])
		var beforeFoldingBlock = bytes.Clone(unfoldedOutput.Bytes()[:i])
		unfoldedOutput.Reset()
		unfoldedOutput.Write(beforeFoldingBlock)
		unfoldedOutput.WriteString(foldingCharacter)
		unfoldedOutput.Write(afterFoldingBlock)
	}
	unfoldedOutput.WriteString(escapeCharacter)
}

func (tP *textPropertyType) ToICalendarPropFormat(output io.Writer) {
	var paramsOutput bytes.Buffer
	var valuesOutput bytes.Buffer
	var unfoldedOutput bytes.Buffer
	if tP.Values != nil {
		for i := 0; i < len(tP.Values); i++ {
			valuesOutput.WriteString(tP.Values[i].S)
			if len(tP.Values)-1 > i {
				valuesOutput.WriteString(fmt.Sprint(","))
			}
		}
	} else {
		valuesOutput.WriteString(tP.Value.S)
	}
	if tP.Parameters != nil {
		computeParameters(&paramsOutput, tP.Parameters)
	}
	unfoldedOutput.Write(
		[]byte(fmt.Sprintf("%s%s:%s", tP.PropName, paramsOutput.String(), valuesOutput.String())),
	)
	foldOutput(&unfoldedOutput)
	unfoldedOutput.WriteTo(output)
}

func (iP *integerPropertyType) ToICalendarPropFormat(output io.Writer) {
	var unfoldedOutput bytes.Buffer
	var paramsOutput bytes.Buffer
	if iP.Parameters != nil {
		computeParameters(&paramsOutput, iP.Parameters)
	}
	unfoldedOutput.Write(
		[]byte(
			fmt.Sprintf(
				"%s%s:%s",
				iP.PropName,
				paramsOutput.String(),
				iP.Value.GetStringValue(),
			),
		),
	)
	foldOutput(&unfoldedOutput)
	unfoldedOutput.WriteTo(output)
}

func (dP *dateTimePropertyType) ToICalendarPropFormat(output io.Writer) {
	var unfoldedOutput bytes.Buffer
	var valuesOutput bytes.Buffer
	var paramsOutput bytes.Buffer
	if dP.Values != nil {
		for i := 0; i < len(dP.Values); i++ {
			valuesOutput.WriteString(dP.Values[i].GetValue())
			if len(dP.Values)-1 > i {
				valuesOutput.WriteString(fmt.Sprint(","))
			}
		}
	} else {
		valuesOutput.WriteString(dP.Value.GetValue())
	}
	if dP.Parameters != nil {
		computeParameters(&paramsOutput, dP.Parameters)
	}
	unfoldedOutput.Write(
		[]byte(
			fmt.Sprintf("%s%s:%s", dP.PropName, paramsOutput.String(), valuesOutput.String()),
		),
	)
	foldOutput(&unfoldedOutput)
	unfoldedOutput.WriteTo(output)
}

func (dP *datePropertyType) ToICalendarPropFormat(output io.Writer) {
	var unfoldedOutput bytes.Buffer
	var valuesOutput bytes.Buffer
	var paramsOutput bytes.Buffer
	if dP.Values != nil {
		for i := 0; i < len(dP.Values); i++ {
			valuesOutput.WriteString(dP.Values[i].GetValue())
			if len(dP.Values)-1 > i {
				valuesOutput.WriteString(fmt.Sprint(","))
			}
		}
	} else {
		valuesOutput.WriteString(dP.Value.GetValue())
	}
	if dP.Parameters != nil {
		computeParameters(&paramsOutput, dP.Parameters)
	}
	unfoldedOutput.Write(
		[]byte(
			(fmt.Sprintf("%s%s:%s", dP.PropName, paramsOutput.String(), valuesOutput.String())),
		),
	)
	foldOutput(&unfoldedOutput)
	unfoldedOutput.WriteTo(output)
}

func (pP *periodPropertyType) ToICalendarPropFormat(output io.Writer) {
	var unfoldedOutput bytes.Buffer
	var valuesOutput bytes.Buffer
	if pP.Values != nil {
		for i := 0; i < len(pP.Values); i++ {
			valuesOutput.WriteString(pP.Values[i].GetValue())
			if len(pP.Values)-1 > i {
				valuesOutput.WriteString(fmt.Sprint(","))
			}
		}
	} else {
		valuesOutput.WriteString(pP.Value.GetValue())
	}
	unfoldedOutput.Write([]byte(fmt.Sprintf("%s:%s", pP.PropName, valuesOutput.String())))
	foldOutput(&unfoldedOutput)
	unfoldedOutput.WriteTo(output)
}

func (dP *durationPropertyType) ToICalendarPropFormat(output io.Writer) {
	var unfoldedOutput bytes.Buffer
	var paramsOutput bytes.Buffer
	if dP.Parameters != nil {
		computeParameters(&paramsOutput, dP.Parameters)
	}
	unfoldedOutput.Write(
		[]byte(fmt.Sprintf("%s%s:%s", dP.PropName, paramsOutput.String(), dP.Value.GetValue())),
	)
	foldOutput(&unfoldedOutput)
	unfoldedOutput.WriteTo(output)
}

func (gP *geoPropertyType) ToICalendarPropFormat(output io.Writer) {
	var unfoldedOutput bytes.Buffer
	unfoldedOutput.Write(
		[]byte(
			fmt.Sprintf(
				"%s:%f;%f",
				gP.PropName,
				gP.Longitude.GetValue(),
				gP.Latitude.GetValue(),
			),
		),
	)
	foldOutput(&unfoldedOutput)
	unfoldedOutput.WriteTo(output)
}

func (gP *calendarUserAddressPropertyType) ToICalendarPropFormat(output io.Writer) {
	var paramsOutput bytes.Buffer
	var unfoldedOutput bytes.Buffer
	if gP.Parameters != nil {
		computeParameters(&paramsOutput, gP.Parameters)
	}
	unfoldedOutput.Write(
		[]byte(fmt.Sprintf("%s%s:%s", gP.PropName, paramsOutput.String(), gP.Value.GetValue())),
	)
	foldOutput(&unfoldedOutput)
	unfoldedOutput.WriteTo(output)
}

func (uP *utcOffsetPropertyType) ToICalendarPropFormat(output io.Writer) {
	var unfoldedOutput bytes.Buffer
	var paramsOutput bytes.Buffer
	if uP.Parameters != nil {
		computeParameters(&paramsOutput, uP.Parameters)
	}
	unfoldedOutput.Write(
		[]byte(fmt.Sprintf("%s%s:%s", uP.PropName, paramsOutput.String(), uP.Value.GetValue())),
	)
	foldOutput(&unfoldedOutput)
	unfoldedOutput.WriteTo(output)
}

func (uP *uriPropertyType) ToICalendarPropFormat(output io.Writer) {
	var unfoldedOutput bytes.Buffer
	var paramsOutput bytes.Buffer
	if uP.Parameters != nil {
		computeParameters(&paramsOutput, uP.Parameters)
	}
	unfoldedOutput.Write(
		[]byte(fmt.Sprintf("%s%s:%s", uP.PropName, paramsOutput.String(), uP.Value.GetValue())),
	)
	foldOutput(&unfoldedOutput)
	unfoldedOutput.WriteTo(output)
}

func (dP *requestStatusPropertyType) ToICalendarPropFormat(output io.Writer) {
	var unfoldedOutput bytes.Buffer
	var paramsOutput bytes.Buffer
	var extraValueOutput bytes.Buffer
	if dP.Parameters != nil {
		computeParameters(&paramsOutput, dP.Parameters)
	}
	if dP.ExtraData.S != "" {
		extraValueOutput.Write([]byte(fmt.Sprintf(";%s", dP.ExtraData.S)))
	}
	unfoldedOutput.Write(
		[]byte(
			fmt.Sprintf(
				"%s%s:%s;%s",
				dP.PropName,
				paramsOutput.String(),
				dP.StatusCode.S,
				dP.StatusDescription.S,
			),
		),
	)
	foldOutput(&unfoldedOutput)
	unfoldedOutput.WriteTo(output)
}

func (rrP *recurrenceRulePropertyType) ToICalendarPropFormat(output io.Writer) {
	var unfoldedOutput bytes.Buffer
	var partsOutput bytes.Buffer
	if rrP.Value.Parts != nil {
		partsOutput.Write([]byte(fmt.Sprint(rrP.Value.GetValue())))
	}
	unfoldedOutput.Write([]byte(fmt.Sprintf("%s:%s", rrP.PropName, rrP.Value.GetValue())))
	foldOutput(&unfoldedOutput)
	unfoldedOutput.WriteTo(output)
}

func (aP *actionPropertyType) ToICalendarPropFormat(output io.Writer) {
	var unfoldedOutput bytes.Buffer
	unfoldedOutput.Write([]byte(fmt.Sprintf("%s:%s", aP.PropName, aP.Value.Value)))
	foldOutput(&unfoldedOutput)
	unfoldedOutput.WriteTo(output)
}

func (cP *classificationPropertyType) ToICalendarPropFormat(output io.Writer) {
	var unfoldedOutput bytes.Buffer
	unfoldedOutput.Write([]byte(fmt.Sprintf("%s:%s", cP.PropName, cP.Value.Value)))
	foldOutput(&unfoldedOutput)
	unfoldedOutput.WriteTo(output)
}

func (sP *statusPropertyType) ToICalendarPropFormat(output io.Writer) {
	var unfoldedOutput bytes.Buffer
	unfoldedOutput.Write([]byte(fmt.Sprintf("%s:%s", sP.PropName, sP.Value.Value)))
	foldOutput(&unfoldedOutput)
	unfoldedOutput.WriteTo(output)
}

func (tP *timeTransparencyPropertyType) ToICalendarPropFormat(output io.Writer) {
	var unfoldedOutput bytes.Buffer
	unfoldedOutput.Write([]byte(fmt.Sprintf("%s:%s", tP.PropName, tP.Value.Value)))
	foldOutput(&unfoldedOutput)
	unfoldedOutput.WriteTo(output)
}
