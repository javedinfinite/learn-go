package main

import ("fmt"
		"time")


func PrintTime () {
	fmt.Println("*******************Printing time*************************")

	fmt.Println(time.Now())

	// time.Date(year, month, day, hour, min, sec, nsec, location)
	myTime := time.Date(2016, time.March, 11, 4, 5, 0, 0, time.Local) 
	fmt.Println("myTime : ", myTime)
	fmt.Println("extracting time components", myTime.Second(), myTime.Day(), myTime.Weekday(), time.Now().Day())
	fmt.Println("Comparing settime and nowtime : ", myTime.Before(time.Now()), myTime.After(time.Now()), myTime.Equal(time.Now()))
	timeDiff := time.Now().Sub(myTime)
	fmt.Println("Check time difference : ", timeDiff, timeDiff.Hours())
	fmt.Println("Add subtract to the time : ", myTime.Add(1 * time.Hour), myTime.Add(-1 * time.Hour), myTime.Add(1 * time.Hour + 2 * time.Minute))
	fmt.Println("Add date : ", myTime.AddDate(1, 2, 3))

	fmt.Println("************* Printing date formatting with fixed placeholder in go*********************")
	// In Go, formatting and parsing time is NOT like other languages.
	// Instead of saying: “use YYYY-MM-DD format”, Go says: “show me an example time, and I’ll copy its shape”
	// That example is always: Mon Jan 2 15:04:05 MST 2006 , This is called the reference time.
	inputFormat := "27-04-2026" 
	timeFormat, err := time.Parse("02-01-2006", inputFormat) //string → time,  telling Go: “My input looks like: day-month-year”, inputFormat is a string becomes a real time object
		if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("parsed date : ", timeFormat)
	fmt.Println("formatted date : ", timeFormat.Format("01-2006-02")) // time → string,  “How do I want to display it using Go’s magic reference date?”, must use Go’s “dictionary words” (2006, 01, 02, etc.)
	fmt.Println("RFC3339 : ", timeFormat.Format(time.RFC3339))

	// An epoch is a fixed point in time that a system uses as a reference to measure and store dates and times.
	// In most modern computing (including Go and Unix-like systems), the standard epoch is: January 1, 1970, 00:00:00 UTC. This is called the Unix epoch.

	// Why epoch is useful: 
	// Simple numeric representation, Easy comparison and sorting, Widely used in databases, APIs, and operating systems
	fmt.Println("epoch time : ", time.Unix(0,0), time.Now().Unix()) // this gives current time in seconds from January 1, 1970, 00:00:00 UTC, same UnixMilli, UnixNano 
	fmt.Println("local epoch time", time.Unix(time.Now().Unix(), 0)) // A clean timestamp rounded to seconds. Take the current time, convert it to epoch seconds, then reconstruct a time object from those seconds


}