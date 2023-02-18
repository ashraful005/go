package main

import "fmt"

func main(){
fmt.Print("Enter your age: ")
var age int
fmt.Scanf("%d", &age)

//if-else if statement
if age>0 && age<3{

fmt.Println("Infant")

}else if age>2 && age<13{

fmt.Println("Children")

}else if age>12 && age<20{

fmt.Println("Teenager")

}else if age>19{

fmt.Println("Adult")
}
}