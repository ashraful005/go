package main

import "fmt"

func main(){
//slices are segments of array
//slicing from an array
students :=[]string{"Ashraful", "Asgor", "Mostain"}

x := students[0:2]

//creating a slice
y := make([]string, 0)
fmt.Println(x)
fmt.Println(y)
}