package main

import (
    "fmt"
    "time"
)

func main(){
currentTime := time.Now()
fmt.Println("The time is", currentTime)


fmt.Println("The year is",currentTime.Year())
fmt.Println("The month is",currentTime.Month())
fmt.Println("The day is",currentTime.Day())
fmt.Println("The hour is",currentTime.Hour())
fmt.Println("The minute is",currentTime.Minute())
fmt.Println("The second is",currentTime.Second())
}