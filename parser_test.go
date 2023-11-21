package timewlib

import (
	"fmt"
	"strings"
	"testing"
)

func TestParseFormOnlyConfig(t *testing.T) {
	var testCase string = `journal.size: -1
reports.day.axis: internal
reports.day.cell: 15
reports.day.day: yes
reports.day.holidays: no
reports.day.hours: all
reports.day.lines: 2
reports.day.month: no
reports.day.spacing: 1
reports.day.summary: yes
reports.day.totals: no
reports.day.week: no
reports.day.weekday: yes
reports.month.cell: 15
reports.month.day: yes
reports.month.holidays: yes
reports.month.hours: all
reports.month.lines: 1
reports.month.month: yes
reports.month.spacing: 1
reports.month.summary: yes
reports.month.totals: yes
reports.month.week: yes
reports.month.weekday: yes
reports.summary.holidays: yes
reports.week.cell: 15
reports.week.day: yes
reports.week.holidays: yes
reports.week.hours: all
reports.week.lines: 1
reports.week.month: no
reports.week.spacing: 1
reports.week.summary: yes
reports.week.totals: yes
reports.week.week: yes
reports.week.weekday: yes
temp.config: /path/to/config
temp.db: /path/to/temp
temp.extensions: /path/to/extensions
temp.report.end: 20230712T220000Z
temp.report.start: 20230711T220000Z
temp.report.tags:
temp.version: 1.5.0
test: test
theme.colors.exclusion: gray8 on gray4
theme.colors.holiday: gray4
theme.colors.label: gray4
theme.colors.today: white
theme.description: Built-in default
verbose: on

[
]`
	reader := strings.NewReader(testCase)
	parsed, err := Parse(reader)
	if err != nil {
		t.Errorf("Parser failed parsing test case %v", err)
	}
	configurationLength := len(parsed.Configuration)
	expectedConfigurationLength := 50
	if configurationLength != expectedConfigurationLength {
		t.Errorf("%d configuration elements where loaded instead of %d", configurationLength, expectedConfigurationLength)
	}
	if len(parsed.Intervals) != 0 {
		t.Error("Intervals are not empty")
	}
}

func TestParseFromConfigurationWithColon(t *testing.T) {
	var testCase = `key: value:with:colons

[
]`
	reader := strings.NewReader(testCase)
	parsed, err := Parse(reader)
	if err != nil {
		t.Errorf("Parser failed parsing test case %v", err)
	}
	if parsed.Configuration["key"] != "value:with:colons" {
		t.Errorf("Parser is not correctly handling colons in values of configuration. Actual value is [%s]", parsed.Configuration["key"])
	}
}

func TestParseFromConfigurationEmpty(t *testing.T) {
	var testCase = `

[
]`
	reader := strings.NewReader(testCase)
	parsed, err := Parse(reader)
	if err != nil {
		t.Errorf("Parser failed parsing test case %v", err)
	}
	if len(parsed.Configuration) != 0 {
		t.Errorf("Configuration is not empty")
	}
}

func TestParseFromConfigurationWithIntervals(t *testing.T) {
	var testCase = `

[
	{"id":1,"start":"20230101T000000Z","end":"20230101T003000Z","tags":["tag1","tag2","tag3"]},
	{"id":2,"start":"20230101T003000Z","end":"20230101T010000Z","tags":["tag1","tag2","tag3"]}
]`
	reader := strings.NewReader(testCase)
	parsed, err := Parse(reader)
	if err != nil {
		t.Errorf("Parser failed parsing test case %v", err)
	}
	if len(parsed.Intervals) != 2 {
		t.Errorf("Parser did not load all intervals")
	}
	first := parsed.Intervals[0]
	if first.Start != "20230101T000000Z" {
		t.Errorf("Expected start to be [20230101T000000Z] found %s instead", first.Start)
	}
	if first.End != "20230101T003000Z" {
		t.Errorf("Expected start to be [20230101T003000Z] found %s instead", first.End)
	}
	if len(first.Tags) != 3 {
		t.Errorf("Parser did not load all tags for interval")
	}
	for i := 0; i < 3; i++ {

		if first.Tags[i] != fmt.Sprintf("tag%d", i+1) {
			t.Errorf("Parser did not properly load tag number %d, actual value is %s", i, first.Tags[i])
		}
	}
}

func TestParseFromWithError(t *testing.T) {
	var testCases = []string{"", `

`,
		`
key: value

{}
`,
		`
key: value

[
{}
]
`}
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			reader := strings.NewReader(testCase)
			_, err := Parse(reader)
			if err == nil {
				t.Errorf("Parser failed to detect invalid configuration [%s]", testCase)
			} else {
				t.Logf("%v", err)
			}
		})
	}

}
