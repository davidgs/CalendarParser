package main

import (
	// "bytes"
	// "encoding/hex"
	"encoding/json"
	"fmt"
	// "log"
	// "regexp"
	// "strings"
	// "time"
)

var testEvent = `{"events" : [
{
  "start": "new Date('2021-07-01T10:30:00-04:00')",
  "end": "new Date('2021-07-01T11:30:00-04:00')",
  "title": "Dr Brenner - Delray office",
  "extendedProps": {
    "location": "Linton Boulevard, Linton Blvd, Delray Beach, FL, USA"
  },
  "classNames": [ "Doctors" ]
}
]}`

type singleEvent struct {
	Start         string `json:"start"`
	End           string `json:"end"`
	Title         string `json:"title"`
	ExtendedProps struct {
		Location string `json:"location"`
	} `json:"extendedProps"`
	ClassNames []string `json:"classNames"`
}

type Event struct {
	Events []struct {
		Start         string `json:"start"`
		End           string `json:"end"`
		Title         string `json:"title"`
		ExtendedProps struct {
			Location string `json:"location"`
		} `json:"extendedProps"`
		ClassNames []string `json:"classNames"`
	} `json:"events"`
}

func main() {
	var obj Event
	err := json.Unmarshal([]byte(testEvent), &obj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", obj)
	fmt.Println("Start: ", obj.Events[0].Start)
	 sev := singleEvent{}
	sev.Start = "new Date('2021-07-01T10:30:00-04:00')"
	sev.End = "new Date('2021-07-01T10:30:00-04:00')"
	sev.Title = "this is my new event"
	sev.ClassNames = make([]string, 1)
	sev.ClassNames[0] = "other class"
	sev.ExtendedProps.Location = "4216 Winding Oak Way Apex, NC 27539 USA"
	obj.Events = append(obj.Events, sev)
}