package timewlib

import (
	"fmt"
)

const c_OUTPUT_FORMAT = "2006-01-02T15:04:05"

func GenerateNoDataMessage(config Configuration) string {
	startDate, startErr := config.GetReportStartDate()
	endDate, endErr := config.GetReportEndDate()
	if startErr != nil || endErr != nil {
		return "No filtered data found."
	}
	return fmt.Sprintf("No filtered data found in the range %s - %s.", startDate.Format(c_OUTPUT_FORMAT), endDate.Format(c_OUTPUT_FORMAT))
}
