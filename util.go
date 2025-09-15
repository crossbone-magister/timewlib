package timewlib

import (
	"fmt"
	"io"
	"log"
	"os"
)

// SetupLogging configures the logging output based on the settings in the configuration.
// If debugging is not enabled on timewarrior, all log output is set to be discarded.
func SetupLogging(configuration Configuration) {
	if !configuration.IsDebug() {
		log.SetOutput(io.Discard)
	}
}

// ExitIfError checks if the provided error is non-nil.
// If it is, it prints the error message and exits the program with a non-zero status code.
func ExitIfError(err error) {
	if err != nil {
		fmt.Println("Blocking error occurred: ", err)
		log.Fatal("Blocking error occurred: ", err)
	}
}

// ExitOnNoData outputs a message indicating that no data was found based report start date and end date from the configuration exiting with a zero status code.
func ExitOnNoData(intervals []TimewarriorInterval, configuration Configuration) {
	if len(intervals) == 0 {
		fmt.Println(GenerateNoDataMessage(configuration))
		os.Exit(0)
	}
}
