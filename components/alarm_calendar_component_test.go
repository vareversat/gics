package components

import (
	"bytes"
	"testing"

	"github.com/vareversat/gics/registries"

	"github.com/stretchr/testify/assert"
	"github.com/vareversat/gics/properties"
)

func TestNewAlarmCalendarComponent(t *testing.T) {
	action := properties.NewActionProperty(registries.Audio)
	trigger, _ := properties.NewTriggerFromDurationProperty(3600 * 1e9)
	component := NewAlarmCalendarComponent(action, trigger)

	assert.NotNil(t, component)
	assert.Equal(t, action.GetActionValue().GetValue(), registries.Audio)
}

func TestAlarmCalendarComponent_SerializeToICSFormat(t *testing.T) {
	action := properties.NewActionProperty(registries.Audio)
	trigger, _ := properties.NewTriggerFromDurationProperty(-3600 * 1e9)
	description := properties.NewDescriptionProperty("Test Alarm")
	summary := properties.NewSummaryProperty("Alarm SummaryProp")

	component := NewAlarmCalendarComponent(action, trigger, description, summary)

	var buf bytes.Buffer
	component.SerializeToICSFormat(&buf)
	serialized := buf.String()

	expected := "BEGIN:VALARM\r\nACTION:AUDIO\r\nTRIGGER:-PT1H\r\nDESCRIPTION:Test Alarm\r\nSUMMARY:Alarm SummaryProp\r\nEND:VALARM\r\n"
	assert.Equal(t, expected, serialized)
}
