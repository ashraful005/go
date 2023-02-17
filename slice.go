package main

import "fmt"

func main(){
//slices are segments of array
//slicing from an array
students :=[]string{"Ashraful", "Asgor", "Mostain"}

x := students[0:2]
fmt.Println(x)
}