// Go offers extensive support for times and durations;
// here are some examples.

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// We'll start by getting the current time.
	now := time.Now()
	p(now)

	// You can build a `time` struct by providing the
	// year, month, day, etc. Times are always associated
	// with a `Location`, i.e. time zone.

	// TODO: Create then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	// TODO: Print then
	fmt.Println(then)

	// You can extract the various components of the time
	// value as expected.

	// TODO: Print then.Year()
	fmt.Println(then.Year())
	// TODO: Print then.Month()
	fmt.Println(then.Month())
	// TODO: Print then.Day()
	fmt.Println(then.Day())
	// TODO: Print then.Hour()
	fmt.Println(then.Hour())
	// TODO: Print then.Minute()
	fmt.Println(then.Minute())
	// TODO: Print then.Second()
	fmt.Println(then.Second())
		// TODO: Print then.Nanosecond()
	fmt.Println(then.Nanosecond())
	// TODO: Print then.Location()
	fmt.Println(then.Location())
	// The Monday-Sunday `Weekday` is also available.

	// TODO: Print then.Weekday()
	fmt.Println(then.Weekday())
	// These methods compare two times, testing if the
	// first occurs before, after, or at the same time
	// as the second, respectively.

	// TODO: Print then.Before(now)
	fmt.Println(then.Before(now))
	// TODO: Print then.After(now)
	fmt.Println(then.After(now))
	// TODO: Print then.Equal(now)
	fmt.Println(then.Equal(now))

	// The `Sub` methods returns a `Duration` representing
	// the interval between two times.

	// TODO: Create diff := now.Sub(then)
	// TODO: Print diff
	diff := now.Sub(then)
	fmt.Println(diff)

	// We can compute the length of the duration in
	// various units.

	// TODO: Print diff.Hours()
	fmt.Println(diff.Hours())
	// TODO: Print diff.Minutes()
	fmt.Println(diff.Minutes())
	// TODO: Print diff.Seconds()
	fmt.Println(diff.Seconds())
	// TODO: Print diff.Nanoseconds()
	fmt.Println(diff.Nanoseconds())

	// You can use `Add` to advance a time by a given
	// duration, or with a `-` to move backwards by a
	// duration.

	// TODO: Print then.Add(diff)
	// TODO: Print then.Add(-diff)
	fmt.Println(then.Add(diff))
	fmt.Println(then.Add(-diff))
}