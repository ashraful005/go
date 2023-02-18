package main

import "fmt"

func main() {
	l := 20
	b := 30
	
	//closure function is a special anonymous function
	//doen't have any parameters and 
	//can access variables outside the function body
	func() {
		var area int
		area = l * b
		fmt.Println(area)
	}()
}