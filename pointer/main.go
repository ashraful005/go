package main

import "fmt"

func main(){
var x int
fmt.Println("x value is ", x)
fmt.Println("x memory address is ", &x)

//pointer declaration
var y *int
fmt.Println("y value is ", y)
fmt.Println("y memory address is ", &y)


//pointer initialization
x = 10
y=&x
fmt.Println("x is ", x)
fmt.Println("y is ", y)
}