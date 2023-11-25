package timewlib

import (
	"fmt"
	"testing"
	"time"
)

func TestNewInterval(t *testing.T) {
	interval := NewInterval(10, 00, 10, 30)
	if interval.StartHour() != 10 {
		t.Errorf("interval.StartHour() is not correct :%d", interval.StartHour())
	}
	if interval.StartMinute() != 0 {
		t.Errorf("interval.StartMinute() is not correct :%d", interval.StartMinute())
	}
	if interval.EndHour() != 10 {
		t.Errorf("interval.EndHour() is not correct :%d", interval.EndHour())
	}
	if interval.EndMinute() != 30 {
		t.Errorf("interval.EndMinute() is not correct :%d", interval.EndMinute())
	}
	if interval.Duration() != 30 {
		t.Errorf("interval.Duration() is not correct :%f", interval.Duration())
	}
	if !interval.IsSameHour() {
		t.Errorf("interval.IsSameHour() is not correct :%t", interval.IsSameHour())
	}
	today := time.Now().Format("2006-01-02")
	timezone := time.Now().Format("Z0700 MST")
	expectedStringValue := fmt.Sprintf("[%s 10:00:00 %s-%s 10:30:00 %s]", today, timezone, today, timezone)
	if interval.String() != expectedStringValue {
		t.Errorf("interval.String() is not correct :%s != %s", expectedStringValue, interval.String())
	}
	if interval.ShortString() != "1000-1030" {
		t.Errorf("interval.ShortString() is not correct :%s", interval.ShortString())
	}
}

func TestStartDate(t *testing.T) {
	interval := NewInterval(10, 00, 10, 30)
	start_year, start_month, start_day := interval.StartDate()
	now_year, now_month, now_day := time.Now().Date()
	if now_year != start_year {
		t.Errorf("interval.StartDate() year is not correct: %d != %d", now_year, start_year)
	}
	if now_month != start_month {
		t.Errorf("interval.StartDate() year is not correct: %d != %d", now_month, start_month)
	}
	if now_day != start_day {
		t.Errorf("interval.StartDate() year is not correct: %d != %d", now_day, start_day)
	}
}

func TestEndDate(t *testing.T) {
	interval := NewInterval(10, 00, 10, 30)
	end_year, end_month, end_day := interval.EndDate()
	now_year, now_month, now_day := time.Now().Date()
	if now_year != end_year {
		t.Errorf("interval.StartDate() year is not correct: %d != %d", now_year, end_year)
	}
	if now_month != end_month {
		t.Errorf("interval.StartDate() year is not correct: %s != %s", now_month, end_month)
	}
	if now_day != end_day {
		t.Errorf("interval.StartDate() year is not correct: %d != %d", now_day, end_day)
	}
}
