package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("The time is", currentTime)

	fmt.Println("The year is", currentTime.Year())
	fmt.Println("The month is", currentTime.Month())
	fmt.Println("The day is", currentTime.Day())
	fmt.Println("The hour is", currentTime.Hour())
	fmt.Println("The minute is", currentTime.Minute())
	fmt.Println("The second is", currentTime.Second())

	//all in a single line
	fmt.Printf("%d-%d-%d %d:%d:%d\n", currentTime.Year(), currentTime.Month(), currentTime.Day(),
		currentTime.Hour(), currentTime.Minute(), currentTime.Second())

//formatting a specific date
theDate := time.Date(2000,6,30,13,24,30,100,time.Local)
fmt.Println("The time is ",theDate)

//formatting the date with specified format
fmt.Println(theDate.Format("2006-1-2 15:4:5"))

fmt.Println(theDate.Format("2006-01-02 03:04:05 pm"))

//using predefined format
fmt.Println(theDate.Format(time.RFC3339Nano))

//parsing strings into time
timeString := "2021-08-15T14:30:45.0000001-05:00"
theTime, err := time.Parse(time.RFC3339Nano, timeString)
if err != nil{
fmt.Println("Could not parse time:", err)
}
fmt.Println("The time is", theTime)

fmt.Println(theTime.Format(time.RFC3339Nano))

dateTime()

//age calculation
hours:=time.Now().Sub(theDate)
fmt.Println(hours)
}



func dateTime() {
	firstTime := time.Date(2021, 8, 15, 14, 30, 45, 100, time.UTC)
	fmt.Println("The first time is", firstTime)

	secondTime := time.Date(2021, 12, 25, 16, 40, 55, 200, time.UTC)
	fmt.Println("The second time is", secondTime)

	fmt.Println("First time before second?", firstTime.Before(secondTime))
	fmt.Println("First time after second?", firstTime.After(secondTime))

	fmt.Println("Second time before first?", secondTime.Before(firstTime))
	fmt.Println("Second time after first?", secondTime.After(firstTime))
}
