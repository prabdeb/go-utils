package calendar

import "time"

// Calendar provides common tasks with time
type Calendar struct {
	StartDate time.Time
	EndDate   time.Time
}

// NewCalendar creates Calendar type
// startDate: the date from where calendar should begin
// endDate: the date where calendar should end, pass current time time.Now() for till lattest
func NewCalendar(startDate time.Time, endDate time.Time) *Calendar {
	if startDate.After(endDate) {
		panic("calendar config error: startDate can not be after endDate")
	}
	return &Calendar{
		StartDate: startDate,
		EndDate:   endDate,
	}
}

// Month returns the month count between two times
func (c *Calendar) Month() int {
	months := 1
	createdAtTime := c.StartDate
	month := createdAtTime.Month()
	for createdAtTime.Before(c.EndDate) {
		createdAtTime = createdAtTime.Add(time.Hour * 24)
		nextMonth := createdAtTime.Month()
		if nextMonth != month {
			months++
		}
		month = nextMonth
	}
	return months
}

// Day returns the day count between two times
func (c *Calendar) Day() int {
	days := 1
	createdAtTime := c.StartDate
	day := createdAtTime.Day()
	for createdAtTime.Before(c.EndDate) {
		createdAtTime = createdAtTime.Add(time.Hour * 24)
		nextDay := createdAtTime.Day()
		if nextDay != day {
			days++
		}
		day = nextDay
	}
	return days
}

// Week returns the week count between two times
func (c *Calendar) Week() int {
	weeks := 1
	createdAtTime := c.StartDate
	for createdAtTime.Before(c.EndDate) {
		dayOfTheWeek := createdAtTime.Weekday()
		createdAtTime = createdAtTime.Add(time.Hour * 24)
		// Considering Sunday as beging of the week
		if dayOfTheWeek == 0 {
			weeks++
		}
	}
	return weeks
}
