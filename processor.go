package timewlib

import (
	"time"
)

const isoLayout = "20060102T150405Z"

func parseIsoLocal(date string) (time.Time, error) {
	t, err := time.Parse(isoLayout, date)
	if err != nil {
		return time.Time{}, err
	}
	return t.Local(), nil
}

func Process(rawIntervals []TimewarriorInterval) ([]Interval, error) {
	var processedIntervals []Interval
	for _, rawInterval := range rawIntervals {
		start, err := parseIsoLocal(rawInterval.Start)
		if err != nil {
			return []Interval{}, err
		}
		var end time.Time
		if rawInterval.End == "" {
			end = time.Now().UTC()
		} else {
			end, err = parseIsoLocal(rawInterval.End)
			if err != nil {
				return []Interval{}, err
			}
		}
		processedIntervals = append(processedIntervals, Interval{
			start: start,
			end:   end,
			Tags:  rawInterval.Tags,
		})
	}
	return processedIntervals, nil
}
