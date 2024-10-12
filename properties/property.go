package properties

import (
	"bytes"
	"fmt"
	"io"

	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/types/recurrence_rule"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

const (
	escapeCharacter  = "\r\n"
	foldingIndex     = 75
	foldingCharacter = escapeCharacter + " "
)

type Property interface {
	// ToICalendarPropFormat format output to the icalendar specs
	ToICalendarPropFormat(output io.Writer)
	// GetName return the property name
	GetName() registries.PropertyRegistry
	// GetValue return the string of a property
	GetValue() string
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
	// GetActionValue return the types.ActionValue associated to this ActionProperty
	GetActionValue() registries.ActionRegistry
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

type BlockDelimiterPropertyType interface {
	Property
}

type NonStandardPropertyType interface {
	Property
}

type textPropertyType struct {
	PropName   registries.PropertyRegistry
	Value      types.TextValue
	Values     []types.TextValue
	Parameters parameters.Parameters
}

func (tP *textPropertyType) GetName() registries.PropertyRegistry {
	return tP.PropName
}

func (tP *textPropertyType) GetValue() string {
	return tP.Value.S
}

type integerPropertyType struct {
	PropName   registries.PropertyRegistry
	Value      types.IntegerValue
	Parameters parameters.Parameters
}

func (iP *integerPropertyType) GetName() registries.PropertyRegistry {
	return iP.PropName
}

func (iP *integerPropertyType) GetValue() string {
	return iP.Value.GetStringValue()
}

type dateTimePropertyType struct {
	PropName   registries.PropertyRegistry
	Value      types.DateTimeValue
	Values     []types.DateTimeValue
	Parameters parameters.Parameters
}

func (dtP *dateTimePropertyType) GetName() registries.PropertyRegistry {
	return dtP.PropName
}

func (dtP *dateTimePropertyType) GetValue() string {
	return dtP.Value.GetValue()
}

type datePropertyType struct {
	PropName   registries.PropertyRegistry
	Value      types.DateValue
	Values     []types.DateValue
	Parameters parameters.Parameters
}

func (dP *datePropertyType) GetName() registries.PropertyRegistry {
	return dP.PropName
}

func (dP *datePropertyType) GetValue() string {
	return dP.Value.GetValue()
}

type periodPropertyType struct {
	PropName   registries.PropertyRegistry
	Value      types.PeriodValue
	Values     []types.PeriodValue
	Parameters parameters.Parameters
}

func (pP *periodPropertyType) GetName() registries.PropertyRegistry {
	return pP.PropName
}

func (pP *periodPropertyType) GetValue() string {
	return pP.Value.GetValue()
}

type durationPropertyType struct {
	PropName   registries.PropertyRegistry
	Value      types.DurationValue
	Parameters parameters.Parameters
}

func (dP *durationPropertyType) GetName() registries.PropertyRegistry {
	return dP.PropName
}

func (dP *durationPropertyType) GetValue() string {
	return dP.Value.GetValue()
}

type geoPropertyType struct {
	PropName   registries.PropertyRegistry
	Longitude  types.FloatValue
	Latitude   types.FloatValue
	Parameters parameters.Parameters
}

func (gP *geoPropertyType) GetName() registries.PropertyRegistry {
	return gP.PropName
}

func (gP *geoPropertyType) GetValue() string {
	return "GEO"
}

type calendarUserAddressPropertyType struct {
	PropName   registries.PropertyRegistry
	Value      types.CalendarUserAddressValue
	Parameters parameters.Parameters
}

func (caP *calendarUserAddressPropertyType) GetName() registries.PropertyRegistry {
	return caP.PropName
}

func (caP *calendarUserAddressPropertyType) GetValue() string {
	return caP.Value.GetValue()
}

type utcOffsetPropertyType struct {
	PropName   registries.PropertyRegistry
	Value      types.UtcOffsetValue
	Parameters parameters.Parameters
}

func (uoP *utcOffsetPropertyType) GetName() registries.PropertyRegistry {
	return uoP.PropName
}

func (uoP *utcOffsetPropertyType) GetValue() string {
	return uoP.Value.GetValue()
}

type uriPropertyType struct {
	PropName   registries.PropertyRegistry
	Value      types.UriValue
	Parameters parameters.Parameters
}

func (uP *uriPropertyType) GetName() registries.PropertyRegistry {
	return uP.PropName
}

func (uP *uriPropertyType) GetValue() string {
	return uP.Value.GetValue()
}

type requestStatusPropertyType struct {
	PropName          registries.PropertyRegistry
	StatusCode        types.TextValue
	StatusDescription types.TextValue
	ExtraData         types.TextValue
	Parameters        parameters.Parameters
}

func (rsP *requestStatusPropertyType) GetName() registries.PropertyRegistry {
	return rsP.PropName
}

func (rsP *requestStatusPropertyType) GetValue() string {
	return "RRST"
}

type recurrenceRulePropertyType struct {
	PropName registries.PropertyRegistry
	Value    recurrence_rule.RecurrenceRuleValue
}

func (rrP *recurrenceRulePropertyType) GetName() registries.PropertyRegistry {
	return rrP.PropName
}

func (rrP *recurrenceRulePropertyType) GetValue() string {
	return rrP.Value.GetValue()
}

type actionPropertyType struct {
	PropName   registries.PropertyRegistry
	Value      types.ActionValue
	Parameters parameters.Parameters
}

func (aP *actionPropertyType) GetName() registries.PropertyRegistry {
	return aP.PropName
}

func (aP *actionPropertyType) GetValue() string {
	return string(aP.Value.Value)
}

type classificationPropertyType struct {
	PropName   registries.PropertyRegistry
	Value      types.ClassificationValue
	Parameters parameters.Parameters
}

func (cP *classificationPropertyType) GetName() registries.PropertyRegistry {
	return cP.PropName
}

func (cP *classificationPropertyType) GetValue() string {
	return string(cP.Value.Value)
}

type statusPropertyType struct {
	PropName   registries.PropertyRegistry
	Value      types.StatusValue
	Parameters parameters.Parameters
}

func (sP *statusPropertyType) GetName() registries.PropertyRegistry {
	return sP.PropName
}

func (sP *statusPropertyType) GetValue() string {
	return string(sP.Value.Value)
}

type timeTransparencyPropertyType struct {
	PropName   registries.PropertyRegistry
	Value      types.TimeTransparencyValue
	Parameters parameters.Parameters
}

func (ttP *timeTransparencyPropertyType) GetName() registries.PropertyRegistry {
	return ttP.PropName
}

func (ttP *timeTransparencyPropertyType) GetValue() string {
	return string(ttP.Value.Value)
}

type blockDelimiterPropertyType struct {
	PropName   registries.PropertyRegistry
	Value      types.BlockDelimiterValue
	Parameters parameters.Parameters
}

func (bdP *blockDelimiterPropertyType) GetName() registries.PropertyRegistry {
	return bdP.PropName
}

func (bdP *blockDelimiterPropertyType) GetValue() string {
	return string(bdP.Value.Value)
}

type nonStandardPropertyType struct {
	PropName   string
	Value      types.TextValue
	Parameters parameters.Parameters
}

func (nSP *nonStandardPropertyType) GetName() registries.PropertyRegistry {
	return registries.NonStandardProp
}

func (nSP *nonStandardPropertyType) GetValue() string {
	return nSP.Value.S
}

func computeParameters(paramsOutput io.Writer, params parameters.Parameters) {
	for i := 0; i < len(params); i++ {
		paramsOutput.Write([]byte(fmt.Sprint(";")))
		params[i].ToICalendarParamFormat(paramsOutput)
	}
}

func foldOutput(unfoldedOutput *bytes.Buffer) {
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
	var paramsOutput bytes.Buffer
	if gP.Parameters != nil {
		computeParameters(&paramsOutput, gP.Parameters)
	}
	unfoldedOutput.Write(
		[]byte(
			fmt.Sprintf(
				"%s%s:%f;%f",
				gP.PropName,
				paramsOutput.String(),
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

func (rsP *requestStatusPropertyType) ToICalendarPropFormat(output io.Writer) {
	var unfoldedOutput bytes.Buffer
	var paramsOutput bytes.Buffer
	var extraValueOutput bytes.Buffer
	if rsP.Parameters != nil {
		computeParameters(&paramsOutput, rsP.Parameters)
	}
	if rsP.ExtraData.S != "" {
		extraValueOutput.Write([]byte(fmt.Sprintf(";%s", rsP.ExtraData.S)))
	}
	unfoldedOutput.Write(
		[]byte(
			fmt.Sprintf(
				"%s%s:%s;%s",
				rsP.PropName,
				paramsOutput.String(),
				rsP.StatusCode.S,
				rsP.StatusDescription.S,
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
	var paramsOutput bytes.Buffer
	if aP.Parameters != nil {
		computeParameters(&paramsOutput, aP.Parameters)
	}
	unfoldedOutput.Write(
		[]byte(fmt.Sprintf("%s%s:%s", aP.PropName, paramsOutput.String(), aP.Value.Value)),
	)
	foldOutput(&unfoldedOutput)
	unfoldedOutput.WriteTo(output)
}

func (cP *classificationPropertyType) ToICalendarPropFormat(output io.Writer) {
	var unfoldedOutput bytes.Buffer
	var paramsOutput bytes.Buffer
	if cP.Parameters != nil {
		computeParameters(&paramsOutput, cP.Parameters)
	}
	unfoldedOutput.Write(
		[]byte(fmt.Sprintf("%s%s:%s", cP.PropName, paramsOutput.String(), cP.Value.Value)),
	)
	foldOutput(&unfoldedOutput)
	unfoldedOutput.WriteTo(output)
}

func (sP *statusPropertyType) ToICalendarPropFormat(output io.Writer) {
	var unfoldedOutput bytes.Buffer
	var paramsOutput bytes.Buffer
	if sP.Parameters != nil {
		computeParameters(&paramsOutput, sP.Parameters)
	}
	unfoldedOutput.Write(
		[]byte(fmt.Sprintf("%s%s:%s", sP.PropName, paramsOutput.String(), sP.Value.Value)),
	)
	foldOutput(&unfoldedOutput)
	unfoldedOutput.WriteTo(output)
}

func (ttP *timeTransparencyPropertyType) ToICalendarPropFormat(output io.Writer) {
	var unfoldedOutput bytes.Buffer
	var paramsOutput bytes.Buffer
	if ttP.Parameters != nil {
		computeParameters(&paramsOutput, ttP.Parameters)
	}
	unfoldedOutput.Write(
		[]byte(fmt.Sprintf("%s%s:%s", ttP.PropName, paramsOutput.String(), ttP.Value.Value)),
	)
	foldOutput(&unfoldedOutput)
	unfoldedOutput.WriteTo(output)
}

func (bT *blockDelimiterPropertyType) ToICalendarPropFormat(output io.Writer) {
	var unfoldedOutput bytes.Buffer
	var paramsOutput bytes.Buffer
	if bT.Parameters != nil {
		computeParameters(&paramsOutput, bT.Parameters)
	}
	unfoldedOutput.Write(
		[]byte(fmt.Sprintf("%s%s:%s", bT.PropName, paramsOutput.String(), bT.Value.Value)),
	)
	foldOutput(&unfoldedOutput)
	unfoldedOutput.WriteTo(output)
}

func (nSP *nonStandardPropertyType) ToICalendarPropFormat(output io.Writer) {
	var unfoldedOutput bytes.Buffer
	var paramsOutput bytes.Buffer
	if nSP.Parameters != nil {
		computeParameters(&paramsOutput, nSP.Parameters)
	}
	unfoldedOutput.Write(
		[]byte(fmt.Sprintf("%s%s:%s", nSP.PropName, paramsOutput.String(), nSP.Value.S)),
	)
	foldOutput(&unfoldedOutput)
	unfoldedOutput.WriteTo(output)
}
