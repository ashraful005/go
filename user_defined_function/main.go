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

func main(){
a, p := rectangle(10,10)
fmt.Println(a, p)
fmt.Println(Add(5,6))

}