package parsers

import (
	"fmt"
	"net/http"
	"time"
)

var dateFormats = []string{
	"02.01.2006",
	"2.1.2006",
	"2.Jan.2006",
	"02.Jan.2006",
	"02 01 2006",
	"2 1 2006",
	"2 Jan 2006",
	"02 Jan 2006",
	"02.01.06",
	"2.1.06",
	"2.Jan.06",
	"02.Jan.06",
	"2.January.2006",
	"02.January.2006",
	time.ANSIC,
	time.UnixDate,
	time.RubyDate,
	time.RFC822,
	time.RFC822Z,
	time.RFC850,
	time.RFC1123,
	time.RFC1123Z,
	time.RFC3339,
	time.RFC3339Nano,
	time.Kitchen,
	time.Stamp,
	time.StampMilli,
	time.StampMicro,
	time.StampNano,
	http.TimeFormat,
}

func ParseDate(date string) (time.Time, error) {
	var (
		format   string
		dateTime time.Time
		err      error
	)

	for _, format = range dateFormats {
		if dateTime, err = time.Parse(format, date); err == nil {
			return dateTime, nil
		}
	}
	return time.Time{}, fmt.Errorf("Could not parse date '%s'", date)

}
