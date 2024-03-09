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
