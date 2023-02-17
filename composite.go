package main 

import "fmt"

func main(){
//primitive data types
//rune, byte = alias primitive
//int, float32, string, bool = original primitive

//composite data types : array, slice, map, struct


//array declaration
var students [3]string
fmt.Println(students)
fmt.Println(len(students))


//array characteristics
//fixed length, same type
//array initialization
students[0]="Ashraful"
students[1]="Mostain"
students[2]="Ashgor"

fmt.Println(students[2])


}