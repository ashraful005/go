package main

import "fmt"

func main(){
var age,sum float32
sum=0
var avg float32
fmt.Println("Enter age:")
for i:=1;i<=5;i++{
fmt.Scanf("%f",&age)
sum=sum+age
}
avg=sum/5.0
fmt.Printf("The average age is %f years", avg)
}