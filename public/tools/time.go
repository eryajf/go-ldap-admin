package tools

import "time"

func ConvertToTimestamp(t time.Time) int64 {
	timestamp := t.Unix()
	return timestamp
}

func CalculateDaysBetweenTimestamps(timestamp1, timestamp2 int64) int {
	t1 := time.Unix(timestamp1, 0)
	t2 := time.Unix(timestamp2, 0)

	duration := t2.Sub(t1)
	days := int(duration.Hours() / 24)
	return days
}
