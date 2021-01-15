package calendar

import "time"

// GetDates 取得當月有幾天
func GetDates(t time.Time) int {
	year, month, _ := t.Date()

	return time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

// AddMonth 增加月份
func AddMonth(t time.Time) time.Time {
	return t.AddDate(0, 0, GetDates(t)-t.Day()+1)
}
