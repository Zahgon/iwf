package timeparser

import (
	"time"
)

const (
	DateTimeFormat = "2006-01-02T15:04:05-07:00" // used for converting UnixNano to string like 2018-02-15T16:16:36-08:00
	// regex expression for parsing time durations, shorter, longer notations and numeric value respectively
	defaultDateTimeRangeShortRE = "^[1-9][0-9]*[smhdwMy]$"                                // eg. 1s, 20m, 300h etc.
	defaultDateTimeRangeLongRE  = "^[1-9][0-9]*(second|minute|hour|day|week|month|year)$" // eg. 1second, 20minute, 300hour etc.
	defaultDateTimeRangeNum     = "^[1-9][0-9]*"                                          // eg. 1, 20, 300 etc.
)

func ParseTime(timeStr string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// try to parse

// treat as raw time

// treat as time range format

// parseTimeRange parses a given time duration string (in format X<time-duration>) and
// returns parsed timestamp given that duration in the past from current time.
// All valid values must contain a number followed by a time-duration, from the following list (long form/short form):
// - second/s
// - minute/m
// - hour/h
// - day/d
// - week/w
// - month/M
// - year/y
// For example, possible input values, and their result:
// - "3d" or "3day" --> three days --> time.Now().Add(-3 * 24 * time.Hour)
// - "2m" or "2minute" --> two minutes --> time.Now().Add(-2 * time.Minute)
// - "1w" or "1week" --> one week --> time.Now().Add(-7 * 24 * time.Hour)
// - "30s" or "30second" --> thirty seconds --> time.Now().Add(-30 * time.Second)
// Note: Duration strings are case-sensitive, and should be used as mentioned above only.
// Limitation: Value of numerical multiplier, X should be in b/w 0 - 1e6 (1 million), boundary values excluded i.e.
// 0 < X < 1e6. Also, the maximum time in the past can be 1 January 1970 00:00:00 UTC (epoch time),
// so giving "1000y" will result in epoch time.
func parseTimeRange(timeRange string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// fallback on to check if it's of longer notation

// using server's local timezone

const (
	// time ranges
	day   = 24 * time.Hour
	week  = 7 * day
	month = 30 * day
	year  = 365 * day
)

// parseTimeDuration parses the given time duration in either short or long convention
// and returns the time.Duration
// Valid values (long notation/short notation):
// - second/s
// - minute/m
// - hour/h
// - day/d
// - week/w
// - month/M
// - year/y
// NOTE: the input "duration" is case-sensitive
func parseTimeDuration(duration string) (dur time.Duration, err error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}
