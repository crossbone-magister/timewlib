package timewlib

import (
	"strings"
	"time"
)

const separator = "."
const debugKey = "debug"
const verboseKey = "verbose"
const confirmationKey = "confirmation"
const reportStartKey = "temp.report.start"
const reportEndKey = "temp.report.end"

type Configuration map[string]string

func (c Configuration) IsDebug() bool {
	return isEnabledValue(c[debugKey])
}

func (c Configuration) IsVerbose() bool {
	return isEnabledValue(c[verboseKey])
}

func (c Configuration) RequireConfirmation() bool {
	return isEnabledValue(c[confirmationKey])
}

func (c Configuration) GetReportStartDate() (time.Time, error) {
	start := c[reportStartKey]
	return parseIsoLocal(start)
}

func (c Configuration) GetReportEndDate() (time.Time, error) {
	end := c[reportEndKey]
	return parseIsoLocal(end)
}

func (c Configuration) GetAllByPrefix(prefix string) map[string]string {
	result := make(map[string]string)
	for key, value := range c {
		if strings.HasPrefix(key, prefix) {
			result[key] = value
		}
	}
	return result
}

func (c Configuration) GetAllByPrefixStripped(prefix string) map[string]string {
	var withPrefix = c.GetAllByPrefix(prefix)
	var result = make(map[string]string)
	for key, value := range withPrefix {
		result[strings.TrimPrefix(key, prefix+separator)] = value
	}
	return result
}

func isEnabledValue(value string) bool {
	return value == "on" || value == "yes" || value == "true" || value == "y" || value == "1"
}
