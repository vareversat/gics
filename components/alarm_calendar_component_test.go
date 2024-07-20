package components

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/vareversat/gics/registries"

	"github.com/stretchr/testify/assert"
	"github.com/vareversat/gics/properties"
	"github.com/vareversat/gics/types"
)

func TestNewAlarmCalendarComponent(t *testing.T) {
	action := properties.NewActionProperty(registries.Audio)
	trigger := properties.NewTriggerProperty(time.Now(), "PT5M", types.WithUtcTime)
	component := NewAlarmCalendarComponent(action, trigger)

	assert.NotNil(t, component)
	assert.Equal(t, action.GetActionValue(), registries.Audio)
}

func TestAlarmCalendarComponent_SerializeToICSFormat(t *testing.T) {
	myTime := time.Now()
	myTimeToString := myTime.Format("20060102T150405Z")

	action := properties.NewActionProperty(registries.Audio)
	trigger := properties.NewTriggerProperty(myTime, "PT5M", types.WithUtcTime)
	description := properties.NewDescriptionProperty("Test Alarm")
	summary := properties.NewSummaryProperty("Alarm Summary")

	component := NewAlarmCalendarComponent(action, trigger, description, summary)

	var buf bytes.Buffer
	component.SerializeToICSFormat(&buf)
	serialized := buf.String()

	expected := fmt.Sprintf(
		"BEGIN:VALARM\r\nACTION:AUDIO\r\nTRIGGER:%s\r\nDESCRIPTION:Test Alarm\r\nSUMMARY:Alarm Summary\r\nEND:VALARM\r\n",
		myTimeToString,
	)
	assert.Equal(t, expected, serialized)
}
