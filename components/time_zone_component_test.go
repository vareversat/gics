package components

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vareversat/gics/properties"
)

func TestNewTimeZoneCalendarComponent_STANDARD(t *testing.T) {
	subComponent := NewTimeZoneCalendarStandardSubcomponent(
		properties.NewDateTimeStampProperty(time.Now()),
		properties.NewTimeZoneOffsetToProperty("-0100"),
		properties.NewTimeZoneOffsetFromProperty("-0300"))

	component := NewTimeZoneCalendarComponent(
		[]TimeZoneCalendarSubComponent{subComponent},
		properties.NewTimeZoneIdProperty("Europe/Paris"),
	)

	assert.NotNil(t, component)
}

func TestNewTimeZoneCalendarComponent_DAYLIGHT(t *testing.T) {
	subComponent := NewTimeZoneDayLightSubcomponent(
		properties.NewDateTimeStampProperty(time.Now()),
		properties.NewTimeZoneOffsetToProperty("-0100"),
		properties.NewTimeZoneOffsetFromProperty("-0300"))

	component := NewTimeZoneCalendarComponent(
		[]TimeZoneCalendarSubComponent{subComponent},
		properties.NewTimeZoneIdProperty("Europe/Paris"),
	)

	assert.NotNil(t, component)
}

func TestTimeZoneCalendarComponent_STANDARD_SerializeToICSFormat(t *testing.T) {
	myTime := time.Now()
	myTimeToString := myTime.Format("20060102T150405Z")

	subComponent := NewTimeZoneCalendarStandardSubcomponent(
		properties.NewDateTimeStampProperty(myTime),
		properties.NewTimeZoneOffsetToProperty("-0100"),
		properties.NewTimeZoneOffsetFromProperty("-0300"))

	component := NewTimeZoneCalendarComponent(
		[]TimeZoneCalendarSubComponent{subComponent},
		properties.NewTimeZoneIdProperty("Europe/Paris"),
	)

	var buf bytes.Buffer
	component.SerializeToICSFormat(&buf)
	serialized := buf.String()

	expected := fmt.Sprintf(
		"BEGIN:VTIMEZONE\r\nTZID:Europe/Paris\r\nBEGIN:STANDARD\r\nDTSTAMP:%s\r\nTZOFFSETTO:-0100\r\nTZOFFSETFROM:-0300\r\nEND:STANDARD\r\nEND:VTIMEZONE\r\n",
		myTimeToString,
	)
	assert.Equal(t, expected, serialized)
}

func TestTimeZoneCalendarComponent_DAYLIGHT_SerializeToICSFormat(t *testing.T) {
	myTime := time.Now()
	myTimeToString := myTime.Format("20060102T150405Z")

	subComponent := NewTimeZoneDayLightSubcomponent(
		properties.NewDateTimeStampProperty(myTime),
		properties.NewTimeZoneOffsetToProperty("-0100"),
		properties.NewTimeZoneOffsetFromProperty("-0300"))

	component := NewTimeZoneCalendarComponent(
		[]TimeZoneCalendarSubComponent{subComponent},
		properties.NewTimeZoneIdProperty("Europe/Paris"),
	)

	var buf bytes.Buffer
	component.SerializeToICSFormat(&buf)
	serialized := buf.String()

	expected := fmt.Sprintf(
		"BEGIN:VTIMEZONE\r\nTZID:Europe/Paris\r\nBEGIN:DAYLIGHT\r\nDTSTAMP:%s\r\nTZOFFSETTO:-0100\r\nTZOFFSETFROM:-0300\r\nEND:DAYLIGHT\r\nEND:VTIMEZONE\r\n",
		myTimeToString,
	)
	assert.Equal(t, expected, serialized)
}
