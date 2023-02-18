package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	nl, err := net.Listener("tcp", ":8888")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	conn, err := nl.Accept()
	if err != nil {
		fmt.Println(err.Error())
	}
	bs := make([]byte, 1024)

	n, e := conn.Read(bs)
	if e != nil {
		fmt.Println(e.Error())
	}

	fmt.Println(n)

	reqstr := string(bs) //conversion
	fmt.Println(reqstr)

}
