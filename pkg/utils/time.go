package utils

import (
	"fmt"
	"time"
)

func ParseAndValidateTime(startEpoch, endEpoch int64) (*time.Time, *time.Time, error) {
	// Parse time
	start := time.Unix(startEpoch, 0)
	end := time.Unix(endEpoch, 0)

	// Validate
	if !IsInTheFuture(start) {
		return nil, nil, fmt.Errorf("event cannot start in the past")
	}

	if !TimeIsBeforeEnd(start, end) {
		return nil, nil, fmt.Errorf("invalid start/end time(s). event is ending before it starts")
	}

	if !IsInHourRange(start, end, 2, 12) {
		return nil, nil, fmt.Errorf("event duration failed ot meet requirement: 2 <= n <= 12")
	}

	return &start, &end, nil
}

func IsInTheFuture(start time.Time) bool {
	return start.After(time.Now())
}

func TimeIsBeforeEnd(start, end time.Time) bool {
	return start.Before(end)
}

func IsInHourRange(start, end time.Time, minHours, maxHours uint) bool {
	diff := end.Sub(start)
	return uint(diff.Hours()) >= minHours && uint(diff.Hours()) <= maxHours
}
