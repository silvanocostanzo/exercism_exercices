package booking

import (
	"fmt"
	"log"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		log.Fatal(err)
	}
	return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		log.Fatal(err)
	}
	return time.Now().After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		log.Fatal(err)
	}
	hour := t.Hour()
	return hour < 20 && hour > 11
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	t := Schedule(date)
	layout := "Monday, January 2, 2006, at 15:04"
	apt := t.Format(layout)

	return fmt.Sprintf("You have an appointment on %s.", apt)
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	layout := "2006/01/02 15:04:05"
	date := "2022/09/15 00:00:00"
	t, err := time.Parse(layout, date)
	if err != nil {
		log.Fatal(err)
	}
	return t
}
