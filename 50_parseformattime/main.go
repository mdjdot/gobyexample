package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	t := time.Now()
	p(t.Format((time.RFC3339)))

	t1, e := time.Parse(time.RFC3339, "2019-12-01T21:52:16+08:00")
	if e != nil {
		p(e)
	}
	p(t1)

	p(t.Format("3:04PM"))                              // 9:55PM
	p(t.Format("15:04"))                               // 21:56
	p(t.Format("Mon Jan _2 15:04:05 2006"))            // Sun Dec  1 21:56:59 2019
	p(t.Format("2006-01-02T15:04:05.999999999Z07:00")) // 2019-12-01T21:57:37.2703532+08:00

	form := "3 04 PM"
	t2, err := time.Parse(form, "8 41 PM")
	if err != nil {
		p(err)
	}
	p(t2) // 0000-01-01 20:41:00 +0000 UTC

	ansic := "Mon Jan _2 15:04:05 2006"
	t3, err := time.Parse(ansic, "8:41PM")
	if err != nil {
		p(e) // <nil>
	}
	p(t3) // 0001-01-01 00:00:00 +0000 UTC
}

/*
time.Format
const (
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
)
*/
