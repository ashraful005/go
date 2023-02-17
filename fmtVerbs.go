package main

import "fmt" 

func main(){
//boolean 
var isFound bool
fmt.Println(isFound)

//logical operators
fmt.Println(true && false)
fmt.Println(true && true)
fmt.Println(true || false)
fmt.Println(false || false)
fmt.Println(!true)

//default values of some datatypes
var c rune
var age int
var result float32
var name string
var found bool

fmt.Println(c, age, result, name, found)

}