package timewlib

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerateNoDataMessageComplete(t *testing.T) {
	var config Configuration = map[string]string{"temp.report.start": "20240308T230000Z", "temp.report.end": "20240309T230000Z"}
	message := GenerateNoDataMessage(config)
	t.Log(message)
	startDate := time.Date(2024, 3, 8, 23, 0, 0, 0, time.UTC).Local().Format("2006-01-02T15:04:05")
	endDate := time.Date(2024, 3, 9, 23, 0, 0, 0, time.UTC).Local().Format("2006-01-02T15:04:05")
	expected := fmt.Sprintf("No filtered data found in the range %s - %s.", startDate, endDate)
	if message != expected {
		t.Fatalf("Actual message [%s] is different from expected [%s]", message, expected)
	}
}

func TestGenerateNoDataMessageNoEndDate(t *testing.T) {
	var config Configuration = map[string]string{"temp.report.start": "20240308T230000Z"}
	message := GenerateNoDataMessage(config)
	t.Log(message)
	expected := "No filtered data found."
	if message != expected {
		t.Fatalf("Actual message [%s] is different from expected [%s]", message, expected)
	}
}

func TestGenerateNoDataMessageNoStartDate(t *testing.T) {
	var config Configuration = map[string]string{"temp.report.end": "20240309T230000Z"}
	message := GenerateNoDataMessage(config)
	t.Log(message)
	expected := "No filtered data found."
	if message != expected {
		t.Fatalf("Actual message [%s] is different from expected [%s]", message, expected)
	}
}
