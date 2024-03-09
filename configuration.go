package timewlib

const c_DEBUG_KEY = "debug"
const c_VERBOSE_KEY = "verbose"
const c_CONFIRMATION_KEY = "confirmation"

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

func isEnabledValue(value string) bool {
	return value == "on" || value == "yes" || value == "true" || value == "y" || value == "1"
}
