package main

import "fmt"

/*
func add(a int,b int) int{
temp := a+b
return temp
}
*/

/*
//example-02
func add(a,b int)int{
temp := a+b
return temp
}
*/

/*
//example-03
func add(a int,b int) (temp int){
temp = a+b
return temp
}
*/


//example-04
func add(a int,b int) (temp int){
temp = a+b
return 
}


//case sensitive
func Add(a int,b int) (temp int){
temp = a*b
return 
}

//rectangle
func rectangle(a,b int) (area int, perimeter int){
area = a*b
perimeter = 2*(a+b)
return
}


//passing pointer or address of a variable
func update(a *int, t *string){
*a = *a + 5
*t = *t + " Doe"
return
}

func main(){
ar, p := rectangle(10,10)
fmt.Println(ar, p)
fmt.Println(Add(5,6))

a := 10
t := "Mostain"
fmt.Println(a, t)
update(&a, &t)
fmt.Println(a, t)

}