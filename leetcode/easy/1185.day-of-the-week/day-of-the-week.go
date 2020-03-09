package _185_day_of_the_week

import "time"

func dayOfTheWeek(day int, month int, year int) string {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local).Weekday().String()
}
