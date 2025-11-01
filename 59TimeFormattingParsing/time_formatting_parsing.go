// Go supports time formatting and parsing via
// pattern-based layouts.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Here's a basic example of formatting a time
	// according to RFC3339, using the corresponding layout
	// constant.

	// TODO: Create t := time.Now()
	// TODO: Print t.Format(time.RFC3339)

	t := time.Now()
	fmt.Println("t:", t.Format(time.RFC3339))
	

	// Time parsing uses the same layout values as `Format`.

	// TODO: Create t1, e := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	// TODO: Print t1
	t1, e := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	fmt.Println("t1:", t1)
	fmt.Println("e:", e)

	// `Format` and `Parse` use example-based layouts. Usually
	// you'll use a constant from `time` for these layouts, but
	// you can also supply custom layouts. Layouts must use the
	// reference time `Mon Jan 2 15:04:05 MST 2006` to show the
	// pattern with which to format/parse a given time/string.
	// The example time must be exactly as shown: the year 2006,
	// 15 for the hour, Monday for the day of the week, etc.

	// TODO: Print t.Format("3:04PM")
	// TODO: Print t.Format("Mon Jan _2 15:04:05 2006")
	// TODO: Print t.Format("2006-01-02T15:04:05.999999-07:00")
	fmt.Println("t.Format(\"3:04PM\"):", t.Format("3:04PM"))
	fmt.Println("t.Format(\"Mon Jan _2 15:04:05 2006\"):", t.Format("Mon Jan _2 15:04:05 2006"))
	fmt.Println("t.Format(\"2006-01-02T15:04:05.999999-07:00\"):", t.Format("2006-01-02T15:04:05.999999-07:00"))
	

	// TODO: Create form := "3 04 PM"
	// TODO: Create t2, e := time.Parse(form, "8 41 PM")
	// TODO: Print t2
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	fmt.Println("t2:", t2)
	fmt.Println("e:", e)

	// For purely numeric representations you can also
	// use standard string formatting with the extracted
	// components of the time value.

	// TODO: Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00")
	// TODO: with t.Year(), t.Month(), t.Day(),
	// t.Hour(), t.Minute(), t.Second()
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	// `Parse` will return an error on malformed input
	// explaining the parsing problem.

	// TODO: Create ansic := "Mon Jan _2 15:04:05 2006"
	// TODO: Create _, e = time.Parse(ansic, "8:41PM")
	// TODO: Print e
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	fmt.Println("e:", e)
}