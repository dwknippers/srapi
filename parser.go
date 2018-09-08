package main

import (
	"regexp"
	"time"
)

// Event is a time when a class or activity takes place
type Event struct {
	Start        time.Time
	End          time.Time
	Course       string
	TeacherName  string
	Type         string
	LocationCode string
}

// Day is a date when event may take place
type Day struct {
	Date   time.Time
	Events []Event
}

var re = map[string]*regexp.Regexp{
	"days":          regexp.MustCompile(`(?sU)<a name="([0-9]{4}-[0-9]{2}-[0-9]{2})".*</table>`),
	"time":          regexp.MustCompile(`(?sU)(([0-9]{2}:[0-9]{2}) - ([0-9]{2}:[0-9]{2})).*(<th.*class="times"|</table>)`),
	"courseAndType": regexp.MustCompile(`<span>(.*)</span>`),
	"teacher":       regexp.MustCompile(`(?sU)teacher.*">(.*)</a>`),
	"nbsp":          regexp.MustCompile(`&nbsp;`),
	"locationCode":  regexp.MustCompile(`(?sU)<span class="pull-right">.*([A-Z][0-9].[0-9]{2}).*</span>`),
}

func assignMatch(out *string, in *string, re *regexp.Regexp, index int) {
	match := re.FindStringSubmatch(*out)
	if len(match) > 0 {
		*in = match[index]
	}
}

func parse(scheduleHTML string) []Day {
	daysMatch := re["days"].FindAllStringSubmatch(scheduleHTML, -1)

	var days = make([]Day, len(daysMatch))

	for d := 0; d < len(daysMatch); d++ {
		date := daysMatch[d][1]
		parsedDate, _ := time.Parse("2006-01-02", date)

		timesMatch := re["time"].FindAllStringSubmatch(string(daysMatch[d][0]), -1)

		day := Day{
			Date:   parsedDate,
			Events: make([]Event, len(timesMatch)),
		}

		for t := 0; t < len(timesMatch); t++ {
			start, _ := time.Parse("15:04", timesMatch[t][2])
			end, _ := time.Parse("15:04", timesMatch[t][3])

			courseAndType := re["courseAndType"].FindAllStringSubmatch(timesMatch[t][0], -1)

			course := courseAndType[0][1]
			var teacherName string
			assignMatch(&timesMatch[t][0], &teacherName, re["teacher"], 1)
			eventType := re["nbsp"].ReplaceAllString(courseAndType[1][1], "")

			var locationCode string
			assignMatch(&timesMatch[t][0], &locationCode, re["locationCode"], 1)

			day.Events[t] = Event{
				Start:        start,
				End:          end,
				Course:       course,
				TeacherName:  teacherName,
				Type:         eventType,
				LocationCode: locationCode,
			}
		}
		days[d] = day
	}
	return days
}
