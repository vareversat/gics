package main

import (
	"fmt"

	"github.com/vareversat/gics/pkg"
)

func main() {
	event := pkg.Event{
		PRODID:      "github.com/vareversat/gics",
		VERSION:     "2.0",
		DESCRIPTION: "My event",
	}

	fmt.Println("Hello, world!")
}
