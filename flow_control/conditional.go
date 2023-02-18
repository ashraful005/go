package main

import "fmt"

func main(){
fmt.Print("Enter your age: ")
var age int
fmt.Scanf("%d", &age)

//if-else statement
if age<10{
fmt.Println("Children")
}else{
fmt.Println("Teenager")
}
}