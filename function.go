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
}

func NumtoChar(x int) {
	fmt.Printf("%c\n", x)
}
