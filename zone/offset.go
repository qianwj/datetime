package zone

import (
	"math"
	"strconv"
	"strings"
)

const (
	secondsPerHour   = 60 * 60
	secondsPerMinute = 60
	minutesPerHour   = 60
)

type Offset struct {
	id      string
	seconds int64
}

func OfHours(hours int64) *Offset {
	return OfHoursMinutes(hours, 0)
}

func OfHoursMinutes(hours, minutes int64) *Offset {
	return OfHoursMinutesSeconds(hours, minutes, 0)
}

func OfHoursMinutesSeconds(hours, minutes, seconds int64) *Offset {
	total := totalSeconds(hours, minutes, seconds)
	return &Offset{
		id:      buildOffsetId(total),
		seconds: total,
	}
}

//
//func OfTotalSeconds(total int64) *Offset {
//
//}

func totalSeconds(hours, minutes, seconds int64) int64 {
	return hours*secondsPerHour + minutes*secondsPerMinute + seconds
}

func buildOffsetId(totalSeconds int64) string {
	if totalSeconds == 0 {
		return "Z"
	} else {
		absTotalSeconds := int64(math.Abs(float64(totalSeconds)))
		var buf strings.Builder
		absHours := absTotalSeconds / secondsPerHour
		absMinutes := (absTotalSeconds / secondsPerMinute) % minutesPerHour
		newCondition(totalSeconds < 0).If(func() { buf.WriteString("-") }).Else(func() { buf.WriteString("+") }).Exec()
		newCondition(absHours < 10).If(func() { buf.WriteString("0") }).Else(func() { buf.WriteString("") }).Exec()
		buf.WriteString(strconv.FormatInt(absHours, 10))
		newCondition(absMinutes < 10).If(func() { buf.WriteString(":0") }).Else(func() { buf.WriteString(":") }).Exec()
		buf.WriteString(strconv.FormatInt(absMinutes, 10))
		absSeconds := absTotalSeconds % secondsPerMinute
		if absSeconds != 0 {
			newCondition(absSeconds < 10).If(func() { buf.WriteString(":0") }).Else(func() { buf.WriteString(":") }).Exec()
			buf.WriteString(strconv.FormatInt(absSeconds, 10))
		}
		return buf.String()
	}
}

type condition struct {
	test       bool
	ifBranch   func()
	elseBranch func()
}

func newCondition(test bool) *condition {
	return &condition{test: test}
}

func (c *condition) If(ifBranch func()) *condition {
	c.ifBranch = ifBranch
	return c
}

func (c *condition) Else(elseBranch func()) *condition {
	c.elseBranch = elseBranch
	return c
}

func (c *condition) Exec() {
	if c.test {
		if c.ifBranch != nil {
			c.ifBranch()
		}
	} else {
		if c.elseBranch != nil {
			c.elseBranch()
		}
	}
}
