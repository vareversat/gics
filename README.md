# gics - Go Internet Calendar Scheduling

> [!WARNING]
> There is still plenty work to do on this project and it certenaly cannot be used in a production environment 

This project is intended to implement the **RFC5545 - Internet Calendaring and Scheduling Core Object Specification** (you can find the *Request Of Comments* [here](https://datatracker.ietf.org/doc/html/rfc5545))

## Installation

To install gics, use the `go get` command:

```sh
go get github.com/vareversat/gics
```

## Example 

 - Create a Calendar

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
	v := components.NewEventCalendarComponent(
		properties.NewUidProperty("My_EVENT"),
		properties.NewDateTimeStampProperty(time.Now(), types.WithUtcTime),
		properties.NewDescriptionProperty("This is an event of my calendar !"))
	c := make(components.Components, 0)
	c = append(c, v)
	calendar, err := pkg.NewCalendar(c, "My_CALENDAR", "PUBLISH", "2.0")
	if err != nil {
		panic(err)
	} else {
		calendar.ToICalendarFormat(os.Stdout)

	}
}
```

