package booking

import (
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/02/2006 15:04:05"
	dateTime, err := time.Parse(layout, date)
	if err != nil {
		panic("Cannot format provided date")
	}
	return dateTime
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	dateTime, err := time.Parse(layout, date)
	if err != nil {
		panic("Wront date format")
	}
	return time.Now().After(dateTime)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	dateTime, err := time.Parse(layout, date)
	if err != nil {
		panic("Wrong date format")
	}
	if dateTime.Hour() >= 12 && dateTime.Hour() < 18 {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	inpoutLayout := "1/2/2006 15:04:05"
	dateTime, err := time.Parse(inpoutLayout, date)
	if err != nil {
		panic("Wrong date format")
	}
	outputLayout := "You have an appointment on Monday, January 2, 2006, at 15:04."
	return dateTime.Format(outputLayout)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
