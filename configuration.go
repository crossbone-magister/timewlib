package timewlib

import "time"

const c_DEBUG_KEY = "debug"
const c_VERBOSE_KEY = "verbose"
const c_CONFIRMATION_KEY = "confirmation"
const c_REPORT_START_KEY = "temp.report.start"
const c_REPORT_END_KEY = "temp.report.end"

type Configuration map[string]string

func (c Configuration) IsDebug() bool {
	return isEnabledValue(c[c_DEBUG_KEY])
}

func (c Configuration) IsVerbose() bool {
	return isEnabledValue(c[c_VERBOSE_KEY])
}

func (c Configuration) RequireConfirmation() bool {
	return isEnabledValue(c[c_CONFIRMATION_KEY])
}

func (c Configuration) GetReportStartDate() (time.Time, error) {
	start := c[c_REPORT_START_KEY]
	return parseIsoLocal(start)
}

func (c Configuration) GetReportEndDate() (time.Time, error) {
	end := c[c_REPORT_END_KEY]
	return parseIsoLocal(end)
}

func isEnabledValue(value string) bool {
	return value == "on" || value == "yes" || value == "true" || value == "y" || value == "1"
}
