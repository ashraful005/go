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
x = 10  //assignment
y=&x  //referencing
fmt.Println("x is ", x)    //accessing
fmt.Println("y is ", y)    


//accessing value from a pointer or dereferencing
fmt.Println("y dereferencing value is ", *y)     //dereferencing
}