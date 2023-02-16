package main

import "fmt"

func main() {

	fmt.Println("Enter name and age: ")
	var name string
	var age int

	fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("%s is %d years old\n", name, age)
	if age >= 1 && age <= 10 {
		fmt.Println("Age between 1 and 10")
	} else {
		fmt.Println("Age not between 1 and 10")
	}
}
