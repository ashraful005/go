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
theDate := time.Date(2000,6,30,6,24,30,100,time.Local)
fmt.Println("The time is ",theDate)

//formatting the date with specified format
fmt.Println(theDate.Format("2006-1-2 15:4:5"))
}
