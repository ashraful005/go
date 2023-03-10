package main

import ("fmt"
        "reflect"
)

func main(){
//slices are segments of array
//slicing from an array
students :=[...]string{"Ashraful", "Asgor", "Mostain"}

x := students[0:2]

//creating a slice
y := make([]string, 0)

//slice declaration
var fruits []string
//slice initialization
fruits = append(fruits, "Apple","Banana","Mango")


fmt.Println(x)
fmt.Println(y)
fmt.Println(fruits, len(fruits))
fmt.Printf("%T\n", fruits)
fmt.Printf("%T\n", students)

//reflect package
a := reflect.TypeOf(students).Kind().String()
b := reflect.TypeOf(fruits).Kind().String()
fmt.Println(a)
fmt.Println(b)
}