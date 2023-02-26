package gtime

import "time"

const TimeLayout = "2006-01-02 15:04:05"

func ParseTime(target string) (time.Time, error) {
	location, _ := time.LoadLocation("Asia/Shanghai")
	inLocation, err := time.ParseInLocation(TimeLayout, target, location)
	if err != nil {
		return time.Time{}, err
	}
	return inLocation, err
}
