package timewlib

import (
	"testing"
	"time"
)

func TestProcess(t *testing.T) {
	testCase := []TimewarriorInterval{
		{
			Start: "20230101T000000Z",
			End:   "20230101T003000Z",
			Tags:  []string{"tag1", "tag2"},
		},
	}
	processed, err := Process(testCase)
	if err != nil {
		t.Errorf("Error occurred while processing %v", err)
	}
	if len(processed) != 1 {
		t.Errorf("Processor Failed to read all records")
	}
	interval := processed[0]
	if len(interval.Tags) != 2 {
		t.Errorf("Processor failed to process interval tags")
	}
	startTime, err := time.Parse("20060102T150405Z", "20230101T000000Z")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	startTime = startTime.Local()
	if interval.start != startTime {
		t.Errorf("Expected start date to be [%v], found instead [%v]", startTime, interval.start)
	}
	endTime, err := time.Parse("20060102T150405Z", "20230101T003000Z")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	endTime = endTime.Local()
	if interval.end != endTime {
		t.Errorf("Expected start date to be [%v], found instead [%v]", endTime, interval.end)
	}
}

func TestProcessStartDateError(t *testing.T) {
	testCase := []TimewarriorInterval{
		{
			Start: "THISISNOTADATE",
			End:   "20230101T003000Z",
			Tags:  []string{"tag1", "tag2"},
		},
	}
	_, err := Process(testCase)
	if err == nil {
		t.Errorf("Processor did not return error for invalid string date")
	} else {
		t.Logf("%v", err)
	}
}

func TestProcessEndDateEmpty(t *testing.T) {
	testCase := []TimewarriorInterval{
		{
			Start: "20230101T003000Z",
			End:   "",
			Tags:  []string{"tag1", "tag2"},
		},
	}
	intervals, err := Process(testCase)
	if err != nil {
		t.Fatalf("Processor returned unexpected error")
	}
	interval := intervals[0]
	if interval.end.After(time.Now()) {
		t.Errorf("Expected end date to be before now, found instead [%v]", interval.end)
	}
}

func TestProcessEndDateError(t *testing.T) {
	testCase := []TimewarriorInterval{
		{
			Start: "20230101T003000Z",
			End:   "THISISNOTADATE",
			Tags:  []string{"tag1", "tag2"},
		},
	}
	_, err := Process(testCase)
	if err == nil {
		t.Errorf("Processor did not return error for invalid string date")
	} else {
		t.Logf("%v", err)
	}
}

func TestParseIsoLocal(t *testing.T) {
	parsedDate, err := parseIsoLocal("20230101T003000Z")
	if err != nil {
		t.Errorf("Error while parsing date")
	}
	expectedDate := time.Date(2023, 01, 01, 01, 30, 0, 0, time.Local)
	if parsedDate != expectedDate {
		t.Errorf("Expected start date to be [%v], found instead [%v]", expectedDate, parsedDate)
	}
}

func TestParseIsoLocalError(t *testing.T) {
	_, err := parseIsoLocal("THISISNOTADATE")
	if err == nil {
		t.Errorf("parseIsoLocal did not return error for invalid string date")
	} else {
		t.Logf("%v", err)
	}
}
