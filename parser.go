package timewlib

import (
	"bufio"
	"encoding/json"
	"io"
	"strings"
)

type TimewarriorInterval struct {
	Id    uint64   `json:"id"`
	Start string   `json:"start"`
	End   string   `json:"end"`
	Tags  []string `json:"tags"`
}

type TimeWarriorInput struct {
	Configuration map[string]string
	Intervals     []TimewarriorInterval
}

func Parse(reader io.Reader) (*TimeWarriorInput, error) {
	scanner := bufio.NewScanner(reader)
	parsingConfiguration := true
	config := map[string]string{}
	var intervalsRaw []string
	var intervals []TimewarriorInterval
	for scanner.Scan() {
		row := scanner.Text()
		if parsingConfiguration {
			if len(row) > 0 {
				rawConfig := strings.Split(row, ": ")
				config[rawConfig[0]] = strings.Join(rawConfig[1:], "")
			} else {
				parsingConfiguration = false
			}
		} else {
			intervalsRaw = append(intervalsRaw, row)
		}
	}
	err := json.Unmarshal([]byte(strings.Join(intervalsRaw, "")), &intervals)
	if err != nil {
		return nil, err
	}
	return &TimeWarriorInput{
		Configuration: config,
		Intervals:     intervals,
	}, nil
}
