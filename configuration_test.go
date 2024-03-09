package timewlib

import (
	"fmt"
	"testing"
)

var testCases = map[string]bool{
	"on":    true,
	"yes":   true,
	"true":  true,
	"y":     true,
	"1":     true,
	"off":   false,
	"no":    false,
	"false": false,
	"0":     false,
	"test":  false,
}

func TestIsDebug(t *testing.T) {
	for value, expected := range testCases {
		t.Run(fmt.Sprintf("%s-%t", value, expected), func(t *testing.T) {
			var config Configuration = map[string]string{"debug": value}
			if config.IsDebug() != expected {
				t.Errorf("Test with value [%s] was not [%t]", value, expected)
			}
		})
	}
}

func TestIsVerbose(t *testing.T) {
	for value, expected := range testCases {
		t.Run(fmt.Sprintf("%s-%t", value, expected), func(t *testing.T) {
			var config Configuration = map[string]string{"verbose": value}
			if config.IsVerbose() != expected {
				t.Errorf("Test with value [%s] was not [%t]", value, expected)
			}
		})
	}
}

func TestRequireConfirmation(t *testing.T) {
	for value, expected := range testCases {
		t.Run(fmt.Sprintf("%s-%t", value, expected), func(t *testing.T) {
			var config Configuration = map[string]string{"confirmation": value}
			if config.RequireConfirmation() != expected {
				t.Errorf("Test with value [%s] was not [%t]", value, expected)
			}
		})
	}
}

func TestGetReportStartDate(t *testing.T) {
	var config Configuration = map[string]string{"temp.report.start": "20240308T230000Z"}
	start, err := config.GetReportStartDate()
	if err != nil {
		t.Fatal(err)
	}
	if start.Day() != 9 || start.Month() != 3 || start.Year() != 2024 || start.Hour() != 0 || start.Minute() != 0 {
		t.Fatalf("Parsed date is not 20240308T230000Z but %s", start)
	}
}

func TestGetReportEndDate(t *testing.T) {
	var config Configuration = map[string]string{"temp.report.end": "20240309T230000Z"}
	start, err := config.GetReportEndDate()
	if err != nil {
		t.Fatal(err)
	}
	if start.Day() != 10 || start.Month() != 3 || start.Year() != 2024 || start.Hour() != 0 || start.Minute() != 0 {
		t.Fatalf("Parsed date is not 20240309T230000Z but %s", start)
	}
}
