package main

import "fmt"


//struct defination
type Book struct{
Title string
Author string
ISBN string
Price float32
Pages int
}


func main(){
//struct combines two or more different data types
//struct declaration
var b1 Book

fmt.Println(b1)
}