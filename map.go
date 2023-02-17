package main 


import ("fmt"
	
)

func main(){
//unordered collection of key-value pairs
//map declaration
x :=  make(map[string]int)

//map initialization
x["Ashraful"] = 23
x["Mostain"] = 40
x["Mehedi"] = 24

fmt.Println(x)
fmt.Println(x["Ashraful"])
}