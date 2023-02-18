package main

import "fmt"

func main(){
fmt.Print("Enter your age: ")
var age int
fmt.Scanf("%d", &age)

//if statement
if age<10{
fmt.Println("Children")
}
}