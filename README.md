# GICS 📅

[![Go badge for gics Github project](https://img.shields.io/badge/go-white?logo=go&style=for-the-badge)](https://go.dev)
[![Build badge for gics Github project](https://img.shields.io/github/actions/workflow/status/vareversat/gics/push.default.yaml?logo=github&style=for-the-badge)](https://github.com/vareversat/gics/actions)
[![Last release badge for gics Github project](https://img.shields.io/github/v/tag/vareversat/gics?label=version&logo=git&logoColor=white&style=for-the-badge)](https://github.com/vareversat/gics/releases)
[![Code coverage badge for gics Github project](https://img.shields.io/codecov/c/github/vareversat/gics?logo=codecov&style=for-the-badge&token=ES462W8D70)](https://codecov.io/gh/vareversat/gics/)

> [!WARNING]
> There is still plenty work to do on this project. This package may be uses in a production environment at your own
> risks !

This project is intended to implement the **RFC5545 - Internet Calendaring and Scheduling Core Object Specification** (
you can find the *Request Of Comments* [here](https://datatracker.ietf.org/doc/html/rfc5545))

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
  - ✅ All are implemented
  - 🟧 Work needed for **EXRULE**, **FREEBUSY**, **RDATE**, **RRULE** and **TRIGGER**
- 🟧 Make use of all [types available](https://datatracker.ietf.org/doc/html/rfc5545#section-3.3)
  - ✅ All are implemented
  - 🟧 Work needed for **RECUR**
- 🟧 Parse an iCalendar file
  - ✅ Parsing ics files is implemented
  - 🟧 Missing the struct computation
- ✅ Write in an iCalendar file

## Installation

To install gics, use the `go get` command:

```sh
go get github.com/vareversat/gics@latest
```

## Type representations

|                                 Type name                                  | Go type                          |
|:--------------------------------------------------------------------------:|----------------------------------|
|   [BINARY](https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.1)    | `string` (base64 representation) |
|   [BOOLEAN](https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.2)   | `bool`                           |
| [CAL-ADDRESS](https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.3) | `uri.URL`                        |
|    [DATE](https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.4)     | `time.Time`                      |
|  [DATE-TIME](https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.5)  | `time.Time`                      |
|  [DURATION](https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.6)   | `string`                         |
|    [FLOAT](https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.7)    | `float32`                        |
|   [INTEGER](https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.8)   | `int32`                          |
|   [PERIOD](https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.9)    | `time.Time` / `time.Time`        |
|   [RECUR](https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10)    | *complex*                        |
|    [TEXT](https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.11)    | `string`                         |
|    [TIME](https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.12)    | `time.Time`                      |
|    [URI](https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.13)     | `uri.URL`                        |
| [UTC-OFFSET](https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.14) | `string`                         |

## Example

- Create and display an iCalendar

```go
package main

import (
  "github.com/vareversat/gics"
  "github.com/vareversat/gics/components"
  "github.com/vareversat/gics/properties"
  "github.com/vareversat/gics/types"
  "os"
  "time"
)

func main() {
  // Create an array of VEVENT component
  calendarComponents := components.CalendarComponents{}
  // Create a VEVENT component
  c := components.NewEventCalendarComponent(
    properties.NewUidProperty("My_EVENT"),
    properties.NewDateTimeStampProperty(time.Now().UTC()),
    []components.AlarmCalendarComponent{}, // Empty VALARM
    properties.NewDateTimeStartProperty(time.Now(), types.WithUtcTime),
    properties.NewDescriptionProperty("This is an event of my calendar !"))
  calendarComponents = append(calendarComponents, c)

  // Create an iCalendar with the previously created VEVENT
  calendar, err := gics.NewCalendar(calendarComponents,
    "-//Valentin REVERSAT//https://github.com/vareversat/gics//FR",
    "PUBLISH",
    "2.0")
  if err != nil {
    panic(err)
  } else {
    // Print the iCalendar in the stdout
    calendar.SerializeToICSFormat(os.Stdout)
  }
}

```
