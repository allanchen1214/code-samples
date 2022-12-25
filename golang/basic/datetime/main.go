package main

import (
	"fmt"
	"time"
)

/*
// Build-in format
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

func main() {
	now := time.Now()
	fmt.Println(now) // Time

	ts := time.Now().Unix()
	fmt.Println(ts) // int64

	fmt.Println(now.Format("Mon, Jan 2, 2006 at 3:04pm")) // Sun, Dec 25, 2022 at 7:11pm

	fmt.Println("Year: ", now.Year())   // 2022
	fmt.Println("Month: ", now.Month()) // December

	input_format := "1/2/2006 3:04pm"
	user_str := "4/16/2014 11:38am"
	user_date, err := time.Parse(input_format, user_str)
	if err != nil {
		fmt.Println(">>> Error parsing date string")
	}
	fmt.Println("User date = ", user_date.Format("Jan 2, 2006 @ 3:04pm"))

	now2 := time.Now().Add(time.Duration(30 * time.Second))
	b1 := now.Before(now2)
	b2 := now.After(now2)
	b3 := now.Equal(now2)
	fmt.Printf("b1: %v b2: %v b3: %v\n", b1, b2, b3)

	diff := now.Sub(now2)
	fmt.Println(diff)
	diff = now2.Sub(now)
	fmt.Println(diff.Hours())
	fmt.Println(diff.Minutes())
	fmt.Println(diff.Seconds())
	fmt.Println(diff.Milliseconds())

	twodaysDiff := time.Hour * 24 * 2
	twodays := now.Add(twodaysDiff)
	fmt.Println("Two Days: ", twodays.Format(time.ANSIC))

	est, _ := time.LoadLocation("EST")
	july4 := time.Date(1776, 7, 4, 12, 15, 0, 0, est)
	fmt.Println(july4)
}
