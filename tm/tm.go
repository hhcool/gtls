package tm

import "time"

const (
	TimeFormat     = "2006-01-02 15:04:05"
	TimeFormatDate = "2006-01-02"
)

func Now() string {
	return time.Now().Format(TimeFormat)
}
func NowDate() string {
	return time.Now().Format(TimeFormatDate)
}

func StartOfDay(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}
func EndOfDay(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 23, 59, 59, 0, d.Location())
}
func StartOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return StartOfDay(d)
}
func EndOfMonth(d time.Time) time.Time {
	return EndOfDay(StartOfMonth(d).AddDate(0, 1, -1))
}
func StartOfYear(d time.Time) time.Time {
	return time.Date(d.Year(), 1, 1, 0, 0, 0, 0, d.Location())
}
func EndOfYear(d time.Time) time.Time {
	return EndOfDay(StartOfYear(d).AddDate(1, 0, -1))
}
