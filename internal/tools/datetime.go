package tools

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ParseDatetime(input string) (string, error) {
	input = strings.TrimSpace(input)

	if isTimestamp(input) {
		return parseTimestamp(input)
	}

	return parseTimeString(input)
}

func isTimestamp(input string) bool {
	matched, _ := regexp.MatchString(`^\d+$`, input)
	return matched
}

func parseTimestamp(input string) (string, error) {
	ts, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return "", fmt.Errorf("invalid timestamp: %v", err)
	}

	var t time.Time
	if len(input) > 10 {
		t = time.UnixMilli(ts)
	} else {
		t = time.Unix(ts, 0)
	}

	return t.Format("2006.01.02 15:04:05 MST"), nil
}

func parseTimeString(input string) (string, error) {
	input = strings.ReplaceAll(input, ".", "-")

	layouts := []string{
		"2006-01-02 15:04:05 MST",
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04:05",
		time.RFC3339,
		time.RFC3339Nano,
	}

	var t time.Time
	var err error

	for _, layout := range layouts {
		t, err = time.Parse(layout, input)
		if err == nil {
			break
		}
	}

	if err != nil {
		return "", fmt.Errorf("unable to parse time: %s", input)
	}

	return fmt.Sprintf("%d", t.Unix()), nil
}

func GetCurrentTimestamp() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}
