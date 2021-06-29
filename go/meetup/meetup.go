package meetup

import "time"

type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Teenth
	Last
)

func Day(week WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	date := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

	switch week {
	case First:
		return getNextWeekday(date, weekday, 1)
	case Second:
		return getNextWeekday(date.AddDate(0, 0, 7), weekday, 1)
	case Third:
		return getNextWeekday(date.AddDate(0, 0, 14), weekday, 1)
	case Fourth:
		return getNextWeekday(date.AddDate(0, 0, 21), weekday, 1)
	case Last:
		return getNextWeekday(date.AddDate(0, 1, -1), weekday, -1)
	case Teenth:
		return getNextWeekday(date.AddDate(0, 0, 12), weekday, 1)
	default:
		return 0
	}
}

func getNextWeekday(date time.Time, weekday time.Weekday, step int) int {
	for {
		if date.Weekday() == weekday {
			return date.Day()
		}
		date = date.AddDate(0, 0, step)
	}
}
