package timewlib

import (
	"bytes"
	"log"
	"testing"
)

func TestSetupLogging(t *testing.T) {
	var config Configuration = map[string]string{"debug": "off"}
	buffer := bytes.NewBufferString("")
	log.SetOutput(buffer)
	SetupLogging(config)
	log.Println("Test message")
	if buffer.String() != "" {
		t.Errorf("Log output was not discarded when debug is off: %s", buffer.String())
	}
}
