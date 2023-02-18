package main 

import "fmt"

func main(){
/*
//for loop
for i:=1;i<=10;i++ {
fmt.Println(i)
}
*/

//range for loop
students := []string{"Asgor", "Mainul", "Anonnya"}
for i, val := range students {
fmt.Printf("At position %d, the character %s is present\n",i,val)
}

}