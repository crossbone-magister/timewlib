package timewlib

import (
	"fmt"
	"testing"
	"time"
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
	var expected = time.Date(2024, 3, 8, 23, 0, 0, 0, time.UTC).Local()
	if err != nil {
		t.Fatal(err)
	}
	if expected != start {
		t.Fatalf("Parsed date is not %s but %s", expected, start)
	}
}

func TestGetReportEndDate(t *testing.T) {
	var config Configuration = map[string]string{"temp.report.end": "20240309T230000Z"}
	end, err := config.GetReportEndDate()
	var expected = time.Date(2024, 3, 9, 23, 0, 0, 0, time.UTC).Local()
	if err != nil {
		t.Fatal(err)
	}
	if expected != end {
		t.Fatalf("Parsed date is not %s but %s", expected, end)
	}
}

func TestGetAllByPrefix(t *testing.T) {
	var config Configuration = map[string]string{"p1.a": "1", "p1.b": "2", "p2.c": "3", "d": "4"}
	allP1 := config.GetAllByPrefix("p1")
	expectedEntries := 2
	if len(allP1) != expectedEntries {
		t.Fatalf("Configuration did not extract expected [%d] but [%d] entries instead with prefix p1", expectedEntries, len(allP1))
	}
	var keys = []string{"p1.a", "p1.b"}
	for _, key := range keys {
		if _, ok := allP1[key]; !ok {
			t.Fatalf("Configuration did not contain expected key %s", key)
		}
	}
}

func TestGetAllByPrefixStripped(t *testing.T) {
	var config Configuration = map[string]string{"p1.a": "1", "p1.b": "2", "p2.c": "3", "d": "4"}
	allP1 := config.GetAllByPrefixStripped("p1")
	expectedEntries := 2
	if len(allP1) != expectedEntries {
		t.Fatalf("Configuration did not extract expected [%d] but [%d] entries instead with prefix p1", expectedEntries, len(allP1))
	}
	var keys = []string{"a", "b"}
	for _, key := range keys {
		if _, ok := allP1[key]; !ok {
			t.Fatalf("Configuration did not contain expected key %s", key)
		}
	}
	for key := range config {
		if _, ok := allP1[key]; ok {
			t.Fatalf("Configuration contained unexpected key %s", key)
		}
	}
}
