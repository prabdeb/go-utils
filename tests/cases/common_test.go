package cases

import (
	"testing"
	"time"

	"github.com/prabdeb/go-utils/common"
)

func TestStringContains(t *testing.T) {
	if !common.Contains([]string{"My", "Awesome", "Search", "Util"}, "My") {
		t.Error("String array search did not worked for true case")
	}
	if common.Contains([]string{"My", "Awesome", "Search", "Util"}, "Hello") {
		t.Error("String array search did not worked for false case")
	}
}

func TestIntContains(t *testing.T) {
	if !common.Contains([]int{1, 2, 3, 4, 15, 16, 18}, 16) {
		t.Error("Int array search did not worked for true case")
	}
	if common.Contains([]int{1, 2, 3, 4, 15, 16, 18}, 100) {
		t.Error("Int array search did not worked for false case")
	}
}

func TestNewCalendar(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic when startDate was after endDate")
		}
	}()
	endDate, _ := time.Parse(time.RFC822, "01 Jan 18 12:00 IST")
	startDate, _ := time.Parse(time.RFC822, "01 Apr 18 12:00 IST")
	common.NewCalendar(startDate, endDate)
}

func TestMonthBetweenTimes(t *testing.T) {
	startDate, _ := time.Parse(time.RFC822, "01 Jan 18 12:00 IST")
	endDate, _ := time.Parse(time.RFC822, "01 Apr 18 12:00 IST")
	calendar := common.NewCalendar(startDate, endDate)
	if calendar.Month() != 4 {
		t.Errorf("Month check expected 4 but got - %d", calendar.Month())
	}
	startDate, _ = time.Parse(time.RFC822, "01 Jan 18 12:00 IST")
	endDate, _ = time.Parse(time.RFC822, "31 Jan 18 12:00 IST")
	calendar = common.NewCalendar(startDate, endDate)
	if calendar.Month() != 1 {
		t.Errorf("Month check expected 1 but got - %d", calendar.Month())
	}
}

func TestDayBetweenTimes(t *testing.T) {
	startDate, _ := time.Parse(time.RFC822, "01 Jan 18 12:00 IST")
	endDate, _ := time.Parse(time.RFC822, "01 Apr 18 12:00 IST")
	calendar := common.NewCalendar(startDate, endDate)
	if calendar.Day() != 91 {
		t.Errorf("Day check expected 90 but got - %d", calendar.Day())
	}
	startDate, _ = time.Parse(time.RFC822, "01 Jan 18 12:00 IST")
	endDate, _ = time.Parse(time.RFC822, "31 Jan 18 12:00 IST")
	calendar = common.NewCalendar(startDate, endDate)
	if calendar.Day() != 31 {
		t.Errorf("Day check expected 31 but got - %d", calendar.Day())
	}
}

func TestWeekBetweenTimes(t *testing.T) {
	startDate, _ := time.Parse(time.RFC822, "01 Jan 19 12:00 IST")
	endDate, _ := time.Parse(time.RFC822, "01 Apr 19 12:00 IST")
	calendar := common.NewCalendar(startDate, endDate)
	if calendar.Week() != 14 {
		t.Errorf("Week check expected 14 but got - %d", calendar.Week())
	}
	startDate, _ = time.Parse(time.RFC822, "01 Jan 19 12:00 IST")
	endDate, _ = time.Parse(time.RFC822, "31 Jan 19 12:00 IST")
	calendar = common.NewCalendar(startDate, endDate)
	if calendar.Week() != 5 {
		t.Errorf("Week check expected 14 but got - %d", calendar.Week())
	}
}
