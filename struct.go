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

//struct initialization
b1.Title = "An Introduction to Programming in Go"
b1.Author = "Caleb Doxsey"
b1.ISBN = "978-1478355823"
b1.Price = 10.50
b1.Pages = 165

fmt.Println(b1)
}