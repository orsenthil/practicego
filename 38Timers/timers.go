// We often want to execute Go code at some point in the
// future, or repeatedly at some interval. Go's built-in
// _timer_ and _ticker_ features make both of these tasks
// easy. We'll look first at timers and then
// at [tickers](tickers).

package main

import (
	"fmt"
	"time"
)

func main() {

	// Timers represent a single event in the future. You
	// tell the timer how long you want to wait, and it
	// provides a channel that will be notified at that
	// time. This timer will wait 2 seconds.

	// TODO: Create timer1 with 2 seconds
	timer1 := time.NewTimer(time.Second * 2)

	// The `<-timer1.C` blocks on the timer's channel `C`
	// until it sends a value indicating that the timer
	// fired.

	// TODO: Receive from timer1.C and print the result
	<-timer1.C
	fmt.Println("Timer 1 fired")
	// If you just wanted to wait, you could have used
	// `time.Sleep`. One reason a timer may be useful is
	// that you can cancel the timer before it fires.
	// Here's an example of that.

	// TODO: Create timer2 with 1 second using NewTimer
	timer2 := time.NewTimer(time.Second * 1)
	// TODO: Creat a goroutine that receives from timer2.C and prints the result
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	// TODO: Create stop2 with timer2.Stop()
	stop2 := timer2.Stop()
	// TODO: If stop2 is true, print "Timer 2 stopped"
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	// TODO: Give the timer2 enough time to fire, if it ever
	// was going to, to show it is in fact stopped.
	time.Sleep(time.Second * 2)
}