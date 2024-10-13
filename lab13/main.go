package main

import (
	"fmt"
	"time"
)

func TimeDifference(start, end time.Time) time.Duration {
	return end.Sub(start)
}

func FormatTimeToString(timestamp time.Time, format string) string {
	return timestamp.Format(format)
}

func ParseStringToTime(dateString, format string) (time.Time, error) {
	ans, err := time.Parse(format, dateString)
	if err != nil {
		panic(err)
	}
	return ans, nil
}

func NextWorkday(start time.Time) time.Time {
	curDay := start.Format("Monday")
	switch curDay {
	case "Friday":
		return start.Add(72 * time.Hour)
	case "Saturday":
		return start.Add(48 * time.Hour)
	default:
		return start.Add(24 * time.Hour)
	}
}

func TimeAgo(pastTime time.Time) string {
	now := time.Now()
	diff := now.Sub(pastTime)

	var unit string
	var value int

	switch {
	case diff < time.Minute:
		return "just now"
	case diff < time.Hour:
		value = int(diff.Minutes())
		unit = "minute"
	case diff < 24*time.Hour:
		value = int(diff.Hours())
		unit = "hour"
	case diff < 30*24*time.Hour:
		value = int(diff.Hours() / 24)
		unit = "day"
	case diff < 365*24*time.Hour:
		value = int(diff.Hours() / (24 * 30))
		unit = "month"
	default:
		value = int(diff.Hours() / (24 * 365))
		unit = "year"
	}

	if value > 1 {
		unit += "s"
	}

	return fmt.Sprintf("%d %s ago", value, unit)
}

func main() {
	fmt.Println(NextWorkday(time.Now()))
}
