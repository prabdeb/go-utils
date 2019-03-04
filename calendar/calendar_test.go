package calendar

import (
	"testing"
	"time"
)

func TestNewCalendar(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic when startDate was after endDate")
		}
	}()
	endDate, _ := time.Parse(time.RFC822, "01 Jan 18 12:00 IST")
	startDate, _ := time.Parse(time.RFC822, "01 Apr 18 12:00 IST")
	NewCalendar(startDate, endDate)
}

func TestMonthBetweenTimes(t *testing.T) {
	startDate, _ := time.Parse(time.RFC822, "01 Jan 18 12:00 IST")
	endDate, _ := time.Parse(time.RFC822, "01 Apr 18 12:00 IST")
	calendarObj := NewCalendar(startDate, endDate)
	if calendarObj.Month() != 4 {
		t.Errorf("Month check expected 4 but got - %d", calendarObj.Month())
	}
	startDate, _ = time.Parse(time.RFC822, "01 Jan 18 12:00 IST")
	endDate, _ = time.Parse(time.RFC822, "31 Jan 18 12:00 IST")
	calendarObj = NewCalendar(startDate, endDate)
	if calendarObj.Month() != 1 {
		t.Errorf("Month check expected 1 but got - %d", calendarObj.Month())
	}
}

func TestDayBetweenTimes(t *testing.T) {
	startDate, _ := time.Parse(time.RFC822, "01 Jan 18 12:00 IST")
	endDate, _ := time.Parse(time.RFC822, "01 Apr 18 12:00 IST")
	calendarObj := NewCalendar(startDate, endDate)
	if calendarObj.Day() != 91 {
		t.Errorf("Day check expected 90 but got - %d", calendarObj.Day())
	}
	startDate, _ = time.Parse(time.RFC822, "01 Jan 18 12:00 IST")
	endDate, _ = time.Parse(time.RFC822, "31 Jan 18 12:00 IST")
	calendarObj = NewCalendar(startDate, endDate)
	if calendarObj.Day() != 31 {
		t.Errorf("Day check expected 31 but got - %d", calendarObj.Day())
	}
}

func TestWeekBetweenTimes(t *testing.T) {
	startDate, _ := time.Parse(time.RFC822, "01 Jan 19 12:00 IST")
	endDate, _ := time.Parse(time.RFC822, "01 Apr 19 12:00 IST")
	calendarObj := NewCalendar(startDate, endDate)
	if calendarObj.Week() != 14 {
		t.Errorf("Week check expected 14 but got - %d", calendarObj.Week())
	}
	startDate, _ = time.Parse(time.RFC822, "01 Jan 19 12:00 IST")
	endDate, _ = time.Parse(time.RFC822, "31 Jan 19 12:00 IST")
	calendarObj = NewCalendar(startDate, endDate)
	if calendarObj.Week() != 5 {
		t.Errorf("Week check expected 14 but got - %d", calendarObj.Week())
	}
}
