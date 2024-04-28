package components

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vareversat/gics/pkg/properties"
	"github.com/vareversat/gics/pkg/types"
)

func TestNewFreeBusyCalendarComponent(t *testing.T) {
	component := NewFreeBusyCalendarComponent(
		properties.NewUidProperty("UID"),
		properties.NewDateTimeStampProperty(time.Now(), types.WithUtcTime),
	)

	assert.NotNil(t, component)
}

func TestFreeBusyCalendarComponent_SerializeToICSFormat(t *testing.T) {
	myTime := time.Now()
	myTimeToString := myTime.Format("20060102T150405Z")

	uid := properties.NewUidProperty("My FreeBusy")
	dateTimestamp := properties.NewDateTimeStampProperty(myTime, types.WithUtcTime)

	component := NewFreeBusyCalendarComponent(uid, dateTimestamp)

	var buf bytes.Buffer
	component.SerializeToICSFormat(&buf)
	serialized := buf.String()

	expected := fmt.Sprintf(
		"BEGIN:VFREEBUSY\r\nUID:My FreeBusy\r\nDTSTAMP:%s\r\nEND:VFREEBUSY\r\n",
		myTimeToString,
	)
	assert.Equal(t, expected, serialized)
}
