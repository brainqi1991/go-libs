package timeutils

import (
	"fmt"
	"strconv"
	"time"
)

// FMT_TYPE_NORMAL
const (
	DATE_TIME_FMT    = "2006-01-02 15:04:05"
	DATE_FMT         = "2006-01-02"
	TIME_FMT         = "15:04:05"
	DATE_TIME_FMT_CN = "2006年01月02日 15时04分05秒"
	DATE_FMT_CN      = "2006年01月02日"
	TIME_FMT_CN      = "15时04分05秒"
)

const SecondInNano = 1000 * 1000 * 1000

// GetTimestamp return 1455690401 in sec
func GetTimestamp() int64 {
	return time.Now().Unix()
}

// GetTimestampString return 1455690401 in sec string
func GetTimestampString() string {
	return strconv.FormatInt(GetTimestamp(), 10)
}

// GetTimestampInMilli return 1455690401 in millosecond
func GetTimestampInMilli() int64 {
	return int64(time.Now().UnixNano() / (1000 * 1000))
}

// GetTimestampInMilliString return millosecond string
func GetTimestampInMilliString() string {
	return strconv.FormatInt(GetTimestampInMilli(), 10)
}

// GetTimestampInMicro return microsecond
func GetTimestampInMicro() int64 {
	return int64(time.Now().UnixNano() / 1000)
}

// GetTimestampInMicroString return microsecond string
func GetTimestampInMicroString() string {
	return strconv.FormatInt(GetTimestampInMicro(), 10)
}

// GetTimeFormat return the time format
func GetTimeFormat(second int64, format string) string {
	return time.Unix(second, 0).Format(format)
}

// GetCurrentTimeFormat return current time with format
func GetCurrentTimeFormat(format string) string {
	return GetTimeFormat(GetTimestamp(), format)
}

// Elapse Timing the cost of function call,unix nano was returned
func Elapse(f func()) int64 {
	now := time.Now().UnixNano()
	f()
	return time.Now().UnixNano() - now
}

// IsLeapYear check whether a year is leap
func IsLeapYear(year int) bool {
	if year%100 == 0 {
		return year%400 == 0
	}
	return year%4 == 0
}

// GetMonthDays return how many day in the month
func GetMonthDays(year, month int) int {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		if IsLeapYear(year) {
			return 29
		}
		return 28
	default:
		panic(fmt.Sprintf("Illegal month:%d", month))
	}
}
