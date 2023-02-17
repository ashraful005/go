package main

import "fmt"


/*
//struct defination
type Book struct{
Title string
Author string
ISBN string
Price float32
Pages int
}
*/

func main(){
//struct combines two or more different data types
//shorthend struct declaration - anonymous struct
b1 := struct{
Title string
Author string
ISBN string
Price float32
Pages int
}{
Title : "An Introduction to Programming in Go",
Author : "CALEB DOXSEY",
ISBN : "978-1478355823",
Price : 10.50, 
Pages : 165,
}

//struct initialization
b1.Title = "An Introduction to Programming in Go"
b1.Author = "Caleb Doxsey"
b1.ISBN = "978-1478355823"
b1.Price = 10.50
b1.Pages = 165

//custom datatype declaration
var w1 Weight

//custom variable initialization
w1 = 30.5

//data retreival
fmt.Println(b1)
fmt.Println(b1.Title)
fmt.Println(b1.Author)
fmt.Println(b1.ISBN)
fmt.Println(b1.Price)
fmt.Println(b1.Pages)

//custom datatype
fmt.Println(w1)

}