# GICS 📅

[![Go Reference](https://pkg.go.dev/badge/github.com/vareversat/gics.svg)](https://pkg.go.dev/github.com/vareversat/gics)

> [!WARNING]
> There is still plenty work to do on this project and it certenaly cannot be used in a production environment 

This project is intended to implement the **RFC5545 - Internet Calendaring and Scheduling Core Object Specification** (you can find the *Request Of Comments* [here](https://datatracker.ietf.org/doc/html/rfc5545))

## The genesis 🧠
I decided to create this go module because existing options didn't meet my needs or were no longer maintained.

## Main features 🚀
Here the main features this module offers (some are already available ✅ and some are still in WIP state 🟧)

  - ✅ Create an **iCalendar** object
    - ✅ With **VEVENT** components
    - ✅ With **VALARM** components
    - ✅ With **VJOURNAL** components
    - ✅ With **VFREEBUSY** components
    - ✅ With **VTIMEZONE** components
    - ✅ With **VTODO** components
  - 🟧 Make use of all [properties available](https://datatracker.ietf.org/doc/html/rfc5545#section-3.2)
  - 🟧 Make use of all [types available](https://datatracker.ietf.org/doc/html/rfc5545#section-3.3)
  - 🟧 Parse an iCalendar file
  - ✅ Write in an iCalendar file

## Installation

To install gics, use the `go get` command:

```sh
go get github.com/vareversat/gics
```

## Example 

 - Create and display an iCalendar

```go
package main

import (
	"github.com/vareversat/gics/pkg"
	"github.com/vareversat/gics/pkg/components"
	"github.com/vareversat/gics/pkg/properties"
	"github.com/vareversat/gics/pkg/types"
	"os"
	"time"
)

func main() {
	// Create an array of VEVENT component
	v := components.NewEventCalendarComponent(
		properties.NewUidProperty("My_EVENT"),
		properties.NewDateTimeStampProperty(time.Now(), types.WithUtcTime),
		properties.NewDescriptionProperty("This is an event of my calendar !"))
	c := make(components.CalendarComponents, 0)
	c = append(c, v)
	
	// Create an iCalendar with the previously created VEVENT
	calendar, err := pkg.NewCalendar(c, "My_CALENDAR", "PUBLISH", "2.0")
	if err != nil {
		panic(err)
	} else {
		// Print the iCalendar in the stdout
		calendar.SerializeToICSFormat(os.Stdout)
	}
}
```

