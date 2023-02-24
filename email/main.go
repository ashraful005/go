package main

import (
	"encoding/base64"
	"fmt"
)

func main(){
	resp := []byte("\x00"+ "alam1809005@stud.kuet.ac.bd" + "\x00" + "your password")
	//fmt.Println(string(resp), resp)

	sEnc := base64.StdEncoding.EncodeToString([]byte(resp))
	fmt.Println(sEnc)
}