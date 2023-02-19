package main

import "fmt"

func main() {
	//Decimal number to character
	fmt.Printf("%c\n", 68)

	//Hexadecimal number to character
	fmt.Printf("%c\n", 0x49)

	//character to decimal
	fmt.Printf("%d\n", 'c')

	//character to hexa
	fmt.Printf("%x\n", 'A')

	//decimal to hexa
	fmt.Printf("%x\n", 65)

	//function calling
	NumtoChar(65)
	
	//multiple variable of different type
	var name, age, gpa = "Ashraful", 23, 5.00

	fmt.Println(name, age, gpa)


	fmt.Println(sum(5,10, 20, 30))
}

func NumtoChar(x int) {
	fmt.Printf("%c\n", x)
}


//variadic
func sum(a ...int) (r int){
for _, v := range a{
r = r + v
}
return
}
