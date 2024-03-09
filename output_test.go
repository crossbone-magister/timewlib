package timewlib

import (
	"testing"
)

func TestGenerateNoDataMessageComplete(t *testing.T) {
	var config Configuration = map[string]string{"temp.report.start": "20240308T230000Z", "temp.report.end": "20240309T230000Z"}
	message := GenerateNoDataMessage(config)
	t.Log(message)
	expected := "No filtered data found in the range 2024-03-09T00:00:00 - 2024-03-10T00:00:00."
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
