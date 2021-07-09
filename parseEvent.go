package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var newItem = `{"Start": "2021-07-08T19:30:00-04:00",
"End": "2021-07-08T20:30:00-04:00",
"Title": "Ivoryton/Downingtown Meeting",
"Content": "Topic: Ivoryton-Downingtown MeetingTime: Jun 24, 2021 07:30 PM Eastern Time (US and Canada) Please download and import the following iCalendar (.ics) files to your calendar system.Weekly:
Join Zoom Meeting https://us02web.zoom.us/j/445777874?pwd=RXNhZnJmLzR1MUFLOURweWc3RTEwUT09
Meeting ID: 445 777 874
Passcode: Ivoryton
One tap mobile+13462487799,,445777874#,,,,*91896410# US (Houston)+16699009128,,445777874#,,,,*91896410# US (San Jose)",
"Category": "calendar#event",
"Location": "https://us02web.zoom.us/j/445777874?pwd=RXNhZnJmLzR1MUFLOURweWc3RTEwUT09",
"Recurrence": "RRULE:FREQ=WEEKLY;BYDAY=TH",
"Duration": "1.0"
}`

type Incoming struct {
	Start      string `json:"Start"`
	End        string `json:"End"`
	Title      string `json:"Title"`
	Content    string `json:"Content"`
	Category   string `json:"Category"`
	Location   string `json:"Location"`
	Recurrence string `json:"Recurrence"`
	Duration   string `json:"Duration"`
}
type Rrule struct {
	Freq      string   `json:"freq,omitempty"`
	Interval  int      `json:"interval,omitempty"`
	Byweekday []string `json:"byweekday,omitempty"`
	Dtstart   string   `json:"dtstart,omitempty"`
	Until     string   `json:"until,omitempty"`
}
type Events struct {
	Start    string `json:"start,omitempty"`
	End      string `json:"end,omitempty"`
	Title    string `json:"title,omitempty"`
	Allday   string   `json:"allday,omitempty"`
	Location string `json:"location,omitempty"`
	URL      string `json:"url,omitempty"`
	Duration string `json:"duration,omitempty"`
	*Rrule   `json:"rrule,omitempty"`
}
type Calendar struct {
	Events []*Events `json:"events"`
}
func main() {
	calendar := new(Calendar)
	file, err := ioutil.ReadFile("/Users/davidgs/github.com/PatsWeb/themes/davidgs/layouts/partials/calData.html")
	if err != nil {
    log.Fatal(err)
  }
	iss := "{" + string(file) + "}"
	foo := new(Incoming)
	newItem = strings.ReplaceAll(newItem, "\n", " ")
	err = json.Unmarshal([]byte(newItem), foo)
	if err != nil {
		log.Fatal(err)
	}
	fab, err := json.MarshalIndent(foo, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(fab))
	//fmt.Println(iss)
	err = json.Unmarshal([]byte(iss), calendar)
	if err != nil {
    log.Fatal(err)
  }
	bar, err := json.MarshalIndent(calendar, "", "    ")
  if err != nil {
    log.Fatal(err)
  }
  //fmt.Println(calendar)
	fmt.Println(string(bar))
	rule := foo.Recurrence
	newRule := new(Rrule)
	rule_strings := strings.Split(rule, ":")
	if len(rule_strings) > 0 {
		sub_rule := strings.Split(rule_strings[1], ";")
		for x := 0; x < len(sub_rule); x++ {
			each_rule := strings.Split(sub_rule[x], "=")
			if strings.ToLower(each_rule[0]) == "byday" {
				// Figure out how many items in the list, make the array, and assign values
				// newRule.Byweekday = strings.ToLower(each_rule[1])
			} else {
				rr := strings.ToLower(each_rule[0])
				newRule.rr = strings.ToLower(each_rule[1])
			}
		}
	}
}