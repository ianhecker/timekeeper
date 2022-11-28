package shift

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// ## 2022-11-28:0830-0930
const ShiftDateTimeFormat = `^##\s(?P<year>\d{4})\-(?P<month>\d{2})\-(?P<day>\d{2}):(?P<startHour>\d{2})(?P<startMinutes>\d{2})\-(?P<endHour>\d{2})(?P<endMinutes>\d{2})$`

type Shift struct {
	Start time.Time
	End   time.Time
}

func (s *Shift) Unmarshal(markdown string) error {
	r := regexp.MustCompile(ShiftDateTimeFormat)
	matches := r.FindStringSubmatch(markdown)

	if len(matches) != 8 {
		return fmt.Errorf("markdown string: %s does not match regex", markdown)
	}

	year, err := strconv.Atoi(matches[r.SubexpIndex("year")])
	if err != nil {
		return err
	}

	monthInt, err := strconv.Atoi(matches[r.SubexpIndex("month")])
	if err != nil {
		return err
	}

	month, err := parseMonth(monthInt)
	if err != nil {
		return err
	}

	day, err := strconv.Atoi(matches[r.SubexpIndex("day")])
	if err != nil {
		return err
	}

	startHour, err := strconv.Atoi(matches[r.SubexpIndex("startHour")])
	if err != nil {
		return err
	}

	startMinutes, err := strconv.Atoi(matches[r.SubexpIndex("startMinutes")])
	if err != nil {
		return err
	}

	endHour, err := strconv.Atoi(matches[r.SubexpIndex("endHour")])
	if err != nil {
		return err
	}

	endMinutes, err := strconv.Atoi(matches[r.SubexpIndex("endMinutes")])
	if err != nil {
		return err
	}

	s.Start = parseTime(year, month, day, startHour, startMinutes)
	s.End = parseTime(year, month, day, endHour, endMinutes)
	return nil
}

func parseTime(year int, month time.Month, day, hour, minutes int) time.Time {
	var sec int = 0
	var nsec int = 0
	var loc *time.Location = time.UTC
	return time.Date(year, month, day, hour, minutes, sec, nsec, loc)
}

func parseMonth(month int) (time.Month, error) {
	tm := time.Month(month + 1)
	if tm > 13 {
		return 0, fmt.Errorf("invalid month %s", tm.String())
	}
	return tm, nil
}
