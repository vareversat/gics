# GICS 📅

<a href="https://go.dev"><img src="https://img.shields.io/badge/go-white?logo=go&style=for-the-badge" alt="Go badge for gics Github project"></a>
<a href="https://github.com/vareversat/gics/actions"><img src="https://img.shields.io/github/actions/workflow/status/vareversat/gics/push.default.yaml?logo=github&style=for-the-badge" alt="Build badge for gics Github project"></a>
<a href="https://github.com/vareversat/gics/releases"><img src="https://img.shields.io/github/v/tag/vareversat/gics?label=version&logo=git&logoColor=white&style=for-the-badge" alt="Last release badge for gics Github project"></a>
<a href="https://codecov.io/gh/vareversat/gics/"><img src="https://img.shields.io/codecov/c/github/vareversat/gics?logo=codecov&style=for-the-badge&token=ES462W8D70" alt="Code coverage badge for gics Github project"></a>

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

