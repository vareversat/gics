package components

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vareversat/gics/pkg/properties"
	"github.com/vareversat/gics/pkg/types"
	"testing"
	"time"
)

func TestNewAlarmCalendarComponent(t *testing.T) {
	action := properties.NewActionProperty(types.AUDIO)
	trigger := properties.NewTriggerProperty(time.Now(), "PT5M", types.WithUtcTime)
	component := NewAlarmCalendarComponent(action, trigger)

	assert.NotNil(t, component)
	assert.Equal(t, action.GetValue(), types.AUDIO)
}

func TestAlarmCalendarComponent_SerializeToICSFormat(t *testing.T) {
	myTime := time.Now()
	myTimeToString := myTime.Format("20060102T150405Z")

	action := properties.NewActionProperty(types.AUDIO)
	trigger := properties.NewTriggerProperty(myTime, "PT5M", types.WithUtcTime)
	description := properties.NewDescriptionProperty("Test Alarm")
	summary := properties.NewSummaryProperty("Alarm Summary")

	component := NewAlarmCalendarComponent(action, trigger, description, summary)

	var buf bytes.Buffer
	component.SerializeToICSFormat(&buf)
	serialized := buf.String()

	expected := fmt.Sprintf("BEGIN:VALARM\r\nACTION:AUDIO\r\nTRIGGER:%s\r\nDESCRIPTION:Test Alarm\r\nSUMMARY:Alarm Summary\r\nEND:VALARM\r\n", myTimeToString)
	assert.Equal(t, expected, serialized)
}
